package main

import (
	"fmt"

	"ybuilds.in/todo-cli/models"
	"ybuilds.in/todo-cli/utils"
)

func main() {
	title, error := utils.GetData("Enter title")

	if error != nil {
		fmt.Printf("Error fetching value from F-GetData")
	}

	content, error := utils.GetData("Enter content")

	if error != nil {
		fmt.Printf("Error fetching value from F-GetData")
	}

	todo := models.NewTodo(title, content)

	error = utils.SaveTodo(todo)

	if error != nil {
		fmt.Println("Error saving todo in F-SaveTodo")
	} else {
		fmt.Println("Todo saved as JSON")
	}
}
