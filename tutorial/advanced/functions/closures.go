package main

import "fmt"

type multiplier func(int) int

// Factory function
func transformNumber(factor int) multiplier {
	//Closure function: Every anonymous function is a closure
	//Function that multiplies the provided number 'n' with 'factor'
	return func(n int) int {
		//factory is locked in the scope of closure
		return n * factor
	}
}

func arrayTransformer(nums *[]int, transformer multiplier) *[]int {
	res := []int{}
	for _, i := range *nums {
		res = append(res, transformer(i))
	}

	return &res
}

func fibbonacci() func() int {
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}

func driver() {
	// nums := []int{1, 2, 3, 4, 5}
	// doubled := arrayTransformer(&nums, transformNumber(10))
	// fmt.Println(doubled)

	fibbo := fibbonacci()

	for i := range 10 {
		fmt.Printf("Term-%d: %d\n", i, fibbo())
	}
}
