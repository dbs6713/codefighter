package main

import "log"

func main() {
	matrix := [][]int{
		{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3},
	}
	log.Printf("total: %d", matrixElementsSum(matrix))

	matrix = [][]int{
		{1, 1, 1, 0},
		{0, 5, 0, 1},
		{2, 1, 3, 10},
	}
	log.Printf("total: %d", matrixElementsSum(matrix))
}

func matrixElementsSum(matrix [][]int) int {
	top := 0
	floors := len(matrix)
	rooms := len(matrix[top])
	total := 0
	for f := top; f < floors; f++ {
		for r := 0; r < rooms; r++ {
			roomVal := matrix[f][r]
			if roomVal != 0 {
				total = total + roomVal
			} else {
				for n := f + 1; n < floors; n++ {
					matrix[n][r] = 0
				}
			}
		}
	}
	return total
}
