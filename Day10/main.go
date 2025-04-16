package main

import "fmt"

func main() {
	arr := [5]int{3, 6, 1}
	//_ = arr
	//sum := SumOfArray() //without passing any values
	//sum := SumOfArray(4, 5) //with two values
	//sum := SumOfArray(1, 2, 3, 4, 5, 6, 7, 8, 9) //with multiple values
	sum := SumOfArray(arr[:]...) //passing array
	//sum := SumOfArray(arr...) //passing slice

	fmt.Printf("Sum of given array %v\n", sum)

	//SquareOfaNum(4, 5, Greeting)
	result := calculation(2, 5, multiplyNumbers)
	fmt.Println("Result of calculation function with values 2, 5 with function multiply numbers:", result)

	SquareOfaNum(2, 3, Greeting)

	SquareOfaNum(200, 300, displayNumbers)
}

func SumOfArray(arr ...int) int { //varied function
	var sum int
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

// // passing a function as an argument
func SquareOfaNum(num1 int, num2 int, op func(int, int)) {

	fmt.Printf("function square of number - num1: %v, num2: %v square: %v\n", num1, num2, num1*num2)
	op(num1, num2)
}

func Greeting(a int, b int) {
	fmt.Printf("Greeting Print a: %v, b: %v\n", a, b)
}

// Passing a Function as an Argument
func calculation(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func multiplyNumbers(x int, y int) int {
	return x * y
}

func displayNumbers(number1 int, number2 int) {
	fmt.Printf("displayNumbers function num 1: %v, num 2: %v\n", number1, number2)
}
