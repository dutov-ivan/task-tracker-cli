package cmd

import (
	"log"
	"strconv"

	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/dutov-ivan/task-tracker-cli/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markInProgress)
	rootCmd.AddCommand(markDone)
	rootCmd.AddCommand(changeDescription)
}

func processIdFromArgs(args []string) int {
	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal("Id must be a valid integer")
	}

	return id
}

var markInProgress = &cobra.Command{
	Use:  "mark-in-progress <id>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		db.MarkStatus(id, models.InProgress)
	},
}

var markDone = &cobra.Command{
	Use:  "mark-done <id>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		db.MarkStatus(id, models.Done)
	},
}

var changeDescription = &cobra.Command{
	Use:  "update <id> <description>",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		description := args[1]
		db.Rename(id, description)
	},
}
