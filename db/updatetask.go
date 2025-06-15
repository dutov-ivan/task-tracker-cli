package db

import (
	"log"
	"time"

	"github.com/dutov-ivan/task-tracker-cli/models"
)

func updateTaskField(id int, updateFn func(task *models.Task)) {
	tasks := ReadAllTasks()
	task := FindTaskById(id, tasks)
	if task == nil {
		log.Fatal("No task with such ID is found. Try listing available tasks first")
	}
	updateFn(task)
	task.UpdatedAt = time.Now()
	SaveTasks(tasks)
}

func MarkStatus(id int, status models.TaskStatus) {
	updateTaskField(id, func(task *models.Task) {
		task.Status = status
	})
}

func Rename(id int, description string) {
	updateTaskField(id, func(task *models.Task) {
		task.Description = description
	})
}
