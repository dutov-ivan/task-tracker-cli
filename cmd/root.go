package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "Task CLI is a CLI-tool for basic task management.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You got me! You can look for options under -h flag!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
