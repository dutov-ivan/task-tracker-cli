package models

type TaskState string

const (
	Todo       TaskState = "TODO"
	InProgress TaskState = "In progress"
	Completed  TaskState = "Completed"
)

type Task struct {
	Id    int
	Title string
	State TaskState
}

func NewTask(title string) Task {
	task := Task{
		Title: title,
		State: Todo,
	}

	return task
}
