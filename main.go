package main

import (
	"fmt"
	"os"

	"github.com/ssantiago065/task_tracker/tasks"
)

func main() {
	// Simple CLI: expect "add" command and a description
	if len(os.Args) < 3 || os.Args[1] != "add" {
		fmt.Println("Usage: task-cli add <description>")
		return
	}

	description := os.Args[2]
	filename := "tasks.json"

	err := tasks.AddTask(filename, description)
	if err != nil {
		fmt.Printf("Error adding task: %v\n", err)
		return
	}

	fmt.Println("Task added successfully")
}
