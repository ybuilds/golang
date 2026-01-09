package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0

	for i < len(nums) {
		if nums[i] == 0 {
			i++
			continue
		}

		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j++
	}
}

func moveZeroesDriver() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
}
