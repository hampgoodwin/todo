package service

import (
	"context"
	"fmt"
	"time"

	"github.com/hampgoodwin/errors"
	eventv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/event/v1"
	"github.com/hampgoodwin/todo/internal/event"
	"github.com/hampgoodwin/todo/internal/meta"
	"github.com/hampgoodwin/todo/internal/todo"
	"github.com/hampgoodwin/todo/internal/transformer"
	"github.com/segmentio/ksuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type ListToDosReqest struct {
	IDs      []string
	Cursor   string
	PageSize int32
}

func (s *Service) ListToDos(ctx context.Context, request ListToDosReqest) ([]todo.ToDo, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "service.ListTransaction", trace.WithAttributes(
		attribute.String("cursor", request.Cursor),
		attribute.Int64("page_size", int64(request.PageSize)),
	))
	defer span.End()

	repoToDos, err := s.repository.ListToDos(ctx, request.IDs, request.Cursor, request.PageSize)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("fetchign transactions from database with cursor %q and limit '%d'", request.Cursor, request.PageSize))
	}

	toDos := transformer.NewToDosFromRepoToDos(repoToDos)

	return toDos, nil
}

func (s *Service) CreateToDos(ctx context.Context, creates []todo.ToDo) ([]todo.ToDo, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "internal.service.CreateToDos", trace.WithAttributes(
		attribute.Int("create_len", len(creates)),
	))
	defer span.End()

	for i := range creates {
		creates[i].ID = ksuid.New().String()
		creates[i].CreatedAt = time.Now()
	}

	repoCreates := transformer.NewRepoToDosFromToDos(creates)

	repoToDos, err := s.repository.CreateToDos(ctx, repoCreates)
	if err != nil {
		return nil, errors.WithErrorMessage(err, errors.NotValidRequestData, "validating transformed todos for data")
	}

	toDos := transformer.NewToDosFromRepoToDos(repoToDos)

	protoToDos := transformer.NewProtoToDosFromToDos(toDos)
	if err := s.publisher.Publish(event.SubjectToDosCreated, &eventv1.ToDosCreated{ToDos: protoToDos}); err != nil {
		s.log.Error("publishing todos created message", zap.Error(err))
	}

	return toDos, nil
}
