package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"ybuilds.in/structs/models"
	"ybuilds.in/structs/utils"
)

func main() {
	var name, email string;

	name = utils.GetUserData("Enter name");
	age, _ := strconv.ParseInt(utils.GetUserData("Enter age"), 10, 32);
	email = utils.GetUserData("Enter email");

	//Creation using CONSTRUCTOR type FUNCTION
	user, error := models.NewUser(name, email, age, time.Now());

	if error != nil {
		fmt.Printf("E[%s]\n", error);
		os.Exit(1);
	}

	user.DisplayDetail();
	user.ResetEmail();
}
