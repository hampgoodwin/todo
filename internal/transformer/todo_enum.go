package transformer

import (
	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	"github.com/hampgoodwin/todo/internal/todo"
)

var toDoPriorityToPBToDoPriorityMap = map[todo.Priority]modelv1.Priority{
	todo.PriorityUnspecified: modelv1.Priority_PRIORITY_UNSPECIFIED,
	todo.PriorityLow:         modelv1.Priority_PRIORITY_LOW,
	todo.PriorityMedium:      modelv1.Priority_PRIORITY_MEDIUM,
	todo.PriorityHigh:        modelv1.Priority_PRIORITY_HIGH,
}

func toDoPriorityToPBToDoPriority(in todo.Priority) modelv1.Priority {
	if v, ok := toDoPriorityToPBToDoPriorityMap[in]; ok {
		return v
	}
	return modelv1.Priority_PRIORITY_UNSPECIFIED
}

var pbToDoPriorityTypeToToDoPriorityMap = map[modelv1.Priority]todo.Priority{
	modelv1.Priority_PRIORITY_UNSPECIFIED: todo.PriorityUnspecified,
	modelv1.Priority_PRIORITY_LOW:         todo.PriorityLow,
	modelv1.Priority_PRIORITY_MEDIUM:      todo.PriorityMedium,
	modelv1.Priority_PRIORITY_HIGH:        todo.PriorityHigh,
}

func pbToDoPriorityTypeToToDoPriority(in *modelv1.Priority) todo.Priority {
	if in == nil {
		return todo.PriorityUnspecified
	}

	if v, ok := pbToDoPriorityTypeToToDoPriorityMap[*in]; ok {
		return v
	}

	return todo.PriorityUnspecified
}

var toDoLevelOfEffortToPBToDoLevelOfEffortMap = map[todo.LevelOfEffort]modelv1.LevelOfEffort{
	todo.LevelOfEffortUnspecified: modelv1.LevelOfEffort_LEVEL_OF_EFFORT_UNSPECIFIED,
	todo.LevelOfEffortOne:         modelv1.LevelOfEffort_LEVEL_OF_EFFORT_ONE,
	todo.LevelOfEffortTwo:         modelv1.LevelOfEffort_LEVEL_OF_EFFORT_TWO,
	todo.LevelOfEffortThree:       modelv1.LevelOfEffort_LEVEL_OF_EFFORT_THREE,
	todo.LevelOfEffortFour:        modelv1.LevelOfEffort_LEVEL_OF_EFFORT_FOUR,
	todo.LevelOfEffortFive:        modelv1.LevelOfEffort_LEVEL_OF_EFFORT_FIVE,
}

func toDoLevelOfEffortToPBToDoLevelOfEffort(in todo.LevelOfEffort) modelv1.LevelOfEffort {
	if v, ok := toDoLevelOfEffortToPBToDoLevelOfEffortMap[in]; ok {
		return v
	}
	return modelv1.LevelOfEffort_LEVEL_OF_EFFORT_UNSPECIFIED
}

var pbToDoLevelOfEffortToToDoLevelOfEffortMap = map[modelv1.LevelOfEffort]todo.LevelOfEffort{
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_UNSPECIFIED: todo.LevelOfEffortUnspecified,
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_ONE:         todo.LevelOfEffortOne,
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_TWO:         todo.LevelOfEffortTwo,
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_THREE:       todo.LevelOfEffortThree,
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_FOUR:        todo.LevelOfEffortFour,
	modelv1.LevelOfEffort_LEVEL_OF_EFFORT_FIVE:        todo.LevelOfEffortFive,
}

func pbToDoLevelOfEffortToToDoLevelOfEffort(in *modelv1.LevelOfEffort) todo.LevelOfEffort {
	if in == nil {
		return todo.LevelOfEffortUnspecified
	}
	if v, ok := pbToDoLevelOfEffortToToDoLevelOfEffortMap[*in]; ok {
		return v
	}
	return todo.LevelOfEffortUnspecified
}

var toDoStatusToPBToDoStatusMap = map[todo.Status]modelv1.Status{
	todo.StatusUnspecified: modelv1.Status_STATUS_UNSPECIFIED,
	todo.StatusCreated:     modelv1.Status_STATUS_CREATED,
	todo.StatusInProgress:  modelv1.Status_STATUS_IN_PROGRESS,
	todo.StatusCompleted:   modelv1.Status_STATUS_COMPLETED,
	todo.StatusCanceled:    modelv1.Status_STATUS_CANCELED,
}

func toDoStatusToPBToDoStatus(in todo.Status) modelv1.Status {
	if v, ok := toDoStatusToPBToDoStatusMap[in]; ok {
		return v
	}
	return modelv1.Status_STATUS_UNSPECIFIED
}

var pbToDoStatusToToDoStatusMap = map[modelv1.Status]todo.Status{
	modelv1.Status_STATUS_UNSPECIFIED: todo.StatusUnspecified,
	modelv1.Status_STATUS_CREATED:     todo.StatusCreated,
	modelv1.Status_STATUS_IN_PROGRESS: todo.StatusInProgress,
	modelv1.Status_STATUS_COMPLETED:   todo.StatusCompleted,
	modelv1.Status_STATUS_CANCELED:    todo.StatusCanceled,
}

func pbToDoStatusToToDoStatus(in modelv1.Status) todo.Status {
	if v, ok := pbToDoStatusToToDoStatusMap[in]; ok {
		return v
	}
	return todo.StatusUnspecified
}
