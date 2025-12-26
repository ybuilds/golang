package models

import (
	"fmt"
	"time"
)

type User struct {
	Name string;
	Email string;
	Age int64;
	Created time.Time;
}

func DisplayDetail(user *User) {
	fmt.Printf("Name: %s(%d) [%s] - ", user.Name, user.Age, user.Email);
	fmt.Printf("Creation time: %v\n", user.Created);
}

/* For sending function that sends a value as agrument
func DisplayDetail(user User) {
	fmt.Printf("Name: %s(%d) [%s] - ", user.Name, user.Age, user.Email);
	fmt.Printf("Creation time: %v\n", user.Created);
} */