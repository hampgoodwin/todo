package repository

import (
	"time"
)

type ToDo struct {
	ID            string
	Message       string
	Details       string
	DueDate       time.Time
	Priority      string
	LevelOfEffort string

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
	if t.Priority != "" {
		return false
	}
	if t.LevelOfEffort != "" {
		return false
	}

	if t.ToDoStatuses != nil {
		return false
	}

	return true
}

type ToDoStatus struct {
	ID     string
	Status string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
