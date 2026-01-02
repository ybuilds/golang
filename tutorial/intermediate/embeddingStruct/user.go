package main

import "fmt"

type User struct {
	firstname string
	lastname  string
	email     string
}

type Admin struct {
	location   string
	department string
	user       User
}

func NewUser(firstname, lastname, email string) *User {
	return &User{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
	}
}

func (user *User) displayUser() {
	fmt.Printf(
		"%s, %s (%s)\n",
		user.lastname,
		user.firstname,
		user.email,
	)
}

func NewAdmin(location, department, firstname, lastname, email string) *Admin {
	return &Admin{
		location:   location,
		department: department,
		user: User{
			firstname: firstname,
			lastname:  lastname,
			email:     email,
		},
	}
}

func (admin *Admin) displayAdmin() {
	fmt.Printf(
		"%s, %s (%s) - %s [BU: %s]\n",
		admin.user.lastname,
		admin.user.firstname,
		admin.user.email,
		admin.location,
		admin.department,
	)
}
