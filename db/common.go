package db

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/dutov-ivan/task-tracker-cli/models"
)

const dbFile string = "db.json"
const defaultPermissions os.FileMode = 0644

func OpenDbOrCreateOne(flag int) *os.File {
	if FileExists(dbFile) {
		f, _ := os.OpenFile(dbFile, flag, defaultPermissions)
		return f
	}

	f, err := os.Create(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
func SaveTasks(tasks []models.Task) {
	contents, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	db := OpenDbOrCreateOne(os.O_WRONLY)
	defer db.Close()

	_, err = db.Write(contents)
	if err != nil {
		log.Fatal(err)
	}

}

func ReadAllTasks() []models.Task {
	db := OpenDbOrCreateOne(os.O_RDONLY)
	defer db.Close()

	contents, err := io.ReadAll(db)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []models.Task
	json.Unmarshal(contents, &tasks)

	return tasks
}

func FindTaskById(id int, tasks []models.Task) *models.Task {
	for i := range tasks {
		if tasks[i].Id == id {
			return &tasks[i]
		}
	}

	return nil
}
