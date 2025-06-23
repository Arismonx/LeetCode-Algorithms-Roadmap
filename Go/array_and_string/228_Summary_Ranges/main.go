package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	output := []string{}
	if len(nums) == 0 {
		return output
	}
	start := nums[0]

	for i := range nums {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if start == nums[i] {
				output = append(output, fmt.Sprintf("%v", start))
			} else {
				output = append(output, fmt.Sprintf("%v->%v", start, nums[i]))
			}
			if i < len(nums)-1 {
				start = nums[i+1]
			}
		}
	}

	return output
}
func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	nums2 := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))  // Output: ["0->2", "4->5", "7"]
	fmt.Println(summaryRanges(nums2)) // Output: ["0", "2->4", "6", "8->9"]
}
