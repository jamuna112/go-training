package main

import "fmt"

/*
NOTE: By using struct
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

type Matrix struct {
	rows, cols int
}

func main() {

	var rows, cols int

	fmt.Print("how many rows and column: ")
	fmt.Scanf("%v", &rows)
	fmt.Scanf("%v", &cols)
	matrix := Matrix{rows: rows, cols: cols}

	fmt.Print("Enter values for matrix 1:\n")
	matrix1 := userMatrix(matrix)
	fmt.Print("Displaying matrix 1:\n")
	displayMatrix(matrix, matrix1)

	fmt.Print("Enter values for matrix 2:\n")
	matrix2 := userMatrix(matrix)
	fmt.Print("Displaying matrix 2:\n")
	displayMatrix(matrix, matrix2)

	performOperation(matrix, matrix1, matrix2)

}

func userMatrix(matx Matrix) [][]int {
	var forRow int
	mat := make([][]int, matx.cols)

	for i := 0; i < matx.rows; i++ {
		for j := 0; j < matx.cols; j++ {
			fmt.Printf("Enter value for [%v][%v] ", i, j)
			fmt.Scanf("%v", &forRow)
			mat[i] = append(mat[i], forRow)
		}
	}
	return mat
}

func displayMatrix(matx Matrix, matrix [][]int) [][]int {

	for i := 0; i < matx.rows; i++ {
		for j := 0; j < matx.cols; j++ {
			fmt.Printf("%4v ", matrix[i][j])
		}
		fmt.Println()
	}
	return matrix
}

func additionMatrix(matx Matrix, matrix1, matrix2 [][]int) [][]int {

	add := make([][]int, matx.rows*matx.cols)
	var output int

	for i := 0; i < matx.rows; i++ {
		for j := 0; j < matx.cols; j++ {
			output = matrix1[i][j] + matrix2[i][j]
			add[i] = append(add[i], output)
		}
	}
	return add
}

func subtractionMatrix(matx Matrix, matrix1, matrix2 [][]int) [][]int {

	sub := make([][]int, matx.rows*matx.cols)
	var output int

	for i := 0; i < matx.rows; i++ {
		for j := 0; j < matx.cols; j++ {
			output = matrix1[i][j] - matrix2[i][j]
			sub[i] = append(sub[i], output)
		}
	}
	return sub
}

func performOperation(matx Matrix, matrix1, matrix2 [][]int) {
	var userInput string
	for {
		fmt.Print("what operation you want to perform? add or sub: ")
		fmt.Scanf("%v", &userInput)

		if userInput == "add" {
			addMatrix := additionMatrix(matx, matrix1, matrix2)
			fmt.Println("Displaying addition of matrix:")
			displayMatrix(matx, addMatrix)
			break
		} else if userInput == "sub" {
			subMatrix := subtractionMatrix(matx, matrix1, matrix2)
			fmt.Println("Displaying subtraction of matrix:")
			displayMatrix(matx, subMatrix)
			break
		} else {
			fmt.Println("Invalid input!")
		}
	}

}
