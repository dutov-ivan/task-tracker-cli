package cmd

import (
	"fmt"

	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listTasks)
}

var listTasks = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ReadAllTasks()
		for _, task := range *tasks {
			task.Print()
			fmt.Print("\n")
		}
	},
}
