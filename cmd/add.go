package cmd

import (
	"fmt"

	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/dutov-ivan/task-tracker-cli/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addTask)
}

var addTask = &cobra.Command{
	Use: "add <task>",
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		task := models.NewTask(title)
		db.AddTask(&task)

		fmt.Printf("Task '%v' successfully created!", task.Title)
	},
}
