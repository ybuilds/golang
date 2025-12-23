package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName string = "balance.dat";

func writeBalanceToFile(balance int) {
	balanceText := fmt.Sprint(balance);
	os.WriteFile(fileName, []byte(balanceText), 0777);
}

func readBalanceFromFile() (int, error) {
	data, error := os.ReadFile(fileName);
	
	if error != nil {
		fmt.Printf("File not found\n");
		return 1000, errors.New("File not found, returning default balance");
	}

	balanceText := string(data);
	balance, error := strconv.ParseFloat(balanceText, 64);

	if error != nil {
		fmt.Printf("Error parsing data\n");
		return 1000, errors.New("Failed to parse, returning default balance");
	}

	return int(balance), nil;
}

func bankDriver() {
	balance, error := readBalanceFromFile();
	
	if error != nil {
		fmt.Printf("ERR:[%s]\n", error);
		os.Exit(1);
	}
	
	var amount int;
	var choice int;

	fmt.Printf("------GO BANK------\n");
	fmt.Printf("Welcome to Go Bank!\n");
	fmt.Printf("-------------------\n");
	
	for i:=0; i==0; {
		fmt.Print("Choices: 1.Balance 2.Withdraw 3.Deposit 4.Exit: ");
		fmt.Scan(&choice);

		switch choice {
			case 1: 
				balance, _ = readBalanceFromFile();
				fmt.Printf("Current balance: %d\n", balance);
			case 2: {
				fmt.Printf("Enter amount to be withdrawn: ");
				fmt.Scan(&amount);

				if amount > 0 && amount <= balance {
					balance -= amount;
					fmt.Printf("Current balance: %d\n", balance);
					writeBalanceToFile(balance);
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
					writeBalanceToFile(balance);
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