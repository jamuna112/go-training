package main

import "fmt"

/*
ask user for values of two matrix, and then ask for operation - addition, subtraction logic
ask user for two dimensions, num of rows,and column
eg: A = {{2, 3}, {4, 5}}, B = {{6, 7}, {8, 9}}

A = | 1  2 |     B = | 5  6 |     A + B = | 6     8 |

	| 3  4 |         | 7  8 |			  | 10   12 |

two funcs: one for user input
- matrix value input
- two diff function for addition and subtraction
- another function for displaying the matrix in a proper format
*/


func main() {

	var rows, cols int

	fmt.Print("how many rows and column: ")
	fmt.Scanf("%v", &rows)
	fmt.Scanf("%v", &cols)

	fmt.Print("Enter values for matrix 1:\n")
	matrix1 := matrix(rows, cols)
	fmt.Print("Displaying matrix 1:\n")
	displayMatrix(rows, cols, matrix1)

	fmt.Print("Enter values for matrix 2:\n")
	matrix2 := matrix(rows, cols)
	fmt.Print("Displaying matrix 2:\n")
	displayMatrix(rows, cols, matrix2)

	fmt.Print("Addition of matrix 1 and matrix 2:\n")
	result := additionMatrix(rows, cols, matrix1, matrix2)
	displayMatrix(rows, cols, result)

}

func matrix(row, col int) [][]int {
	var forRow int
	mat := make([][]int, col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter value for [%v][%v] ", i, j)
			fmt.Scanf("%v", &forRow)
			mat[i] = append(mat[i], forRow)
		}
	}
	return mat
}

func displayMatrix(row, col int, matrix [][]int) [][]int {

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%4v ", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

func additionMatrix(row, col int, matrix1, matrix2 [][]int) [][]int {

	add := make([][]int, row*col)
	var output int

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			output = matrix1[i][j] + matrix2[i][j]
			add[i] = append(add[i], output)
		}
	}
	return add
}
