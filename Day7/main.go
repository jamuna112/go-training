package main

import "fmt"

func main() {
	// var number1 int
	// fmt.Print("Enter a number for multiplication table?")
	// fmt.Scanf("%d", &number1)

	/*
		  This is multiplication table of 4
		  _________________________________

			4 * 1 = 4
			4 * 2 = 8
			..
			..
			4 * 10 = 40
	*/

	// fmt.Printf("This is multiplication table of %v\n", number1)
	// fmt.Println("____________________________________\n")
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%v * %2v = %3v\n", number1, i, number1*i)
	// }

	/*
		take two numbers from user, and operations to be performed give the output based on operation.

		Enter two numbers for operation:
		Enter first number? 7
		Enter second number? 8
		Enter operation(+, -, and /)? /

		output of division operation is
		7 / 8 = 0.875
	*/
	var num1 int
	var num2 int
	var sign string
	var operation string
	var output float32

	fmt.Printf("Enter two numbers for operation:\n")
	fmt.Printf("Enter first number?: ")
	fmt.Scanf("%v\n", &num1)
	fmt.Printf("Enter second number?: ")
	fmt.Scanf("%v\n", &num2)
	fmt.Printf("Enter operation(+, -, and /)?")
	fmt.Scanf("%v", &sign)
	if sign == "+" {
		operation = "addition"
		output = float32(num1 + num2)
	} else if sign == "-" {
		operation = "subtraction"
		output = float32(num1 - num2)
	} else if sign == "/" {
		operation = "division"
		output = float32(float32(num1) / float32(num2))
	}
	fmt.Printf("output of %v operation is\n", operation)
	fmt.Printf("%v %v %v = %.2f\n", num1, sign, num2, output)

}
