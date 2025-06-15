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
	Use:   "list",
	Short: "Lists all tasks",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ReadAllTasks()
		PrintTasks(tasks)
	},
}

var listTasksTodo = &cobra.Command{
	Use:   "todo",
	Short: "Lists tasks with status 'TODO'",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasksTodo := db.ListTasksFilteredByStatus(models.Todo)
		PrintTasks(tasksTodo)
	},
}

var listTasksInProgress = &cobra.Command{
	Use:   "in-progress",
	Short: "Lists tasks with status 'In progress'",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasksInProgress := db.ListTasksFilteredByStatus(models.InProgress)
		PrintTasks(tasksInProgress)
	},
}

var listTasksDone = &cobra.Command{
	Use:   "done",
	Short: "Lists tasks with status 'Done'",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasksDone := db.ListTasksFilteredByStatus(models.Done)
		PrintTasks(tasksDone)
	},
}
