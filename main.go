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
	fmt.Println("  task-cli update <task_ID> <description>")
	fmt.Println("  task-cli mark-in-progress <task_ID>")
	fmt.Println("  task-cli mark-done <task_ID>")
	fmt.Println("  task-cli mark-todo <task_ID>")
	fmt.Println("  task-cli list")
	fmt.Println("  task-cli list-done")
	fmt.Println("  task-cli list-todo")

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

	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Error: task ID is required.")
			printUsage()
			return
		}
		if len(os.Args) < 4 {
			fmt.Println("Error: new description is required")
			printUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: invalid task ID '%s'.", os.Args[2])
			return
		}
		description := strings.Join(os.Args[3:], " ")
		err = tasks.UpdateTask(filename, id, description)
		if err != nil {
			fmt.Printf("Error updating task: %v", err)
			return
		}
		fmt.Println("Task updated successfully.")

	case "mark-in-progress":
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
		err = tasks.MarkInProgress(filename, id)
		if err != nil {
			fmt.Printf("Error updating task: %v", err)
			return
		}
		fmt.Println("Task marked in progress successfully.")

	case "mark-done":
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
		err = tasks.MarkDone(filename, id)
		if err != nil {
			fmt.Printf("Error updating task: %v", err)
			return
		}
		fmt.Println("Task marked done successfully.")

	case "mark-todo":
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
		err = tasks.MarkTodo(filename, id)
		if err != nil {
			fmt.Printf("Error updating task: %v", err)
			return
		}
		fmt.Println("Task marked todo successfully.")

	case "list":
		err := tasks.ListTasks(filename)
		if err != nil {
			fmt.Printf("Error listing tasks: %v", err)
			return
		}
		fmt.Println("Tasks listed successfully.")

	case "list-done":
		err := tasks.ListDone(filename)
		if err != nil {
			fmt.Printf("Error listing tasks: %v", err)
			return
		}
		fmt.Println("Done tasks listed successfully.")

	case "list-todo":
		err := tasks.ListDone(filename)
		if err != nil {
			fmt.Printf("Error listing tasks: %v", err)
			return
		}
		fmt.Println("Todo tasks listed successfully.")

	default:
		fmt.Printf("Unknown command: %s", command)
		printUsage()
	}
}
