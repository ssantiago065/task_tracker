package tasks

import (
	"fmt"
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
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	deleted := false
	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks = append(store.Tasks[:index], store.Tasks[index+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		return fmt.Errorf("task %d not found", id)
	}

	return storage.SaveTasks(filename, store)
}

func UpdateTask(filename string, id int, description string) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	updated := false
	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[index].Description = description
			store.Tasks[index].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("task %d not found", id)
	}

	return storage.SaveTasks(filename, store)
}

func MarkInProgress(filename string, id int) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	marked := false
	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[index].Status = "in-progress"
			store.Tasks[index].UpdatedAt = time.Now()
			marked = true
			break
		}
	}

	if !marked {
		return fmt.Errorf("task %d not found", id)
	}

	return storage.SaveTasks(filename, store)
}

func MarkDone(filename string, id int) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	marked := false
	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[index].Status = "done"
			store.Tasks[index].UpdatedAt = time.Now()
			marked = true
			break
		}
	}

	if !marked {
		return fmt.Errorf("task %d not found", id)
	}

	return storage.SaveTasks(filename, store)
}

func MarkTodo(filename string, id int) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	marked := false
	for index, task := range store.Tasks {
		if task.ID == id {
			store.Tasks[index].Status = "todo"
			store.Tasks[index].UpdatedAt = time.Now()
			marked = true
			break
		}
	}

	if !marked {
		return fmt.Errorf("task %d not found", id)
	}

	return storage.SaveTasks(filename, store)
}

func ListTasks(filename string) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	if len(store.Tasks) == 0 {
		return fmt.Errorf("no tasks found")
	}

	for _, task := range store.Tasks {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
			task.ID,
			task.Description,
			task.Status,
			task.CreatedAt.Format(time.ANSIC),
			task.UpdatedAt.Format(time.ANSIC))
	}

	return nil
}

func ListDone(filename string) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	if len(store.Tasks) == 0 {
		return fmt.Errorf("no tasks found")
	}

	for _, task := range store.Tasks {
		if task.Status == "done" {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.ANSIC),
				task.UpdatedAt.Format(time.ANSIC))
		}
	}

	return nil
}

func ListTodo(filename string) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	if len(store.Tasks) == 0 {
		return fmt.Errorf("no tasks found")
	}

	for _, task := range store.Tasks {
		if task.Status == "todo" {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.ANSIC),
				task.UpdatedAt.Format(time.ANSIC))
		}
	}

	return nil
}

func ListInProgress(filename string) error {
	store, err := storage.LoadTasks(filename)
	if err != nil {
		return err
	}

	if len(store.Tasks) == 0 {
		return fmt.Errorf("no tasks found")
	}

	for _, task := range store.Tasks {
		if task.Status == "in-progress" {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.ID,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.ANSIC),
				task.UpdatedAt.Format(time.ANSIC))
		}
	}

	return nil
}
