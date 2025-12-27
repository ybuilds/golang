package models

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	Name string;
	Email string;
	Age int64;
	Created time.Time;
}

//Constructor utility function with validation
func NewUser(name string, email string, age int64, created time.Time) (*User, error) {
	if name == "" {
		return nil, errors.New("Name is empty");
	}

	if age < 18 {
		return nil, errors.New("User is not an adult");
	}

	return &User{
		Name: name,
		Age: age,
		Email: email,
		Created: created,
	}, nil;
}

//Attaching methods
func (user *User) /*receiver argument*/ DisplayDetail() {
	fmt.Printf("Name: %s(%d) [%s] - ", user.Name, user.Age, user.Email);
	fmt.Printf("Creation time: %v\n", user.Created);
}

//Mutation methods
func (user *User) ResetEmail() {
	var email string;

	fmt.Printf("Erasing email...\n");

	fmt.Printf("Enter an email address: ");
	fmt.Scan(&email);

	user.Email = email;

	user.DisplayDetail();
}
