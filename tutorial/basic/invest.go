package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5;

func calculateValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64){
	var futureValue float64 = investmentAmount * math.Pow(1 + expectedReturnRate / 100, float64(years));
	var futureRealValue float64 = futureValue / math.Pow(1 + inflationRate / 100.0, years);
	return futureValue, futureRealValue;
}

func investment() {
	var investmentAmount, expectedReturnRate, years float64;

	fmt.Print("Enter investment amount: ");
	fmt.Scan(&investmentAmount);

	fmt.Print("Enter expected return rate: ");
	fmt.Scan(&expectedReturnRate);

	fmt.Print("Enter time period in years: ");
	fmt.Scan(&years);
	
	futureValue, futureRealValue := calculateValues(investmentAmount, expectedReturnRate, years);
	
	fmt.Println("Expected return value: ", math.Round(futureValue));
	fmt.Println("Expected real return value: ", math.Round(futureRealValue));
}
