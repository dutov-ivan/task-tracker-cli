package cmd

import (
	"fmt"

	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/dutov-ivan/task-tracker-cli/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTasks)
	listTasks.AddCommand(listTasksTodo)
	listTasks.AddCommand(listTasksInProgress)
	listTasks.AddCommand(listTasksDone)

}

func PrintTasks(tasks []models.Task) {
	for _, task := range tasks {
		task.Print()
		fmt.Print("\n")
	}
}

var listTasks = &cobra.Command{
	Use:  "list",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ReadAllTasks()
		PrintTasks(tasks)
	},
}

var listTasksTodo = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		tasksTodo := db.ListTasksFilteredByStatus(models.Todo)
		PrintTasks(tasksTodo)
	},
}

var listTasksInProgress = &cobra.Command{
	Use: "in-progress",
	Run: func(cmd *cobra.Command, args []string) {
		tasksInProgress := db.ListTasksFilteredByStatus(models.InProgress)
		PrintTasks(tasksInProgress)
	},
}

var listTasksDone = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		tasksDone := db.ListTasksFilteredByStatus(models.Done)
		PrintTasks(tasksDone)
	},
}
