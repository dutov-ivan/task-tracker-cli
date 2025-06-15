package models

import "time"

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "In progress"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(title string) Task {
	now := time.Now()
	task := Task{
		Description: title,
		Status:      Todo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	return task
}
