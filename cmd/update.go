package cmd

import (
	"github.com/dutov-ivan/task-tracker-cli/db"
	"github.com/dutov-ivan/task-tracker-cli/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markInProgress, markDone, changeDescription)
}

var markInProgress = &cobra.Command{
	Use:   "mark-in-progress <id>",
	Short: "Sets status to 'In progress' for task with specified ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		db.MarkStatus(id, models.InProgress)
	},
}

var markDone = &cobra.Command{
	Use:   "mark-done <id>",
	Short: "Sets status to 'Done' for task with specified ID.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		db.MarkStatus(id, models.Done)
	},
}

var changeDescription = &cobra.Command{
	Use:   "update <id> <description>",
	Short: "Changes description for task with specified ID",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := processIdFromArgs(args)
		description := args[1]
		db.Rename(id, description)
	},
}
