package utils

import "fmt"

func GetNoteData() (string, string, error) {
	title, titleError := GetUserData("Enter note title")

	if titleError != nil {
		fmt.Printf("E - error fetching title in GetNoteData")
		return "", "", titleError
	}

	desc, descError := GetUserData("Enter note description")

	if descError != nil {
		fmt.Printf("E - error fetching description in GetNoteData")
		return "", "", descError
	}

	return title, desc, nil
}
