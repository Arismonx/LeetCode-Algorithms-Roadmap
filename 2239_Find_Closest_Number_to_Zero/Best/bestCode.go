package main

import "fmt"

func findClosestNumber(nums []int) int {
	closest := nums[0]

	for _, num := range nums {
		if abs(num) < abs(closest) {
			closest = num
		} else if abs(num) == abs(closest) && num > closest {
			// ถ้าเท่ากัน เลือกค่าบวก
			closest = num
		}
	}

	return closest
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	num1 := []int{-4, -2, 1, 4, 8}
	num2 := []int{2, -1, 1}
	num3 := []int{-10000, -10000}
	f := findClosestNumber(num1)
	f2 := findClosestNumber(num2)
	f3 := findClosestNumber(num3)
	fmt.Println("Result1: ", f)
	fmt.Println("Result2: ", f2)
	fmt.Println("Result3: ", f3)
}
