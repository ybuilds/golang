package models

import (
	"fmt"
	"time"
)

type Todo struct {
	Title   string
	Content string
	Created time.Time
}

func NewTodo(title, content string) *Todo {
	return &Todo{
		Title:   title,
		Content: content,
		Created: time.Now(),
	}
}

func (todo *Todo) DisplayTodo() {
	fmt.Printf("%s: %s (%v)\n", todo.Title, todo.Content, todo.Created)
}
