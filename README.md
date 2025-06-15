# Task CLI Overview

Task CLI is a simple command-line tool for basic task management, designed to help you track your tasks efficiently from the terminal. It allows you to add, list, update, and delete tasks, as well as manage their statuses (TODO, In Progress, Done).

## Features
- **Add Tasks:** Quickly add new tasks with a description.
- **List Tasks:** View all tasks or filter by status (TODO, In Progress, Done).
- **Update Tasks:** Change the description or status of existing tasks.
- **Delete Tasks:** Remove tasks by their ID.
- **Persistent Storage:** Tasks are stored in a local JSON file (`db.json`).

## Usage
Run `task-cli --help` to see all available commands and options.

Example commands:
- `task-cli add "Buy groceries"` — Add a new task
- `task-cli list` — List all tasks
- `task-cli list todo` — List only TODO tasks
- `task-cli mark-in-progress <id>` — Mark a task as In Progress
- `task-cli mark-done <id>` — Mark a task as Done
- `task-cli update <id> "New description"` — Update a task's description
- `task-cli delete <id>` — Delete a task

