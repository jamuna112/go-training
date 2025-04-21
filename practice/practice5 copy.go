package main

import "fmt"

/*
ask user for values of two matrix, and then ask for operation - addition, subtraction logic
ask user for two dimensions, num of rows,and column
eg: A = {{2, 3}, {4, 5}}, B = {{6, 7}, {8, 9}}

two funcs: one for user input
- matrix value input
- two diff function for addition and subtraction
- another function for displaying the matrix in a proper format
*/

//do this using struct

type Matrix struct {
	row, col int
}

func main() {
	var row, col int
	fmt.Println("User input for matrix 1 and matrix 2:")
	fmt.Print("Enter number of rows and column for matrix: ")
	fmt.Scanf("%v", &row)
	fmt.Scanf("%v", &col)
	fmt.Print("Enter value for matrix 1: ")
	matrix1 := inputMatrix(row, col)
	fmt.Println("Displaying matrix 1:")
	displayMatrix(row, col, matrix1)
	fmt.Print("Enter value for matrix 2: ")
	matrix2 := inputMatrix(row, col)
	fmt.Println("Displaying matrix 2:")
	displayMatrix(row, col, matrix2)

	addMatrix := matrixAddition(row, col, matrix1, matrix2)
	fmt.Println("Displaying addition of matrix:")
	displayMatrix(row, col, addMatrix)

}

// enter value for matrix 1
func inputMatrix(row, col int) [][]int {
	mat := make([][]int, col)
	var arrValue int
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter value for [%v][%v] ", i, j)
			fmt.Scanf("%d", &arrValue)
			mat[i] = append(mat[i], arrValue)
		}
	}
	return mat
}

func displayMatrix(row, col int, mat [][]int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%v ", mat[i][j])
		}
		fmt.Println()
	}
}

func matrixAddition(row, col int, matrix1, matrix2 [][]int) [][]int {
	result := make([][]int, row*col)
	var add int

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			add = matrix1[i][j] + matrix2[i][j]
			result[i] = append(result[i], add)
		}

	}
	return result
}
