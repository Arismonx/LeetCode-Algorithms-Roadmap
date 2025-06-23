package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		left := 0
		right := n - 1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
	fmt.Println(matrix) //[[7,4,1],[8,5,2],[9,6,3]]
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}} //[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
	rotate(matrix)
	rotate(matrix2) //[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
}
