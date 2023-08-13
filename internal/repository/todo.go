package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/hampgoodwin/errors"
	"github.com/hampgoodwin/todo/internal/meta"
	"github.com/jackc/pgx/v5"
	"github.com/segmentio/ksuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func (r *Repository) ListToDos(ctx context.Context, ids []string, cursor string, limit int32) ([]ToDo, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "internal.repository.CreateToDos", trace.WithAttributes(
		attribute.StringSlice("ids", ids),
		attribute.String("ids", cursor),
		attribute.Int64("ids", int64(limit)),
	))
	defer span.End()

	tx, err := r.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, errors.WithErrorMessage(err, errors.NotKnown, "beginning list to dos db transaction")
	}

	query := `select id, message, details, due_date, priority, level_of_effort, created_at, updated_at, deleted_at
		FROM to_dos
		WHERE 1=1`
	var params []interface{}
	if ids != nil {
		params = append(params, ids)
		query += fmt.Sprintf(" AND id IN ($%d)", len(params))
	}
	if cursor != "" {
		params = append(params, cursor)
		query += fmt.Sprintf(" AND id = $%d", len(params))
	}
	query += " ORDER BY id DESC"
	params = append(params, limit)
	query += fmt.Sprintf(" LIMIT $%d", len(params))
	query += ";"

	rows, err := tx.Query(ctx, query, params...)
	if err != nil {
		return nil, errors.WithErrorMessage(err, errors.NotKnown, "querying to dos from database")
	}
	defer rows.Close()

	returningToDos := []ToDo{}
	for rows.Next() {
		returningToDo := ToDo{}
		if err := rows.Scan(
			&returningToDo.ID,
			&returningToDo.Message,
			&returningToDo.Details,
			&returningToDo.DueDate,
			&returningToDo.Priority,
			&returningToDo.LevelOfEffort,
			&returningToDo.CreatedAt,
			&returningToDo.UpdatedAt,
			&returningToDo.DeletedAt,
		); err != nil {
			return nil, errors.WithErrorMessage(err, errors.NotKnown, "scanning list of to dos result row")
		}

		// a more performant way definitely exists, but I've already spent too long on this project to look at doing that :)
		statusRows, err := r.database.Query(ctx, `select id, status, created_at, updated_at, deleted_at
			FROM to_dos_statuses
			WHERE to_do_id = $1;`, returningToDo.ID)
		if err != nil {
			return nil, errors.WithErrorMessage(err, errors.NotKnown, "querying to dos statuses from database")
		}
		defer statusRows.Close()

		returningToDoStatuses := []ToDoStatus{}
		for statusRows.Next() {
			returningToDoStatus := ToDoStatus{}
			if err := statusRows.Scan(
				&returningToDoStatus.ID,
				&returningToDoStatus.Status,
				&returningToDoStatus.CreatedAt,
				&returningToDoStatus.UpdatedAt,
				&returningToDoStatus.DeletedAt,
			); err != nil {
				return nil, errors.WithErrorMessage(err, errors.NotKnown, "scanning list of to dos statuses result row")
			}

			returningToDoStatuses = append(returningToDoStatuses, returningToDoStatus)
		}

		returningToDo.ToDoStatuses = returningToDoStatuses

		returningToDos = append(returningToDos, returningToDo)
	}

	if err := tx.Commit(ctx); err != nil {
		return returningToDos, errors.WithErrorMessage(err, errors.NotKnown, "commiting ToDos query(s)")
	}

	return returningToDos, nil
}

