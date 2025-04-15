package main

import "fmt"

func main() {

	//variables declaration
	array := [5]int{3, 9, 2, 5, 1}
	array1 := []int{12, 26, 11, 18}
	math1 := [3][3]int{{10, 20, 30}, {40, 50, 60}, {70, 80, 90}}
	math2 := [3][3]int{{100, 200, 300}, {400, 500, 600}, {700, 800, 900}}

	//1. complete addition subtraction program
	ListOfOperation()

	//2. Pass an array and find max and min, average
	max, min, average := findMaxMinAvg(array1)
	fmt.Printf("max: %v\nmin: %v\naverage: %v\n", max, min, average)

	// 3. pass an array and enter number print if number exits in a array
	numExists := CheckNumberIfExists(11, array[:])
	fmt.Printf("Given number exists: %t\n", numExists)

	//4. reverse string program
	reverseWord := ReverseStr("hello")
	fmt.Printf("Reverse word: %v\n", reverseWord)

	//5. Pass a string and check if it's a palindrome
	palindromeString := isPalindrome("kayak")
	fmt.Printf("%v\n", palindromeString)

	//6. Pass to dimension matrix, A and B and return A + B
	result := AdditionMatrix(math1, math2)
	fmt.Printf("sum of math1 and math2:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%4v", result[i][j])
		}
		fmt.Println()
	}

	//7. Program to print first 3 letters of given string
	word := "Sunday"
	fmt.Printf("first 3 letters of %v are %v", word, First3Letters(word))
	//================================================================================//

}

func Greeting() {
	fmt.Println("Greeting function")
}

func Add(num1 int, num2 int) {
	fmt.Printf("addition of %d + %d = %d\n", num1, num2, num1+num2)
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
- if yes continue, otherwise break

5 * 8 = 40
5 * 2 = 10, 5 + 5 = 10
5 * 3 = 15, 5 + 5 + 5 = 15
*/

func ListOfOperation() {
	var num1, num2, result int
	var operation string
	for {
		fmt.Printf("what operation you want to perform? +,-,*,/")
		fmt.Scanf("%v\n", &operation)

		fmt.Printf("Enter first number: ")
		fmt.Scanf("%v\n", &num1)

		fmt.Printf("Enter second number: ")
		fmt.Scanf("%v\n", &num2)

		switch operation {
		case "+":
			result = Addition(num1, num2)
			fmt.Printf("addition of %d + %d = %d\n", num1, num2, result)
		case "-":
			result = Subtraction(num1, num2)
			fmt.Printf("subtraction of %d - %d = %d\n", num1, num2, result)
		case "*":
			result = Multiplication(num1, num2)
			fmt.Printf("multiplication of %d * %d = %d\n", num1, num2, result)
		case "/":
			result = Division(num1, num2)
			fmt.Printf("division of %d / %d = %d\n", num1, num2, result)
		}
		var response string
		fmt.Printf("Do you want to continue?\n")
		fmt.Scanf("%v\n", &response)
		if response == "yes" {
			continue
		} else {
			break
		}
	}
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
*/
func ReverseStr(word string) string {
	runes := []byte(word)
	length := len(runes)

	for i := 0; i < length; i++ {
		l := word[length-1-i]
		runes[i] = word[l]
		//runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}

// create new one with the reverse function
func isPalindrome(str string) bool {
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func findMaxMinAvg(num []int) (int, int, float32) {
	max := num[0]
	min := num[0]
	sum := 0
	var average float32
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
		if num[i] < min {
			min = num[i]
		}
		sum += num[i]
		average = float32(sum) / (float32)(len(num))
	}
	return max, min, average
}

func CheckNumberIfExists(num int, array []int) bool {
	for i := 0; i < len(array); i++ {
		if num == array[i] {
			return true
		}
	}
	return false
}

func AdditionMatrix(math1 [3][3]int, math2 [3][3]int) [3][3]int {
	var math3 [3][3]int
	//add matrix 1 and matrix 2
	for i := 0; i < len(math1); i++ {
		for j := 0; j < len(math2); j++ {
			math3[i][j] = math1[i][j] + math2[i][j]
		}
		fmt.Println()
	}
	return math3
}

func First3Letters(str string) string {
	//create byte array
	var r [3]byte

	//read each character from string
	for i := 0; i < len(str); i++ {
		if i >= 3 {
			break
		}
		r[i] = str[i]
	}
	//return byte array as string
	return string(r[:])
}

/*
1. complete addition subtraction program
2. Pass an array and find max and min, average
3. pass an array and enter number print if number exits in a array
4. Fix reverse string program
5. Pass a string and check if it's a palindrome
6. Pass to dimension matrix, A and B and return A + B
*/
