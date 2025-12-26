package main

import "fmt"

func countDigits(num int) int {
	var count int = 0;

	for num > 0 {
		num = num / 10;
		count++;
	}

	return count;
}

func main() {
	var num int;

	fmt.Print("Enter a number: ");
	fmt.Scan(&num);

	fmt.Printf("Number of digits in %d is %d\n", num, countDigits(num));
}