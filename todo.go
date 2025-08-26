package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	ID          int
	Title       string
	CreatedAt   *time.Time
	CompletedAt *time.Time
	Completed   bool
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	now := time.Now()
	todo := Todo{
		ID:        len(*todos) + 1,
		Title:     title,
		CreatedAt: &now,
		// CompletedAt: &now,
		Completed: false,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) List() Todos {
	return *todos
}
func (todos *Todos) find(index int) (*Todo, error) {
	for i, todo := range *todos {
		if i == index {
			return &todo, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range") /* invalid index */
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	// Delete the todo at the specified index before and after the validation
	*todos = append((*todos)[:index], (*todos)[index+1:]...)
	return nil
}
func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	todo := &(*todos)[index]
	if !(todo.Completed) {
		now := time.Now()
		todo.CompletedAt = &now
	}
	todo.Completed = !todo.Completed

	return nil
}
func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}
	todo := &(*todos)[index]
	todo.Title = title
	return nil
}

/* func (todos *Todos) print (){
	table := ""
	for _, todo := range *todos {
		table += fmt.Sprintf("| %d | %s | %s | %s |\n",
			todo.ID,
			todo.Title,
			todo.CreatedAt.Format(time.RFC3339),
			todo.CompletedAt.Format(time.RFC3339))
	}
	fmt.Println(table)
} */
func (todos *Todos) print() {
	t := table.New(os.Stdout)
	t.SetHeaders("ID", "Title", "Created At", "Completed At", "Completed")

	for index, todo := range *todos {
		var completedAtStr string
		if todo.CompletedAt != nil {
			completedAtStr = todo.CompletedAt.Format(time.RFC3339)
		} else {
			completedAtStr = ""
		}
		var completedStr string
		if todo.Completed {
			completedStr = "✅"
		} else {
			completedStr = "❌"
		}
		t.AddRow(
			// strconv.Itoa(todo.ID),
			strconv.Itoa(index),
			todo.Title,
			todo.CreatedAt.Format(time.RFC3339),
			completedAtStr,
			// fmt.Sprintf("%v", todo.Completed),
			completedStr,
		)
	}
	t.Render()
}