func (r *Repository) CreateToDos(ctx context.Context, creates []ToDo) ([]ToDo, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "internal.repository.CreateToDos", trace.WithAttributes(
		attribute.Int("creates_len", len(creates)),
	))
	defer span.End()

	now := time.Now()

	tx, err := r.database.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, errors.WithErrorMessage(err, errors.NotKnown, "beginning create to dos db transaction")
	}

	returningToDos := []ToDo{}
	for _, create := range creates {
		returningToDo := ToDo{}
		if err := tx.QueryRow(ctx,
			`INSERT INTO to_dos (id, message, details, due_date, priority, level_of_effort, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, message, details, due_date, priority, level_of_effort, created_at, updated_at, deleted_at;`,
			create.ID, create.Message, create.Details, create.DueDate, create.Priority, create.LevelOfEffort, create.CreatedAt, create.UpdatedAt).Scan(
			&returningToDo.ID,
			&returningToDo.Message,
			&returningToDo.Details,
			&returningToDo.DueDate,
			&returningToDo.Priority,
			&returningToDo.LevelOfEffort,
			&returningToDo.CreatedAt,
			&returningToDo.UpdatedAt,
			&returningToDo.DeletedAt,
		); err != nil {
			return returningToDos, errors.WithErrorMessage(err, errors.NotKnown, "inserting and scanning return todo")
		}

		returningToDoStatus := ToDoStatus{}
		if err := tx.QueryRow(ctx,
			`INSERT INTO to_dos_statuses (id, to_do_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, status, created_at, updated_at, deleted_at;`,
			ksuid.New().String(), returningToDo.ID, "created", now, now).Scan(
			&returningToDoStatus.ID,
			&returningToDoStatus.Status,
			&returningToDoStatus.CreatedAt,
			&returningToDoStatus.UpdatedAt,
			&returningToDoStatus.DeletedAt,
		); err != nil {
			return returningToDos, errors.WithErrorMessage(err, errors.NotKnown, "inserting status for created todo")
		}

		returningToDo.ToDoStatuses = append(returningToDo.ToDoStatuses, returningToDoStatus)

		returningToDos = append(returningToDos, returningToDo)
	}

	if err := tx.Commit(ctx); err != nil {
		return returningToDos, errors.WithErrorMessage(err, errors.NotKnown, "committing create ToDos query(s)")
	}

	return returningToDos, nil
}

// // ListTransactions get's transactions paginated by cursor and limit
// func (r *Repository) ListTransactions(ctx context.Context, transactionID string, createdAt time.Time, limit uint64) ([]Transaction, error) {
// 	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "repository.ListTransaction", trace.WithAttributes(
// 		attribute.String("cursor.id", transactionID),
// 		attribute.String("cursor.created_at", createdAt.String()),
// 		attribute.Int64("limit", int64(limit)),
// 	))
// 	defer span.End()

// 	tx, err := r.database.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return nil, errors.WithErrorMessage(err, errors.NotKnown, "beginning get transactions db transactoion")
// 	}

// 	query := `SELECT id, description, created_at
// 		FROM transaction
// 		WHERE 1=1`
// 	var params []interface{}
// 	if transactionID != "" && !createdAt.IsZero() {
// 		params = append(params, transactionID)
// 		query += fmt.Sprintf(" AND transaction.id <= $%d", len(params))
// 		params = append(params, createdAt)
// 		query += fmt.Sprintf(" AND created_at <= $%d", len(params))
// 	}
// 	query += " ORDER BY created_at DESC"
// 	if limit != 0 {
// 		params = append(params, limit)
// 		query += fmt.Sprintf(" LIMIT $%d", len(params))
// 	}
// 	query += ";"

// 	rows, err := tx.Query(ctx, query, params...)
// 	if err != nil {
// 		return nil, errors.WithErrorMessage(err, errors.NotKnown, "fetching transactions from database")
// 	}
// 	defer rows.Close()
// 	returning := []Transaction{}
// 	for rows.Next() {
// 		transaction := Transaction{}
// 		if err := rows.Scan(
// 			&transaction.ID,
// 			&transaction.Description,
// 			&transaction.CreatedAt,
// 		); err != nil {
// 			return nil, errors.WithErrorMessage(err, errors.NotKnown, "scanning transaction result row")
// 		}
// 		returning = append(returning, transaction)
// 	}

// 	for i, transaction := range returning {
// 		entries, err := getEntriesByTransactionID(ctx, tx, transaction.ID)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "getting entries by transaction id")
// 		}
// 		returning[i].Entries = append(returning[i].Entries, entries...)
// 	}

// 	if err := validate.Validate(returning); err != nil {
// 		return nil, errors.WithErrorMessage(err, errors.NotValidInternalData, "validating transactions fetched from database")
// 	}
// 	if err := tx.Commit(ctx); err != nil {
// 		return nil, errors.WithErrorMessage(err, errors.NotKnown, "committing get transactions transaction")
// 	}

// 	return returning, nil
// }
