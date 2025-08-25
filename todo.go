package main

import "time"

type Todo struct {
	ID        int
	Title     string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Completed bool
}

type Todos []Todo