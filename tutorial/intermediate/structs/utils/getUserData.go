package utils

import "fmt"

func GetUserData(prompt string) string {
	fmt.Printf("%s: ", prompt);
	
	var data string;
	fmt.Scan(&data);

	return data;
}