package main

import (
	"fmt"

	"ybuilds.in/pointers/utils"
)

func getAdultYear(age *int) int {
	if *age > 18 {
		return *age - 18;
	}

	return -1;
}

func main() {
	var age int;
	var agePointer *int;

	fmt.Printf("Enter age: ");
	fmt.Scan(&age);
	agePointer = &age;

	var adultAge int = getAdultYear(agePointer);
	fmt.Printf("Adult age is %d\n", adultAge);

	var x, y int;

	fmt.Printf("Enter two numbers to be swapped: ");
	fmt.Scan(&x);
	fmt.Scan(&y);

	fmt.Printf("Values of X and Y before swapping: %d & %d\n", x, y);

	utils.Swap(&x, &y);
	fmt.Printf("Values of X and Y after swapping: %d & %d\n", x, y);
}