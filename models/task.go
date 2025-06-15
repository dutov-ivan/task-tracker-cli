package models

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	Todo       TaskStatus = "TODO"
	InProgress TaskStatus = "In Progress"
	Done       TaskStatus = "Done"
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

func (task *Task) Print() {
	fmt.Printf("ID: %v\n", task.Id)
	fmt.Printf("Description: %v\n", task.Description)
	fmt.Printf("Status: %v\n", task.Status)
	fmt.Printf("Created at: %v\n", task.CreatedAt.Local().Format("3:04PM, 02 Jan 2006"))
	fmt.Printf("Updated at: %v\n", task.UpdatedAt.Local().Format("3:04PM, 02 Jan 2006"))
}
