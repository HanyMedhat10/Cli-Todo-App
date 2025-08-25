package main

import (
	"errors"
	"fmt"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Completed bool
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	now := time.Now()
	todo := Todo{
		ID:        len(*todos) + 1,
		Title:     title,
		CreatedAt: &now,
		UpdatedAt: &now,	
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
		err := errors.New("Index out of range") /* invalid index */
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