package main

import (
	"fmt"
)

type Todoss struct {
	Items []string
}

func main() {
	todos := &Todos{}
	storage := NewStorage[Todos]("todos.json")

	// Load existing todos from file
	loadedTodos, err := storage.Load()
	if err == nil {
		*todos = loadedTodos
	}

	cmd := &cmdFlags{}
	cmd.parse(nil)
	cmd.Execute(todos)

	// Save updated todos to file
	err = storage.Save(*todos)
	if err != nil {
		fmt.Println("Error saving todos:", err)
	}
}
