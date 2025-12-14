package main

import "fmt"

func profitCalculator() {
	fmt.Println("Profit Calculator");
	fmt.Println("-----------------");

	var revenue, expense, taxRate, earning, profit, profitRatio float64;

	fmt.Print("Enter revenue: ");
	fmt.Scan(&revenue);

	fmt.Print("Enter expenses: ");
	fmt.Scan(&expense);

	fmt.Print("Enter taxRate: ");
	fmt.Scan(&taxRate);

	earning = revenue - expense;
	profit = earning - (earning * taxRate / 100);
	profitRatio = profit / earning;
	
	fmt.Println("\nSummary");
	fmt.Println("-------");
	fmt.Println("Earning before tax =", earning);
	fmt.Println("Earning after tax =", profit);
	fmt.Println("Profit ratio =", profitRatio);
}