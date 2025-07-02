package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	num := make([]int, 0, len(nums))

	for _, v := range nums {
		curr := abs(v)
		num = append(num, curr)
	}
	sort.Ints(num)
	return num

}
func abs(x int) int {
	if x < 0 {
		return -x * -x
	}
	return x * x
}

func main() {
	nums := []int{-4, -1, 0, 3, 10} //16 1 0 9 100
	fmt.Println(sortedSquares(nums))
}
