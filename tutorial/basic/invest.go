package main

import (
	"fmt"
	"math"
)

func investment() {
	var investmentAmount, expectedReturnRate, years float64;

	fmt.Print("Enter investment amount: ");
	fmt.Scan(&investmentAmount);

	fmt.Print("Enter expected return rate: ");
	fmt.Scan(&expectedReturnRate);

	fmt.Print("Enter time period in years: ");
	fmt.Scan(&years);

	const inflationRate float64 = 2.5;
	
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, float64(years));
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100.0, years);
	
	fmt.Println("Expected return value: ", math.Round(futureValue));
	fmt.Println("Expected real return value: ", math.Round(futureRealValue));
}
