package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

	fmt.Println("\n======== Guessing game: ========")
	GuessingGame()

	fmt.Println("\n======== Four operation: ========")
	sum, sub, mul, div := PerformOperation(10, 5)
	fmt.Printf("sum: %v sub: %v mul: %v div: %v\n", sum, sub, mul, div)

	fmt.Println("\n======== Printing fibonacci numbers: ========")
	findFibonacciSeries(7)

	fmt.Println("\n======== Printing sum of all digit: ========")
	output := FindSumOfAllDigits(63)
	fmt.Printf("sum of all digits: %v\n", output)

	fmt.Println("\n======== FizzBuzz: ========")
	ar := []int{30, 15, 7, 10, 15, 21, 22} //fizz = 1 , buzz =3  , fizzbuzz=2 ,  other = 1
	fizMap := fizzBuzz(ar)
	fmt.Printf("%v\n", fizMap)

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

// 3. fibonacci series
func findFibonacciSeries(num int) {
	num1 := 0
	num2 := 1
	var num3 int
	fib := []int{num1, num2}
	for i := 0; i < num; i++ {
		num3 = num1 + num2
		fib = append(fib, num3)
		//shift value of num1 and num2
		num1 = num2
		num2 = num3
	}
	fmt.Printf("fibonacci series of number: %v is %v\n", num, fib)

}

func FindSumOfAllDigits(num int) int {
	/*
		eg: 23 = 2+3 = 5
		236 = 2 + 3 + 6 = 11
		plan: loop through the each digit
		- using modulo operator gives the last digit which is going to add each digit
		- using divide operator which is going to update num by dividing number by 10
	*/
	var sum int

	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

// guessing game
func GuessingGame() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1 // random number from 1 - 100

	var (
		guessNumber  int
		numOfAttempt int
	)

	fmt.Println("Guess the number from 1 - 100")

	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%v", &guessNumber)

		numOfAttempt++

		if guessNumber > randomNumber {
			fmt.Println("Too high")
		} else if guessNumber < randomNumber {
			fmt.Println("Too low")
		} else {
			fmt.Printf("Congratulations! you guessed it in %v attempts\n", numOfAttempt)
			break
		}

	}

}

// create function which will return all four operation and no switch statement
func PerformOperation(num1 int, num2 int) (sum, sub, mul int, div float32) {

	sum = num1 + num2
	sub = num1 - num2
	for i := 1; i <= num2; i++ {
		mul += num1
	}
	/*
		i .   i >= num2   i = i-num2  div++
		15      T         15 - 5 = 10     1
		10      T .       10 - 5 = 5      2
		5		T		  5 - 5 = 0		  3
		0       F (loop end)
	*/
	for i := num1; i >= num2; i = i - num2 {
		div += 1
	}

	return sum, sub, mul, div
}

//fizzbuzz
/*
	if a number is divisible by 3, print value fizz
	if number is divisible by 5, print value buzz
	if number is divisible by 3 and 5, print fizz buzz
	plan:
	- create array
	- check which number associates with which string
	- link them in a map - key value pair
	- return map with count
*/
//3, 5, 7, 10, 15, 20, 30 - has some doubt
func fizzBuzz(arr []int) map[string]int {

	fizzbuzzMap := make(map[string]int)
	var fizz, buzz, fizzbuzz, other int

	for i := 0; i < len(arr); i++ {
		if arr[i]%3 == 0 {
			fizz++
			fizzbuzzMap["fizz"] = fizz

		} else if arr[i]%5 == 0 {
			buzz++
			fizzbuzzMap["buzz"] = buzz

		} else if arr[i]%3 == 0 && arr[i]%5 == 0 {
			fizzbuzz++
			fizzbuzzMap["fizzbuzz"] = fizzbuzz

		} else {
			other++
			fizzbuzzMap["other"] = other

		}
	}

	return fizzbuzzMap
}

/*
1. random number - guessing number => done
2. create a function, which will return all four operations and no switch statement => done
3. fibonacci series => done
4. fiz buzz, take numbers in an array and create function which will take input as parameter and return map (name, value pair - map)
5. Given number, sum of all digits => done
	- eg: 628 result should be 6+2+8
*/
