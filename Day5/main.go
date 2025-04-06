package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [5]int{2, 4, 5, 6} // arrays are fixed
	arr1 := [5]int{1: 10, 3: 20}

	fmt.Printf("Checking array - arr:\n")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("i: %v\n", arr[i])
	}
	fmt.Printf("Checking array - arr1\n")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("i: %v\n", arr1[i])
	}

	// //slices are not fixed
	fmt.Printf("Checking array - arr2\n")
	arr2 := []int{12, 43, 50, 60}
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("i: %v\n", arr2[i])
	}
	fmt.Printf("Initial arr2: %v, length: %d, capacity: %d\n", arr2, len(arr2), cap(arr2))

	//array of string
	fmt.Printf("Checking array of string - arrStr\n")
	var arrStr = [3]string{"toyota", "Hundai", "Audi"}
	for i := 0; i < len(arrStr); i++ {
		fmt.Printf("i: %v\n", arrStr[i])
	}

	//append values in slice
	arr2 = append(arr2, 99, 01)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("after appending arr2: %v, length: %d, capacity: %d\n", arr2, len(arr2), cap(arr2))

	fmt.Printf("original array value of arr: %v\n", arr)
	var newArr = arr[1:]
	fmt.Printf("new array value of arr: %v\n", newArr)

	var newArr2 = arr[:3]
	fmt.Printf("new array value of newArr2: %v\n", newArr2)

	var newArr3 = arr[1:3]
	fmt.Printf("new array value of newArr3: %v\n", newArr3)

	/*
		a = [3][4]int{
	   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
	// {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
	//{8, 9, 10, 11}   /*  initializers for row indexed by 2 */

	var math1 = [2][2]int{{2, 3}, {8, 9}}
	fmt.Printf("math1 value: %v, length: %d\n", math1, len(math1))
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	/* output each array element's value */
	fmt.Println("Matrix math1:")
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("%4d", a[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Matrix math2 and math3:")
	var math2 = [3][3]int{{10, 20, 30}, {40, 50, 60}, {70, 80, 90}}
	var math3 = [3][3]int{{100, 200, 300}, {400, 500, 600}, {700, 800, 900}}

	fmt.Println("Matrix math2")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%4d", math2[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Matrix math3:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%4d", math3[i][j])
		}
		fmt.Println()
	}

	/*
		Create a new variable for math4 which will have addition of math2 and math 3 - matrix addition
	*/

	//Map
	//var c1 map[string]string

	fmt.Println("Map data types:")

	var c2 = map[string]string{

		"USA":    "+1",
		"India":  "+91",
		"China":  "+86",
		"Brazil": "+55",
		"UK":     "+44",
	}
	var c3 = map[string]int{
		"USA":    +1,
		"India":  +91,
		"China":  +86,
		"Brazil": +55,
		"UK":     +44,
	}
	for k, v := range c2 {
		fmt.Printf("k:%v, v:%v\n", k, v)
	}

	fmt.Printf("ISD code of USA: %v\n", c2["USA"])

	fmt.Printf("ISD code of Brazil: %v\n", c3["Brazil"])

	var c4 = map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
	}

	for i, j := range c4 {
		fmt.Printf("Roman value of %v, is %v\n", i, j)
	}
	/*
		From a Given number print roman number
	*/
	// fmt.Println("Convert given number to roman number")
	// var num int = 10
	// romanNumberMap := map[int]string{
	// 	100: "C",
	// 	50:  "L",
	// 	10:  "X",
	// 	9:   "IX",
	// 	5:   "V",
	// 	4:   "IV",
	// 	1:   "I",
	// }
	// result := ""
	// for key, value := range romanNumberMap {
	// 	fmt.Printf("number: %v, roman: %v\n", key, value)
	// 	for num >= key {
	// 		fmt.Printf("value: %v\n", value)
	// 		result += romanNumberMap[key]
	// 		num -= key
	// 	}
	// }
	// fmt.Printf("Result: %v\n", result)

	/*
		Given array of number, create a map
		keys: Even number, odd, perfect square, prime number
		for each of these keys I want count values
		a = {1, 2, 3, 4} //given array
		v = {"even": 2, "odd":2, "perfect square": 1, "prime number": 2 } //output should be in map
	*/

	givenArray := [6]int{4, 5, 9, 8, 6, 11}
	resultMap := map[string]int{
		"even":           0,
		"odd":            0,
		"perfect square": 0,
		"prime number":   0,
	}

	var squareNum int
	_ = resultMap

	//count of even and odd
	for i := 0; i < len(givenArray); i++ {
		if givenArray[i]%2 == 0 {
			resultMap["even"]++
		} else {
			resultMap["odd"]++
		}

		//count of square number
		squareNum = int(math.Sqrt(float64(givenArray[i])))
		if givenArray[i] == squareNum*squareNum {
			resultMap["perfect square"]++
		}

		//count of prime number
		if isPrime(givenArray[i]) {
			resultMap["prime number"]++
		}

	}
	fmt.Println(resultMap)

}

func isPrime(number int) bool {

	if number <= 1 || number%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
