# Go CLI Todo App

A simple and lightweight Command Line Todo Application built with Go (Golang).  
Manage your daily tasks directly from the terminal using command flags.

---

# Features

✅ Add new todos  
✅ List all todos  
✅ Mark todos as completed  
✅ Delete todos  
✅ Toggle todo completion status  
✅ Edit existing todos  
✅ Simple CLI interface using Go flags

---

# Technologies Used

- Go (Golang)
- Go `flag` package
- Structs & Methods
- Slices
- Command Line Interface (CLI)

---

# Installation

## 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-cli-todo.git
```

---

## 2. Navigate to Project Folder

```bash
cd go-cli-todo
```

---

## 3. Run the Application

```bash
go run .
```

---

# Command Flags

| Flag | Type | Description | Example |
|---|---|---|---|
| `-add` | string | Add a new todo | `go run . -add "Learn Go"` |
| `-list` | bool | List all todos | `go run . -list` |
| `-complete` | int | Mark todo as completed by index | `go run . -complete 0` |
| `-delete` | int | Delete todo by index | `go run . -delete 1` |
| `-toggle` | int | Toggle completion status | `go run . -toggle 0` |
| `-edit` | string | Edit todo title | `go run . -edit "0:Learn Advanced Go"` |

---

# Usage Guide

---

# Add a Todo

```bash
go run . -add "Learn Golang"
```

### Output

```bash
Todo added successfully
```

---

# List All Todos

```bash
go run . -list
```

### Example Output

```bash
0. [ ] Learn Golang
1. [x] Build CLI Todo App
```

### Symbols Meaning

| Symbol | Meaning |
|---|---|
| `[ ]` | Incomplete Todo |
| `[x]` | Completed Todo |

---

# Complete a Todo

Mark a todo as completed using its index.

```bash
go run . -complete 0
```

### Before

```bash
0. [ ] Learn Golang
```

### After

```bash
0. [x] Learn Golang
```

---

# Delete a Todo

Delete a todo using its index.

```bash
go run . -delete 1
```

---

# Toggle Todo Status

Toggle between completed and incomplete.

```bash
go run . -toggle 0
```

### Example

```bash
[x] Learn Golang
```

becomes

```bash
[ ] Learn Golang
```

---

# Edit a Todo

Edit the title of a todo.

## Format

```bash
-edit "index:new title"
```

## Example

```bash
go run . -edit "0:Learn Advanced Golang"
```

---

# Complete Example Workflow

## Add Todos

```bash
go run . -add "Learn Go"
go run . -add "Build Todo App"
```

---

## List Todos

```bash
go run . -list
```

Output:

```bash
0. [ ] Learn Go
1. [ ] Build Todo App
```

---

## Complete First Todo

```bash
go run . -complete 0
```

---

## Toggle Second Todo

```bash
go run . -toggle 1
```

---

## Edit Todo

```bash
go run . -edit "1:Build Advanced Todo App"
```

---

## Delete Todo

```bash
go run . -delete 0
```

---

# Project Structure

```bash
.
├── main.go
├── todos.go
├── cmdflags.go
└── README.md
```

---

# Code Concepts Used

This project helps practice:

- Structs
- Methods
- Pointer Receivers
- Slices
- Go Flag Package
- Command-Line Applications
- Error Handling
- CRUD Operations

---

# Error Handling

The application validates indexes before operations.

Example invalid command:

```bash
go run . -complete 10
```

Output:

```bash
Invalid index for complete command
```

---

# Build Executable

You can create a standalone executable.

## Linux / Mac

```bash
go build -o todo
```

Run:

```bash
./todo -list
```

---

## Windows

```bash
go build -o todo.exe
```

Run:

```bash
todo.exe -list
```

---

# Future Improvements

- Save todos in JSON file
- Add due dates
- Add priorities
- Search todos
- Colored terminal output
- Interactive terminal UI (TUI)
- Database support

---

# Example Commands Summary

```bash
go run . -add "Study Go"
go run . -list
go run . -complete 0
go run . -toggle 0
go run . -edit "0:Master Golang"
go run . -delete 0
```

---

# Author

Mahidul Anik

---

# License

This project is open-source and available under the MIT License.