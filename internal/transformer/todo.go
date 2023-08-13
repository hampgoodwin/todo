package transformer

import (
	"time"

	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	servicev1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/service/v1"
	"github.com/hampgoodwin/todo/internal/repository"
	"github.com/hampgoodwin/todo/internal/todo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewToDoFromRepoToDo(in repository.ToDo) todo.ToDo {
	if in.IsZero() {
		return todo.ToDo{}
	}

	outStatuses := NewToDoStatusesFromRepoToDoStatuses(in.ToDoStatuses)

	out := todo.ToDo{
		ID:            in.ID,
		Message:       in.Message,
		Details:       in.Details,
		DueDate:       in.DueDate,
		Priority:      todo.ParsePriority(in.Priority),
		LevelOfEffort: todo.ParseLevelOfEffort(in.LevelOfEffort),
		ToDoStatuses:  outStatuses,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
		DeletedAt:     in.DeletedAt,
	}

	return out
}

func NewToDosFromRepoToDos(in []repository.ToDo) []todo.ToDo {
	if in == nil {
		return nil
	}

	out := []todo.ToDo{}
	for _, v := range in {
		out = append(out, NewToDoFromRepoToDo(v))
	}

	return out
}

func NewRepoToDoFromToDo(in todo.ToDo) repository.ToDo {
	if in.IsZero() {
		return repository.ToDo{}
	}

	outStatuses := NewRepoToDoStatusesFromToDoStatuses(in.ToDoStatuses)

	out := repository.ToDo{
		ID:            in.ID,
		Message:       in.Message,
		Details:       in.Details,
		DueDate:       in.DueDate,
		Priority:      in.Priority.String(),
		LevelOfEffort: in.LevelOfEffort.String(),
		ToDoStatuses:  outStatuses,
		CreatedAt:     in.CreatedAt,
		UpdatedAt:     in.UpdatedAt,
		DeletedAt:     in.DeletedAt,
	}

	return out
}

func NewRepoToDosFromToDos(in []todo.ToDo) []repository.ToDo {
	if in == nil {
		return nil
	}

	out := []repository.ToDo{}

	for _, v := range in {
		out = append(out, NewRepoToDoFromToDo(v))
	}

	return out
}

func NewProtoToDoFromToDo(in todo.ToDo) *modelv1.ToDo {
	if in.IsZero() {
		return nil
	}

	var outDetails *string
	if in.Details == "" {
		outDetails = nil
	} else {
		outDetails = &in.Details
	}

	outStatuses := NewProtoToDoStatusesFromToDoStatuses(in.ToDoStatuses)

	priority := toDoPriorityToPBToDoPriority(in.Priority)
	levelOfEffort := toDoLevelOfEffortToPBToDoLevelOfEffort(in.LevelOfEffort)

	out := &modelv1.ToDo{
		Id:            in.ID,
		Message:       in.Message,
		Details:       outDetails,
		DueDate:       timestamppb.New(in.DueDate),
		Priority:      &priority,
		LevelOfEffort: &levelOfEffort,
		Statuses:      outStatuses,
		CreatedAt:     timestamppb.New(in.CreatedAt),
		UpdatedAt:     timestamppb.New(in.UpdatedAt),
		DeletedAt:     timestamppb.New(in.DeletedAt),
	}

	return out
}

func NewProtoToDosFromToDos(in []todo.ToDo) []*modelv1.ToDo {
	if in == nil {
		return nil
	}

	out := []*modelv1.ToDo{}
	for _, v := range in {
		out = append(out, NewProtoToDoFromToDo(v))
	}

	return out
}

func NewToDoFromProtoToDo(in *modelv1.ToDo) todo.ToDo {
	if in == nil {
		return todo.ToDo{}
	}

	var outDetails string
	if in.Details != nil {
		outDetails = *in.Details
	}
	var outDueDate time.Time
	if in.DueDate != nil {
		outDueDate = in.DueDate.AsTime()
	}

	var outCreatedAt time.Time
	if in.CreatedAt != nil {
		outCreatedAt = in.CreatedAt.AsTime()
	}
	var outUpdatedAt time.Time
	if in.UpdatedAt != nil {
		outUpdatedAt = in.UpdatedAt.AsTime()
	}
	var outDeletedAt time.Time
	if in.UpdatedAt != nil {
		outDeletedAt = in.DeletedAt.AsTime()
	}

	outStatuses := NewToDoStatusesFromProtoToDoStatuses(in.Statuses)

	out := todo.ToDo{
		ID:            in.Id,
		Message:       in.Message,
		Details:       outDetails,
		DueDate:       outDueDate,
		Priority:      pbToDoPriorityTypeToToDoPriority(in.Priority),
		LevelOfEffort: pbToDoLevelOfEffortToToDoLevelOfEffort(in.LevelOfEffort),
		ToDoStatuses:  outStatuses,
		CreatedAt:     outCreatedAt,
		UpdatedAt:     outUpdatedAt,
		DeletedAt:     outDeletedAt,
	}

	return out
}

func NewToDoFromProtoCreateToDo(in *servicev1.CreateToDos) todo.ToDo {
	if in == nil {
		return todo.ToDo{}
	}

	var outDetails string
	if in.Details != nil {
		outDetails = *in.Details
	}
	var outDueDate time.Time
	if in.DueDate != nil {
		outDueDate = in.DueDate.AsTime()
	}

	out := todo.ToDo{
		Message:       in.Message,
		Details:       outDetails,
		DueDate:       outDueDate,
		Priority:      pbToDoPriorityTypeToToDoPriority(in.Priority),
		LevelOfEffort: pbToDoLevelOfEffortToToDoLevelOfEffort(in.LevelOfEffort),
	}

	return out
}

func NewToDosFromProtoCreateToDos(in []*servicev1.CreateToDos) []todo.ToDo {
	if in == nil {
		return nil
	}

	out := []todo.ToDo{}
	for _, v := range in {
		out = append(out, NewToDoFromProtoCreateToDo(v))
	}

	return out
}
