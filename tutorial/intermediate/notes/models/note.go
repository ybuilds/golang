package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"desc"`
	Created     time.Time `json:"created_at"`
}

func NewNote(title, description string) (*Note, error) {
	if title == "" || description == "" {
		fmt.Printf("E - need title and description for note in NewNote")
		return nil, errors.New("")
	}

	return &Note{
		Title:       title,
		Description: description,
		Created:     time.Now(),
	}, nil
}

func (note *Note) ViewNote() {
	fmt.Printf("Title: %s\n", note.Title)
	fmt.Printf("Description: %s\n", note.Description)
}

func (note *Note) SaveAsJson() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename)
	jsonData, jsonError := json.Marshal(note)

	if jsonError != nil {
		fmt.Printf("E - error converting note to json in SaveAsJson")
		return jsonError
	}

	os.WriteFile(
		"./json/"+filename+".json",
		jsonData,
		0644,
	)

	return nil
}
