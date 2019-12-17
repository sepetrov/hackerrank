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

	l := int(math.Round(math.Max(float64(rows), float64(cols))) / 2.0)
	layers := make([][]int32, l)

	// Walk the matrix clockwise starting from top left and collect the values
	// for each layer.
	for i := 0; i < l; i++ {
		layer := []int32{}

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

		layers[i] = layer
	}

	// Left shift each layer r times.
	for i := 0; i < len(layers); i++ {
		for j := int32(0); j < r; j++ {
			layers[i] = append(layers[i][1:], layers[i][0])
		}
	}

	// Populate the matrix with the values from the layers.
	for i := 0; i < l; i++ {

		// top left to top right
		for j := i; j < cols-i-1; j++ {
			matrix[i][j] = layers[i][0]
			layers[i] = layers[i][1:]
		}

		// top right to bottom right
		for j := i; j < rows-i-1; j++ {
			matrix[j][cols-i-1] = layers[i][0]
			layers[i] = layers[i][1:]
		}

		// bottom right to bottom left
		for j := cols - i - 1; j > i; j-- {
			matrix[rows-i-1][j] = layers[i][0]
			layers[i] = layers[i][1:]
		}

		// bottom left to top left
		for j := rows - i - 1; j > i; j-- {
			matrix[j][i] = layers[i][0]
			layers[i] = layers[i][1:]
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
