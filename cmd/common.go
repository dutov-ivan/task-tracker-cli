package cmd

import (
	"log"
	"strconv"
)

func processIdFromArgs(args []string) int {
	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal("Id must be a valid integer")
	}

	return id
}
