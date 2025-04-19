package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ssantiago065/task_tracker/tasks"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  task-cli add <description>")
	fmt.Println("  task-cli delete <task ID>")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	filename := "tasks.json"

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: description is required.")
			printUsage()
			return
		}
		description := strings.Join(os.Args[2:], " ")
		err := tasks.AddTask(filename, description)
		if err != nil {
			fmt.Printf("Error adding task: %v", err)
			return
		}
		fmt.Println("Task added successfully.")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: task ID is required.")
			printUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: invalid task ID '%s'.", os.Args[2])
			return
		}
		err = tasks.DeleteTask(filename, id)
		if err != nil {
			fmt.Printf("Error deleting task: %v", err)
			return
		}
		fmt.Println("Task deleted successfully.")

	default:
		fmt.Printf("Unknown command: %s", command)
		printUsage()
	}
}
