# Task Tracker

Task Tracker is a command-line interface (CLI) application for managing tasks. It allows users to add, update, delete, and list tasks with different statuses (to-do, in progress, completed). This project is designed to practice backend skills in Golang, including file handling, CLI argument processing, and JSON data manipulation.

## Features

- Add, update, and delete tasks.
- Mark tasks as "in progress" or "completed".
- List tasks by status.
- Store data in a JSON file.
- Uses only Go standard libraries.

## Requirements

- Go 1.22 or later

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ssantiago065/task-tracker.git
   cd task-tracker
   ```

2. Initialize the Go module (if not already created):
   ```sh
   go mod tidy
   ```

3. Build the application:
   ```sh
   go build -o task-cli
   ```

## Usage

Run the application from the terminal:

### Add a task
```sh
task-cli add "Buy groceries"
```
Expected output:
```
Task added successfully (ID: 1)
```

### Update a task
```sh
task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task
```sh
task-cli delete 1
```

### Change task status
```sh
task-cli mark-in-progress 1
task-cli mark-done 1
```

### List tasks
```sh
task-cli list
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Project Structure

```
task-tracker/
│── main.go           # Entry point
│── tasks.go          # Task management logic
│── storage.go        # JSON and file handling
│── go.mod            # Go module definition
│── go.sum            # Dependencies (if any)
│── tasks.json        # JSON file storing tasks
│── README.md         # Documentation
│── .gitignore        # Git ignored files
```
