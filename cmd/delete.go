package cmd

import (
	"fmt"
	"log"

	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteTask)
}

var deleteTask = &cobra.Command{
	Use:   "delete <id>",
	Short: "Deletes tasks with corresponding id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		task := db.DeleteTask(id)
		if task == nil {
			log.Fatal("No task with such id was found. Cannot delete")
		}

		fmt.Printf("Task %v was successfully deleted!\n", task.Description)
	},
}
