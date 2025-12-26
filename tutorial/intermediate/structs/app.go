package main

import (
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

	user := new(models.User);

	user.Name = name;
	user.Age = age;
	user.Email = email;
	user.Created = time.Now();

	models.DisplayDetail(user);

	/* Pass struct type as a value directly

	user := models.User {
		Name: name,
		Age: age,
		Email: email,
		Created: time.Now(),
	}

	models.DisplayDetail(user); */
}
