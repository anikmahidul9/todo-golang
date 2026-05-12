package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*t = append(*t, todo)
}

func (t *Todos) complete(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Index out of range")
		fmt.Println(err)
		return err
	}
	(*t)[index].Completed = true
	now := time.Now()
	(*t)[index].CompletedAt = &now
	return nil
}

func (t *Todos) delete(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Index out of range")
		fmt.Println(err)
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *Todos) toggle(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Index out of range")
		fmt.Println(err)
		return err
	}
	(*t)[index].Completed = !(*t)[index].Completed
	if (*t)[index].Completed {
		now := time.Now()
		(*t)[index].CompletedAt = &now
	} else {
		(*t)[index].CompletedAt = nil
	}
	return nil
}

func (t *Todos) edit(index int, newTitle string) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("Index out of range")
		fmt.Println(err)
		return err
	}
	(*t)[index].Title = newTitle
	return nil
}

func (t *Todos) print() {
	for i, todo := range *t {
		status := "Pending"
		if todo.Completed {
			status = "Completed"
		}

		completedAt := "N/A"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
		}

		fmt.Printf("%d: %s [%s] %s %s\n", i, todo.Title, status, todo.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)
	}
}