package db

import (
	"slices"

	"github.com/dutov-ivan/task-tracker-cli/models"
)

func DeleteTask(id int) *models.Task {
	tasks := ReadAllTasks()
	for i := range tasks {
		if tasks[i].Id == id {
			deletedTask := tasks[i]
			newTasks := slices.Delete(tasks, i, i+1)
			SaveTasks(newTasks)
			return &deletedTask
		}
	}

	return nil
}
