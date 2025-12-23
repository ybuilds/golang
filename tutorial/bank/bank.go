package main

import (
	"fmt"
	"os"
)

func bankDriver() {
	var balance int = 0;
	var amount int;
	var choice int;

	fmt.Print("Welcome to Go Bank!\n");
	fmt.Print("-------------------\n");
	fmt.Print("Choices: 1.Balance 2.Withdraw 3.Deposit 4.Exit: ");
	fmt.Scan(&choice);

	if choice == 1 {
		fmt.Printf("Current balance: %d\n", balance);
	} else if choice == 2 {
		fmt.Printf("Enter amount to be withdrawn: ");
		fmt.Scan(&amount);

		balance -= amount;
		fmt.Printf("Current balance: %d", balance);
	} else if choice == 3 {
		fmt.Printf("Enter amount to be deposited: ");
		fmt.Scan(&amount);
		
		balance += amount;
		fmt.Printf("Current balance: %d", balance);
	} else if choice == 4 {
		fmt.Printf("Thanks for using Go Bank!");
		os.Exit(0);
	} else {
		fmt.Printf("Invalid choice!\n");
	}
}