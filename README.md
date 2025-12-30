# Task Tracker CLI

Task Tracker CLI is a simple command-line based task management application written in Go (Golang). This application allows users to create, read, update, and delete tasks (CRUD) and track their status. Task data is persistently stored in JSON format.

## Features

*   **Add New Task**: Create a new task with a description and an auto-generated ID.
*   **Edit Task**: Modify the description or status of an existing task.
*   **Delete Task**: Remove a task based on its unique ID.
*   **Change Status**: Update a task's status (e.g., Pending, On Progress, Done).
*   **Show All Tasks**: Display a list of all stored tasks.
*   **Filter Tasks**:
    *   Show completed tasks (*Done*).
    *   Show tasks currently being worked on (*On Progress*).
    *   Show incomplete tasks (*Pending*).
*   **Persistent Storage**: All data is saved in the `tasks.json` file, ensuring data is not lost when the program is closed.

## Prerequisites

Before running this application, ensure you have installed:

*   **Go**: Version 1.25.5 or later.

## Installation and execution

1.  **Clone Repository** (if using git) or download the source code.
    ```bash
    git clone <repository-url>
    cd task-tracker
    ```

2.  **Run the Application**
    You can run the application directly using the `go run` command:
    ```bash
    go run main.go
    ```

    Or build it into an executable binary first:
    ```bash
    go build -o task-tracker
    ./task-tracker     # For Linux/Mac
    task-tracker.exe   # For Windows
    ```

## Usage Guide

After running the program, you will be presented with an interactive menu as follows:

```text
========================================
              TASK TRACKER
========================================
Menu:
1. Add new task
2. Edit existing task
3. Delete existing task
4. Change task's status
5. Show all tasks
6. Show all done tasks
7. Show all on progress tasks
8. Show all not done tasks
9. Close program
Enter your choice (1 - 9):
```

### Menu Details:

1.  **Add new task**:
    *   Select this option to add a task.
    *   Enter the task description when prompted.
    *   The task will be automatically created with a `Pending` status and a unique ID (e.g., T0001).

2.  **Edit existing task**:
    *   Enter the **UID** (Unique ID) of the task you want to edit.
    *   Choose whether to edit the **Description** or **Status**.

3.  **Delete existing task**:
    *   Enter the **UID** of the task you want to delete.
    *   Confirm deletion by selecting `Yes` or `No`.

4.  **Change task's status**:
    *   Enter the **UID** of the task.
    *   Type the new desired status (e.g., "Done", "On Progress").

5.  **Show all tasks**:
    *   Displays the entire list of tasks along with their details (UID, Description, Status, Created At, Updated At).

6.  **Show all done tasks**:
    *   Only displays tasks with the status "Done".

7.  **Show all on progress tasks**:
    *   Only displays tasks with the status "On Progress".

8.  **Show all not done tasks**:
    *   Only displays tasks with the status "Pending".

9.  **Close program**:
    *   Exits the application.

## Data Structure

Data is stored in the `tasks.json` file with the following structure:

```json
[
 {
  "id": "T0001",
  "desc": "Learn Golang",
  "status": "Pending",
  "created_at": "31-12-25 10:00:00",
  "updated_at": "31-12-25 10:00:00"
 }
]
```

*   **id**: Unique task identification (Auto-increment).
*   **desc**: Task description.
*   **status**: Task completion status.
*   **created_at**: Task creation time.
*   **updated_at**: Last time the task was modified.

## Project Structure

```
task-tracker/
├── main.go       # Main application code
├── go.mod        # Go module definition
├── tasks.json    # Database file (automatically created/updated at runtime)
└── README.md     # Project documentation
```

## Development

If you wish to contribute or modify the code:

1.  Open the `main.go` file.
2.  The `Task` struct defines the data shape.
3.  The `process(menu int)` function handles menu branching logic.
4.  The `loadTasks()` and `saveFile()` functions handle JSON file operations.

## License

[MIT License](https://opensource.org/licenses/MIT) - Feel free to use and modify as needed.
