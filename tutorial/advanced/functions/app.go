package main

import "fmt"

type intCompute func(int) int

func multiplyNumber(nums *[]int, multiplier intCompute) []int {
	dNums := []int{}

	for _, i := range *nums {
		dNums = append(dNums, multiplier(i))
	}

	return dNums
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func quadruple(num int) int {
	return num * 4
}

func fetchMultiplier() intCompute {
	var choice int

	fmt.Printf("1.x2 2.x3 3.x4: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		{
			return double
		}
	case 2:
		{
			return triple
		}
	case 3:
		{
			return quadruple
		}
	default:
		{
			return func(n int) int {
				return n
			}
		}
	}
}

func main() {
	// nums := []int{1, 2, 3}

	// multiplier := fetchMultiplier()

	// modNums := multiplyNumber(&nums, multiplier)

	// doubledNums := multiplyNumber(&nums, func(n int) int {
	// 	return n * 2
	// })

	// tripledNums := multiplyNumber(&nums, func(n int) int {
	// 	return n * 3
	// })

	// fmt.Println(modNums)
	// fmt.Println(doubledNums)
	// fmt.Println(tripledNums)

	driver()
}
