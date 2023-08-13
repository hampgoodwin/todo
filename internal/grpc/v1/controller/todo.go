package controller

import (
	"context"

	"github.com/hampgoodwin/errors"
	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	servicev1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/service/v1"
	"github.com/hampgoodwin/todo/internal/grpc/pagination/v1"
	"github.com/hampgoodwin/todo/internal/meta"
	"github.com/hampgoodwin/todo/internal/service"
	"github.com/hampgoodwin/todo/internal/transformer"
	"github.com/hampgoodwin/todo/internal/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const (
	maxPageSize     = 100
	defaultPageSize = 10
)

func (c *Controller) CreateToDos(ctx context.Context, req *servicev1.CreateToDosRequest) (*servicev1.CreateToDosResponse, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "internal.grpc.v1.controller.CreateAccount", trace.WithAttributes(
		attribute.Int("to_dos_len", len(req.GetCreateToDos())),
	))
	defer span.End()

	if err := validate.Validate(req); err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotValidRequest, "validating request"))
	}

	createToDos := transformer.NewToDosFromProtoCreateToDos(req.CreateToDos)

	toDos, err := c.service.CreateToDos(ctx, createToDos)
	if err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotKnown, "creating to dos"))
	}

	if err := validate.Validate(toDos); err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotValidInternalData, "validating created todos"))
	}

	protoTodos := transformer.NewProtoToDosFromToDos(toDos)

	response := &servicev1.CreateToDosResponse{
		ToDos: protoTodos,
	}

	return response, nil
}

func (c *Controller) ListToDos(ctx context.Context, req *servicev1.ListToDosRequest) (*servicev1.ListToDosResponse, error) {
	ctx, span := otel.Tracer(meta.ServiceName).Start(ctx, "internal.grpc.v1.controller.ListToDos", trace.WithAttributes(
		attribute.StringSlice("ids", req.GetIds()),
		attribute.Int64("page_size", int64(req.GetPageSize())),
		attribute.String("page_token", req.GetPageToken()),
	))
	defer span.End()

	if err := validate.Validate(req); err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotValidRequest, "validating list todos request"))
	}

	pageSize := req.GetPageSize()
	if req.GetPageSize() == 0 {
		pageSize = defaultPageSize
	}
	if req.GetPageSize() > maxPageSize {
		pageSize = maxPageSize
	}

	pageToken, err := pagination.ParsePageToken(req)
	if err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotKnown, "parsing page token"))
	}

	todos, err := c.service.ListToDos(ctx, service.ListToDosReqest{
		IDs:      req.GetIds(),
		Cursor:   pageToken.LastID,
		PageSize: pageSize + 1, // get one additional to determine if there is a next page
	})
	if err != nil {
		return nil, c.respondError(ctx, c.log, errors.WithErrorMessage(err, errors.NotKnown, "getting to dos"))
	}

	var nextPageToken string
	protoTodos := []*modelv1.ToDo{}
	if int32(len(todos)) > pageSize {
		nextPageToken = pageToken.Next(todos[len(todos)-1].ID).String()
		// we asked for one more todo than requested in order to handle pagination completely grpc interface side
		protoTodos = transformer.NewProtoToDosFromToDos(todos[:len(todos)-2]) // take all but
	} else {
		// Not enough records to meet the request pagesize
		protoTodos = transformer.NewProtoToDosFromToDos(todos) // take all but
	}

	response := &servicev1.ListToDosResponse{
		ToDos:         protoTodos,
		NextPageToken: nextPageToken,
	}

	return response, nil
}
