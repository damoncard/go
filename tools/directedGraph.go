package tools

import "math"

func TransitiveClosure(matrix [][]int) [][]int {
	for lock := 0 ; lock < len(matrix) ; lock++ {
		for row := 0 ; row < len(matrix) ; row++ {
			for col := 0 ; col < len(matrix) ; col++ {
				if matrix[row][lock] == 1 && matrix[lock][col] == 1 {
					matrix[row][col] = 1
				}
			}
		}
	}
	return matrix
}

func DistanceMatrix(matrix [][]int) [][]int {
	for lock := 0 ; lock < len(matrix) ; lock++ {
		for row := 0 ; row < len(matrix) ; row++ {
			for col := 0 ; col < len(matrix) ; col++ {
				if matrix[lock][col] > 0 && matrix[row][lock] > 0 {
					if matrix[row][col] == -1 {
						matrix[row][col] = 999
					}
					matrix[row][col] = int(math.Min(float64(matrix[row][col]), float64(matrix[lock][col] + matrix[row][lock])))
				}
			}
		}
	}
	return matrix
}