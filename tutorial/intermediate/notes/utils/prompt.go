package utils

import (
	"errors"
	"fmt"
)

func GetUserData(prompt string) (string, error) {
	var value string

	fmt.Printf("%s: ", prompt)
	fmt.Scan(&value)

	if value == "" {
		fmt.Printf("E - empty user data for prompt '%s'", prompt)
		return "", errors.New("empty user data")
	}

	return value, nil
}
