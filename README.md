# 📝 Task Tracker CLI

**Task Tracker CLI** is a simple command-line tool written in Go for managing personal tasks. It allows you to create, view, update, and delete tasks, as well as mark them by status: `todo`, `in-progress`, or `done`.

## ⚙️ Installation

1. Make sure you have [Go](https://go.dev/doc/install) installed.
2. Clone the repository:

   git clone https://github.com/ssantiago065/task_tracker
   cd task_tracker

4. Build the CLI tool:

   go build -o task-cli

5. Use the tool:

   ./task-cli <command> [arguments]

---

## 🚀 Usage

The CLI stores your tasks in a local `tasks.json` file. Below is a breakdown of the available commands, along with examples.

### ➕ `add [description]`
Adds a new task with the given description.

./task-cli add "Finish Go project"


---

### 🗑️ `delete [task_id]`
Deletes the task with the given ID.

./task-cli delete 1


---

### ✏️ `update [task_id] [new description]`
Updates the description of the specified task.

./task-cli update 1 "Finish Go CLI project with filters"


---

### 🚧 `mark-in-progress [task_id]`
Marks a task as `in-progress`.

./task-cli mark-in-progress 2


---

### ✅ `mark-done [task_id]`
Marks a task as `done`.

./task-cli mark-done 3


---

### ⏳ `mark-todo [task_id]`
Resets a task’s status to `todo`.

./task-cli mark-todo 3


---

### 📋 `list`
Lists all tasks regardless of status.

./task-cli list


---

### 🔍 `list-done`
Lists only the tasks marked as `done`.

./task-cli list-done


---

### 🧮 `list-todo`
Lists only the tasks marked as `todo`.

./task-cli list-todo


---

### 🛠️ `list-in-progress`
Lists only the tasks marked as `in-progress`.

./task-cli list-in-progress

---

## 📌 Notes

- Tasks are stored locally in a `tasks.json` file.
- This tool is designed for local use and learning purposes.
- No external dependencies or databases required.

---

