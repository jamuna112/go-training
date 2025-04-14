package main

import "fmt"

func main() {
	var num1, num2 int = 8, 9
	array := []int{-5, -16, -7, -8, -3}
	Add(num1, num2)
	result := Radd(num1, num2)
	fmt.Printf("addition of two number is: %v\n", result)
	Greeting()
	fmt.Println("Hello World..")
	Greeting()
	fmt.Printf("Testing multiplication function: %v %v\n", Multiplication(5, 2), Multiplication(5, 8))
	fmt.Printf("Testing division function: %v %v\n", Division(20, 2), Division(15, 3))

	// output := ListOfOperation("+")
	// fmt.Printf("Output : %v\n", output)

	fmt.Printf("finding max element from given array: %v\n", FindMaxElement(array))

}

func Greeting() {
	fmt.Println("Greeting function")
}

func Add(num1 int, num2 int) {
	fmt.Printf("addition of num1 %d plus num2 %d is %d\n", num1, num2, num1+num2)
}

func Radd(num1 int, num2 int) int {
	return num1 + num2
}

/*
Writing a program, getting input from user and operation. Write function for each operation
and give the result
Plan:
- list of operations +, -, *, /
- take two numbers from user
- take operations
- each operations will have separate functions
- display user input and output performed
- ask user do you want to continue
- if yes, continue otherwise break

5 * 8 = 40
5 * 2 = 10, 5 + 5 = 10
5 * 3 = 15, 5 + 5 + 5 = 15
*/

func ListOfOperation(operation string) int {
	var num1, num2, result int
	fmt.Printf("what operation you want to perform? +,-,*,/")
	fmt.Scanf("%v\n", &operation)

	fmt.Printf("Enter first number: ")
	fmt.Scanf("%v\n", &num1)

	fmt.Printf("Enter second number: ")
	fmt.Scanf("%v\n", &num2)

	switch operation {
	case "+":
		result = Addition(num1, num2)
	case "-":
		result = Subtraction(num1, num2)
	case "*":
		result = Multiplication(num1, num2)
	case "/":
		result = Division(num1, num2)
	}
	return result
}

func Addition(num1 int, num2 int) int {
	return num1 + num2
}

func Subtraction(num1 int, num2 int) int {
	return num1 - num2
}

func Multiplication(num1 int, num2 int) int {
	var result int
	/*
			5 * 8 = 40
		5 * 2 = 10, 5 + 5 = 10
		num1 = 5, num2 = 2 =>
		5 * 3 = 15, 5 + 5 + 5 = 15
	*/
	for i := 1; i <= num2; i++ {
		result += num1
	}
	return result
}

// Do this math in note
func Division(num1 int, num2 int) int {
	var result int
	/*
			5 / 2 = 2 => 5-2 = 3, 3-2 = 1
		15 / 3 = 5, 15 - 3 = 12, 12- 3 = 9, 9 - 3 = 6, 6 - 3 = 3, 3 - 3 = 0
		20 / 2 = 10
		num1 = 5, num2 = 2 =>
		5 * 3 = 15, 5 + 5 + 5 = 15
	*/
	for i := num1; i >= num2; i -= num2 {
		result += 1
	}
	return result
}

/*
Write a function of int arrays and give max of array elements.
array:= [4]int{5, 6, 7, 8}
result should be 8
*/

func FindMaxElement(array []int) int {
	var temp int = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > temp {
			temp = array[i]
		}
	}
	return temp
}

/*
Do a reverse string
eg: abc => cba
hello => olleh
NOTE: check this example - use rune from documentation (FIX IT!)
*/
func ReverseStr(word string) string {
	var reverseWord string //o
	var count int = 0
	for i := len(word) - 1; i >= 0; i-- {
		//reverseWord[count] = word[i]
		count++
	}
	return reverseWord
}

/*
1. complete addition subtraction program
2. Pass an array and find max and min, average
3. pass an array and enter number print if number exits in a array
4. Fix reverse string program
5. Pass a string and check if it's a palindrome
6. Pass to dimension matrix, A and D and return A + B
*/
