package main

import (
	"fmt"

	"ybuilds.in/notes/models"
	"ybuilds.in/notes/utils"
)

func main() {
	title, description, dataErr := utils.GetNoteData()

	if dataErr != nil {
		fmt.Printf("E - error fetching title and description in Main")
		return
	}

	note, titleErr := models.NewNote(title, description)

	if titleErr != nil {
		fmt.Printf("E - error creating a note in Main")
		return
	}

	note.ViewNote()
	saveError := note.SaveAsJson()

	if saveError != nil {
		fmt.Printf("E - error saving note in Main - %v", saveError)
	} else {
		fmt.Printf("Note saved as JSON\n")
	}
}
