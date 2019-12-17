package matrix_rotation_test

import "github.com/sepetrov/hackerrank/matrix_rotation"

func ExampleMatrixRotation1() {
	m := [][]int32{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	r := int32(2)
	matrix_rotation.MatrixRotation(m, r)

	// Output:
	//
	// 3 4 8 12
	// 2 11 10 16
	// 1 7 6 15
	// 5 9 13 14
}

func ExampleMatrixRotation2() {
	m := [][]int32{
		{1, 2, 3, 4},
		{7, 8, 9, 10},
		{13, 14, 15, 16},
		{19, 20, 21, 22},
		{25, 26, 27, 28},
	}
	r := int32(7)
	matrix_rotation.MatrixRotation(m, r)

	// Output:
	//
	// 28 27 26 25
	// 22 9 15 19
	// 16 8 21 13
	// 10 14 20 7
	// 4 3 2 1
}

func ExampleMatrixRotation3() {
	m := [][]int32{
		{1, 1},
		{1, 1},
	}
	r := int32(3)
	matrix_rotation.MatrixRotation(m, r)

	// Output:
	//
	// 1 1
	// 1 1
}
