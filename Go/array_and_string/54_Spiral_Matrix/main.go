package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	output := []int{}
	rows, cols := len(matrix), len(matrix[0])
	x, y, dx, dy := 0, 0, 1, 0
	for i := 0; i < rows*cols; i++ {
		output = append(output, matrix[y][x])
		matrix[y][x] = -101
		if !(0 <= x+dx && x+dx < cols && 0 <= y+dy && y+dy < rows) || matrix[y+dy][x+dx] == -101 {
			temp := dx
			dx = -dy
			dy = temp
		}
		x += dx
		y += dy
	}
	return output
}

func main() {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	merged := spiralOrder(input)

	fmt.Println(merged)
}
