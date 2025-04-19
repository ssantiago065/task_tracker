package storage

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "todo", "in-progress", "done"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskStore struct {
	LastID int    `json:"last_id"`
	Tasks  []Task `json:"tasks"`
}

func LoadTasks(filename string) (TaskStore, error) {
	var store TaskStore

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			store = TaskStore{LastID: 0, Tasks: []Task{}}
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

func SaveTasks(filename string, store TaskStore) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(&store)
	if err != nil {
		return err
	}

	return nil
}
