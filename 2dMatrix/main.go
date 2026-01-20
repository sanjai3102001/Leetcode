package main

import "fmt"

func main() {
	matrix := [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
	count := 0
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] < 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
