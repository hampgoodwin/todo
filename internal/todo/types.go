package todo

import (
	"time"
)

type ToDo struct {
	ID            string `validate:"required,KSUID"`
	Message       string `validate:"required"`
	Details       string
	DueDate       time.Time
	Priority      Priority
	LevelOfEffort LevelOfEffort

	ToDoStatuses []ToDoStatus

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (t ToDo) IsZero() bool {
	if t.ID != "" {
		return false
	}
	if t.Message != "" {
		return false
	}
	if t.Details != "" {
		return false
	}
	if t.Priority != (Priority{}) {
		return false
	}
	if t.LevelOfEffort != (LevelOfEffort{}) {
		return false
	}

	if t.ToDoStatuses != nil {
		return false
	}

	return true
}

type ToDoStatus struct {
	ID     string
	Status Status

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//
// ENUM-Y TYPES
// We use something called "safer" enums, you can see more about it in the link below
// https://threedots.tech/post/safer-enums-in-go/

// Priority represents the priority of a todo
type Priority struct {
	// Slug should not be accessed by dependent code. It is exported for validation reasons
	Slug string `validate:"oneof=low medium high"`
}

var (
	PriorityUnspecified = Priority{""}
	PriorityLow         = Priority{"low"}
	PriorityMedium      = Priority{"medium"}
	PriorityHigh        = Priority{"high"}
)

// priorityAsStringMap is used in parsing a string to a type
var priorityAsStringMap = map[string]Priority{
	"":       PriorityUnspecified,
	"low":    PriorityLow,
	"medium": PriorityMedium,
	"high":   PriorityHigh,
}

func ParsePriority(t string) Priority {
	if v, ok := priorityAsStringMap[t]; ok {
		return v
	}
	return PriorityUnspecified
}

func (t Priority) String() string {
	return t.Slug
}

// LevelOfEffort represents the effort to complete a todo
type LevelOfEffort struct {
	// Slug should not be accessed by dependent code. It is exported for validation reasons
	Slug string `validate:"oneof=one two three four five"`
}

var (
	LevelOfEffortUnspecified = LevelOfEffort{""}
	LevelOfEffortOne         = LevelOfEffort{"one"}
	LevelOfEffortTwo         = LevelOfEffort{"two"}
	LevelOfEffortThree       = LevelOfEffort{"three"}
	LevelOfEffortFour        = LevelOfEffort{"four"}
	LevelOfEffortFive        = LevelOfEffort{"five"}
)

// levelofEffortAsStringMap is used in parsing a string to a type
var levelOfEffortAsStringMap = map[string]LevelOfEffort{
	"":      LevelOfEffortUnspecified,
	"one":   LevelOfEffortOne,
	"two":   LevelOfEffortTwo,
	"three": LevelOfEffortThree,
	"four":  LevelOfEffortFour,
	"five":  LevelOfEffortFive,
}

func ParseLevelOfEffort(t string) LevelOfEffort {
	if v, ok := levelOfEffortAsStringMap[t]; ok {
		return v
	}
	return LevelOfEffortUnspecified
}

func (t LevelOfEffort) String() string {
	return t.Slug
}

// Status
type Status struct {
	// Slug should not be accessed by dependent code. It is exported for validation reasons
	Slug string `validate:"oneof=created in_progress completed canceled"`
}

var (
	StatusUnspecified = Status{""}
	StatusCreated     = Status{"created"}
	StatusInProgress  = Status{"in_progress"}
	StatusCompleted   = Status{"completed"}
	StatusCanceled    = Status{"canceled"}
)

var statusAsStringMap = map[string]Status{
	"":            StatusUnspecified,
	"created":     StatusCreated,
	"in_progress": StatusInProgress,
	"completed":   StatusCompleted,
	"canceled":    StatusCanceled,
}

func ParseStatus(s string) Status {
	if v, ok := statusAsStringMap[s]; ok {
		return v
	}
	return StatusUnspecified
}

func (s Status) String() string {
	return s.Slug
}
