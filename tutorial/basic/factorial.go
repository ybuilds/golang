package main

func recursiveFactorial(num int) int {
	if num <= 1 {
		return 1;
	}
	return num * recursiveFactorial(num-1);
}
