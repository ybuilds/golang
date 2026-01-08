package utils

import (
	"encoding/json"
	"errors"
	"os"

	"ybuilds.in/todo-cli/models"
)

func SaveTodo(todo *models.Todo) error {
	jsonData, error := json.Marshal(todo)

	if error != nil {
		return errors.New("Error saving todo as JSON")
	}

	os.WriteFile(
		"data/todos.json",
		jsonData,
		0644,
	)

	return nil
}
