package main

import "fmt"

func gcd(a int, b int) int {
	if b == 0 {
		return a;
	}
	return gcd(b, a%b);
}

func eulerPhi(n int) {
	var count int = 0;
	for i:=1; i<n; i++ {
		if gcd(i, n) == 1 {
			fmt.Print(i, " ");
			count++;
		}
	}
	fmt.Println("\nNumber of co primes upto", n, "=", count);
}