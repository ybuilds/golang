package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetData(prompt string) (string, error) {
	fmt.Print(prompt + ": ")

	scanner := bufio.NewReader(os.Stdin)
	value, err := scanner.ReadString('\n')

	if err != nil {
		return "", errors.New("Error fetching value from standard input")
	}

	value = strings.Trim(value, "\n")
	value = strings.Trim(value, "\r")

	return value, nil
}
