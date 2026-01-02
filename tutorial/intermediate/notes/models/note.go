package models

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title       string
	description string
	created     time.Time
}

func NewNote(title, description string) (*Note, error) {
	if title == "" || description == "" {
		fmt.Printf("E - need title and description for note in NewNote")
		return nil, errors.New("")
	}

	return &Note{
		title:       title,
		description: description,
		created:     time.Now(),
	}, nil
}

func (note *Note) ViewNote() {
	fmt.Printf("Title: %s\n", note.title)
	fmt.Printf("Description: %s\n", note.description)
}
