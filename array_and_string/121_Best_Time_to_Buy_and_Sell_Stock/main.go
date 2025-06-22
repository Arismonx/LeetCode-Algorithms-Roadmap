package main

import "fmt"

func maxProfit(price []int) int {
	min, max := price[0], 0

	for i := 1; i < len(price); i++ {
		num := price[i]
		if num < min {
			min = num
		} else if num-min > max {
			max = num - min
		}
	}
	return max
}

func main() {
	price := []int{7, 1, 5, 3, 6, 4}
	price2 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(price))
	fmt.Println(maxProfit(price2))
}
