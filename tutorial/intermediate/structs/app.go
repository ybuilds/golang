package main

import (
	"fmt"

	"ybuilds.in/structs/utils"
)

func main() {
	var name, email string;

	name = utils.GetUserData("Enter name");
	email = utils.GetUserData("Enter email");

	fmt.Printf("Welcome %s![%s]\n", name, email);
}
