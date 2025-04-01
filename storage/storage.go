package storage

import (
	"encoding/json"
	"os"

	"github.com/ssantiago065/task_tracker/tasks"
)

func LoadTasks(filename string) ([]tasks.Task, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) { // If file does not exist, return empty list
			return []tasks.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []tasks.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(filename string, tasks []tasks.Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}
