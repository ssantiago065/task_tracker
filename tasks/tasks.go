package tasks

import (
	"time"

	"github.com/ssantiago065/task_tracker/storage"
)

func AddTask(filename string, description string) error {
	var store storage.TaskStore

	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	store.LastID++
	now := time.Now()
	newTask := storage.Task{
		ID:          store.LastID,
		Description: description,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	store.Tasks = append(store.Tasks, newTask)

	err = storage.SaveTasks(filename, store)
	return err
}

func DeleteTask(filename string, id int) error {
	var store storage.TaskStore

	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks = append(store.Tasks[:index], store.Tasks[index+1:]...)
		}
	}

	err = storage.SaveTasks(filename, store)
	return err
}
