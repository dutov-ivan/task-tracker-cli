package db

import (
	"github.com/dutov-ivan/task-tracker-cli/models"
)

func AddTask(newTask *models.Task) {
	tasks := ReadAllTasks()

	AssignNextHighestId(tasks, newTask)

	*tasks = append(*tasks, *newTask)

	SaveTasks(tasks)
}

func AssignNextHighestId(tasks *[]models.Task, newTask *models.Task) {
	maxId := 0
	for _, task := range *tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	newTask.Id = maxId + 1
}
