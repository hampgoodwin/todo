package transformer

import (
	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	"github.com/hampgoodwin/todo/internal/repository"
	"github.com/hampgoodwin/todo/internal/todo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewToDoStatusFromRepoToDoStatus(in repository.ToDoStatus) todo.ToDoStatus {
	if in == (repository.ToDoStatus{}) {
		return todo.ToDoStatus{}
	}

	out := todo.ToDoStatus{
		ID:        in.ID,
		Status:    todo.ParseStatus(in.Status),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}

	return out
}

func NewToDoStatusesFromRepoToDoStatuses(in []repository.ToDoStatus) []todo.ToDoStatus {
	out := []todo.ToDoStatus{}
	for _, v := range in {
		out = append(out, NewToDoStatusFromRepoToDoStatus(v))
	}
	return out
}

func NewRepoToDoStatusFromToDoStatus(in todo.ToDoStatus) repository.ToDoStatus {
	if in == (todo.ToDoStatus{}) {
		return repository.ToDoStatus{}
	}

	out := repository.ToDoStatus{
		ID:        in.ID,
		Status:    in.Status.String(),
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
		DeletedAt: in.DeletedAt,
	}

	return out
}

func NewRepoToDoStatusesFromToDoStatuses(in []todo.ToDoStatus) []repository.ToDoStatus {
	out := []repository.ToDoStatus{}
	for _, v := range in {
		out = append(out, NewRepoToDoStatusFromToDoStatus(v))
	}
	return out
}

func NewToDoStatusFromProtoToDoStatus(in *modelv1.ToDoStatus) todo.ToDoStatus {
	if in == nil {
		return todo.ToDoStatus{}
	}

	out := todo.ToDoStatus{
		ID:        in.Id,
		Status:    pbToDoStatusToToDoStatus(in.Status),
		CreatedAt: in.CreatedAt.AsTime(),
		UpdatedAt: in.UpdatedAt.AsTime(),
		DeletedAt: in.DeletedAt.AsTime(),
	}

	return out
}

func NewToDoStatusesFromProtoToDoStatuses(in []*modelv1.ToDoStatus) []todo.ToDoStatus {
	if in == nil {
		return nil
	}
	out := []todo.ToDoStatus{}
	for _, v := range in {
		out = append(out, NewToDoStatusFromProtoToDoStatus(v))
	}
	return out
}

func NewProtoToDoStatusFromToDoStatus(in todo.ToDoStatus) *modelv1.ToDoStatus {
	if in == (todo.ToDoStatus{}) {
		return nil
	}
	out := &modelv1.ToDoStatus{
		Id:     in.ID,
		Status: toDoStatusToPBToDoStatus(in.Status),

		CreatedAt: timestamppb.New(in.CreatedAt),
		UpdatedAt: timestamppb.New(in.UpdatedAt),
		DeletedAt: timestamppb.New(in.DeletedAt),
	}

	return out
}

func NewProtoToDoStatusesFromToDoStatuses(in []todo.ToDoStatus) []*modelv1.ToDoStatus {
	if in == nil {
		return nil
	}
	out := []*modelv1.ToDoStatus{}
	for _, v := range in {
		out = append(out, NewProtoToDoStatusFromToDoStatus(v))
	}
	return out
}
