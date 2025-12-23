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
	
	for i:=0; i==0; {
		fmt.Print("Choices: 1.Balance 2.Withdraw 3.Deposit 4.Exit: ");
		fmt.Scan(&choice);

		switch choice {
			case 1: 
				fmt.Printf("Current balance: %d\n", balance);
			case 2: {
				fmt.Printf("Enter amount to be withdrawn: ");
				fmt.Scan(&amount);

				if amount > 0 {
					balance -= amount;
					fmt.Printf("Current balance: %d\n", balance);
				} else {
					fmt.Printf("Enter a valid amount\n");
				}
			}
			case 3: {
				fmt.Printf("Enter amount to be deposited: ");
				fmt.Scan(&amount);
				
				if amount > 0 {
					balance += amount;
					fmt.Printf("Current balance: %d\n", balance);
				} else {
					fmt.Printf("Enter a valid amount\n");
				}
			}
			case 4: {
				fmt.Printf("Thanks for using Go Bank!");
				os.Exit(0);
			}
			default: fmt.Printf("Invalid choice!\n");
		}
	}
}