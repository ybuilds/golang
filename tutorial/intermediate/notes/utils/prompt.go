package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetUserData(prompt string) (string, error) {
	var value string

	fmt.Printf("%s: ", prompt)

	scanner := bufio.NewReader(os.Stdin)
	value, err := scanner.ReadString('\n')

	if err != nil {
		fmt.Printf("E - error reading data for prompt '%s'", prompt)
		return "", errors.New("error reading data")
	}

	value = strings.Trim(value, "\n")
	value = strings.Trim(value, "\r")

	return value, nil
}
