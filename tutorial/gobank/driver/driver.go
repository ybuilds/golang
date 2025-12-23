package driver

import (
	"fmt"
	"os"

	"ybuilds.in/gobank/utils"
)

func Start() {
	var choice int;
	balance, err := utils.ReadBalanceFromFile();
	
	if err != nil {
		fmt.Println(err);
		panic("Error starting Go Bank");
	}

	fmt.Printf("--------------Welcome to Go Bank--------------\n");

	for {
		fmt.Printf("Choices: 1.Balance 2.Deposit 3.Withdraw 4.Exit: ");
		fmt.Scan(&choice);

		switch choice {
			case 1: {
				fmt.Printf("Current balance: %.2f\n", balance);
			}
			case 2: {
				var amount float64;
				fmt.Printf("Enter amount to be deposited: ");
				fmt.Scan(&amount);

				if amount > 0 {
					balance += amount;
					utils.WriteToFile(balance);
					fmt.Printf("Deposited %f. Current balance: %.2f\n", amount, balance);
				} else {
					fmt.Printf("Enter a valid amount\n");
				}
			}
			case 3: {
				var amount float64;
				fmt.Printf("Enter amount to be withdrawn: ");
				fmt.Scan(&amount);

				if amount > 0 && amount <= balance {
					balance -= amount;
					utils.WriteToFile(balance);
					fmt.Printf("Withdrawn %f. Current balance: %.2f\n", amount, balance);
				} else {
					fmt.Printf("Enter a valid amount\n");
				}
			}
			case 4: {
				fmt.Printf("Thanks for using Go Bank!");
				os.Exit(0);
			}
			default: {
				fmt.Printf("Enter a correct choice\n");
			}
		}
	}
}