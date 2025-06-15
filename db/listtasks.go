package db

import "github.com/dutov-ivan/task-tracker-cli/models"

func ListTasksFilteredByStatus(status models.TaskStatus) []models.Task {
	tasks := ReadAllTasks()
	filteredTasks := []models.Task{}
	for i := range tasks {
		if tasks[i].Status == status {
			filteredTasks = append(filteredTasks, tasks[i])
		}
	}

	return filteredTasks
}
