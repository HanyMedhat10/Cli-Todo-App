# Todo CLI Application

A simple command-line interface application built with Go for managing your todo tasks. This application allows you to create, list, edit, delete, and toggle the completion status of your todos.

## Features

- Add new todo items
- List all todo items in a formatted table
- Edit existing todo items
- Delete todo items
- Toggle completion status of todo items
- Persistent storage using JSON

## Installation

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository:
   ```
   git clone <repository-url>
   cd cli-todo-app
   ```

2. Build the application:
   ```
   go build
   ```

## Usage

The application supports the following commands:

| Command | Description |
|---------|-------------|
| `-add` | Add a new todo item |
| `-list` | List all todo items |
| `-edit` | Edit an existing todo item |
| `-del` | Delete a todo item |
| `-toggle` | Toggle the completion status of a todo item |
| `-help` | Display help information |

### Examples

#### Add a new todo
```
./cli-todo-app -add="Learn Go programming"
```

#### List all todos
```
./cli-todo-app -list
```

#### Edit a todo (format: id:new_title)
```
./cli-todo-app -edit="0:Complete Go project"
```

#### Delete a todo
```
./cli-todo-app -del=1
```

#### Toggle completion status
```
./cli-todo-app -toggle=0
```

#### Display help
```
./cli-todo-app -help
```

## Data Storage

Todo items are stored in a JSON file (`todos.json`) in the application directory. The application automatically loads and saves todo items from/to this file.

## Project Structure

- `main.go`: Entry point of the application
- `todo.go`: Contains the Todo struct and methods for managing todos
- `command.go`: Handles command-line arguments and execution
- `storage.go`: Manages data persistence using JSON

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
