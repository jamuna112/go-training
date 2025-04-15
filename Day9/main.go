package main

import "fmt"

func main() {
	var arr []int
	var arrNum, num int
	fmt.Println("How many numbers you want to enter")
	fmt.Scanf("%d", &arrNum)
	fmt.Println("Enter numbers to find min and max: ")
	for i := 0; i < arrNum; i++ {
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}

	min, max := findMinAndMax(arr)
	fmt.Printf("Min: %v, Max: %v\n", min, max)

	min2, max2 := findMinAndMax2(arr)
	fmt.Printf("Min: %v, Max: %v\n", min2, max2)

	fmt.Printf("given array %v\n", arr)
	result := SumOfNumArrays(arr)
	fmt.Printf("sum of given arrray's number: %v\n", result)

	sumOfArrays := SumOfNumArrays
	result2 := sumOfArrays(arr)
	fmt.Printf("result2:  %v, sum of arrays: %T\n", result2, sumOfArrays)

	morning := Greet
	evening := Greet
	fmt.Printf("%v, data types: %T\n", morning("John"), morning)
	fmt.Printf("%v, data types: %T\n", evening("Joe"), evening)

	//anonymous function - check this out
	sumOfNumbers := func(num1 int, num2 int) int {
		return num1 + num2
	}
	fmt.Println("sum of numbers 2, 3", sumOfNumbers(2, 3))

	//factorial
	FactorialRec(3)
	//recursion
	result3 := FactorialRec1(4)
	fmt.Printf("recursion of number 4 is %v\n", result3) // 4*3*2*1
}

/*
read number of an array from user and return min and max
*/

func findMinAndMax(array []int) (int, int) {
	var min, max int = array[0], array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
		if array[i] > max {
			max = array[i]
		}

	}

	return min, max
}

func findMinAndMax2(array []int) (min int, max int) {
	min, max = array[0], array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
		if array[i] > max {
			max = array[i]
		}

	}

	return
}

// Write a function to sum of numbers of arrays
func SumOfNumArrays(array []int) int {
	var sum int = 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

func Greet(str string) string {
	return "Greetings to" + str
}

/*
1. Write a recursion program for factorial
5! = 5 * 4 * 3 * 2 * 1
3! = 3 * 2 * 1
*/

func FactorialRec(num int) {
	var output int = 1
	for i := 1; i <= num; i++ {
		output *= i
	}
	fmt.Printf("Output of given num : %v\n", output)
}

// recursion
func FactorialRec1(num int) int {
	if num <= 1 {
		return 1
	}
	return num * FactorialRec1(num-1)
}

/*
1. random number - guessing number
2. create a function, which will return all four operations and no switch statement
3. fibonacci series
4. fiz buzz, take numbers in an array and create function which will take input as parameter and return map (name, value pair - map)
5. Given number, sum of all digits
	- eg: 628 result should be 6+2+8
*/
