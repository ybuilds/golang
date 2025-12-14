package main

import (
	"fmt"
	"math"
)

func investment(investmentAmount float64, expectedReturnRate float64, years float64) {
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, float64(years));
	fmt.Println("Expected return value: ", math.Round(futureValue));
}

func customerDetail() {
	var name, city string = "yashas", "sbc";
	fmt.Println("User", name, "from", city);
}
