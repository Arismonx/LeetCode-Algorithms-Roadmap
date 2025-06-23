package main

import "fmt"

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	for i := range output {
		output[i] += 1
	}

	left, right := 1, 1

	for i := range len(nums) {
		output[i] *= left
		left *= nums[i]
	}
	for i := len(nums) - 1; -1 < i; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}

func main() {
	nums := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
	fmt.Println(productExceptSelf(nums2))
}
