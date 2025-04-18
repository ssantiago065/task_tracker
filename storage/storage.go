package storage

import (
	"encoding/json"
	"os"

	"github.com/ssantiago065/task_tracker/tasks"
)

func LoadTasks(filename string) (tasks.TaskStore, error) {
	var store tasks.TaskStore

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			store = tasks.TaskStore{LastID: 0, Tasks: []tasks.Task{}}
			return store, nil
		}
		return store, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&store)
	if err != nil {
		return store, err
	}

	return store, nil
}

func SaveTasks(filename string, store tasks.TaskStore) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&store)
	if err != nil {
		return err
	}

	return nil
}
