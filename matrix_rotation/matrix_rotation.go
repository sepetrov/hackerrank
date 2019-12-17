// Package matrix_rotation contains the solution for HackerRank challenge
// https://www.hackerrank.com/challenges/matrix-rotation-algo/problem.
package matrix_rotation

import (
	"fmt"
	"math"
)

// MatrixRotation prints matrix after rotating its elements anticlockwise r times.
func MatrixRotation(matrix [][]int32, r int32) {
	rows := len(matrix)
	cols := len(matrix[0])

	l := int(math.Round(math.Min(float64(rows), float64(cols))) / 2.0)

	for i := 0; i < l; i++ {
		layer := []int32{}

		// Walk the matrix clockwise starting from top left and collect the values
		// for each layer.

		// top left to top right
		for j := i; j < cols-i-1; j++ {
			layer = append(layer, matrix[i][j])
		}
		// top right to bottom right
		for j := i; j < rows-i-1; j++ {
			layer = append(layer, matrix[j][cols-i-1])
		}
		// bottom right to bottom left
		for j := cols - i - 1; j > i; j-- {
			layer = append(layer, matrix[rows-i-1][j])
		}
		// bottom left to top left
		for j := rows - i - 1; j > i; j-- {
			layer = append(layer, matrix[j][i])
		}

		// Left shift layer r times.
		for j := int32(0); j < r; j++ {
			layer = append(layer[1:], layer[0])
		}

		// Populate the matrix with the values from the layer.

		c := 0

		// top left to top right
		for j := i; j < cols-i-1; j++ {
			matrix[i][j] = layer[c]
			c++
		}
		// top right to bottom right
		for j := i; j < rows-i-1; j++ {
			matrix[j][cols-i-1] = layer[c]
			c++
		}
		// bottom right to bottom left
		for j := cols - i - 1; j > i; j-- {
			matrix[rows-i-1][j] = layer[c]
			c++
		}
		// bottom left to top left
		for j := rows - i - 1; j > i; j-- {
			matrix[j][i] = layer[c]
			c++
		}
	}

	// Print the matrix.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(matrix[i][j])
			if j < cols-1 {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
