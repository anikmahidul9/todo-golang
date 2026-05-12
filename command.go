package main

import (
	"flag"
	"fmt"
)

type cmdFlags struct {
	Add      string
	Complete int
	Delete   int
	List     bool
	Toggle   int
	Edit     string
}

func (f *cmdFlags) parse(args []string) {
	cf := cmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Complete, "complete", -1, "Mark a todo as completed by index")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion status of a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo's title by index (format: index:new title)")

	flag.Parse()
	*f = cf
}

func (f *cmdFlags) Execute(todos *Todos) {
	if f.Add != "" {
		todos.add(f.Add)
	}
	if flag.Lookup("complete").Value.String() != "-1" {
		if f.Complete < 0 || f.Complete >= len(*todos) {
			fmt.Println("Invalid index for complete command")
		} else {
			todos.complete(f.Complete)
		}
	}
	if flag.Lookup("delete").Value.String() != "-1" {
		if f.Delete < 0 || f.Delete >= len(*todos) {
			fmt.Println("Invalid index for delete command")
		} else {
			todos.delete(f.Delete)
		}
	}
	if f.List {
		todos.print()
	}
	if flag.Lookup("toggle").Value.String() != "-1" {
		if f.Toggle < 0 || f.Toggle >= len(*todos) {
			fmt.Println("Invalid index for toggle command")
		} else {
			todos.toggle(f.Toggle)
		}
	}
	if f.Edit != "" {
		var index int
		var newTitle string
		n, err := fmt.Sscanf(f.Edit, "%d:%s", &index, &newTitle)
		if err != nil || n != 2 {
			fmt.Println("Invalid format for edit. Use: index:new title")
			return
		}
		if index < 0 || index >= len(*todos) {
			fmt.Println("Invalid index for edit command")
		} else {
			todos.edit(index, newTitle)
		}
	}
}
