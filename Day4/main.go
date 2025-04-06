package main

import (
	"fmt"
	"math"
)

func main() {
	// 	/*
	// 	   	c.	for {}		// for ever, infinite loop
	// 	   d.	for i,j:=range array_variable {
	// 	   i // is for index
	// 	   j // is for value
	// 	   }
	// 	   e.	for condition {} – loop runs till condition is true
	// 	   f.	break – to exit loop
	// 	   g.	continue – to continue loop
	// 	*/
	// 	fmt.Println("Using for loop:")
	// 	for i := 1; i < 10; i++ {
	// 		fmt.Println(i, "Hello, Go!")
	// 	}
	// 	fmt.Println("Using for loop only with conditions:")
	// 	i := 1
	// 	for i < 10 {
	// 		fmt.Println(i, "Hello, Go!")
	// 		i++
	// 	}

	// 	//Write a program to generate a loop from 1-10 and print each number if it is even or odd
	// 	for i := 0; i <= 10; i++ {
	// 		if i%2 == 0 {
	// 			fmt.Println(i, "Even number")
	// 		} else {
	// 			fmt.Println(i, "Odd number")
	// 		}
	// 	}

	// 	//Write a program to generate a loop from 1-10 and print statement only used once
	// 	fmt.Println("Program to display even or odd in a single print")
	// 	var evenNumber string
	// 	for i := 0; i <= 10; i++ {
	// 		if i%2 == 0 {
	// 			evenNumber = "even"
	// 		} else {
	// 			evenNumber = "odd"
	// 		}
	// 		fmt.Printf("%v, %v number\n", i, evenNumber)
	// 	}

	// 	//create a loop of 10 numbers and exit after printing 4 numbers
	// 	fmt.Println("create a loop of 10 numbers and exit after printing 4 numbers")
	// 	for i, j := 10, 0; i >= 1; i-- {
	// 		fmt.Printf("value of i: %v, j: %v\n", i, j)
	// 		j++
	// 		if j == 4 {
	// 			break
	// 		}

	// 	}

	// 	var i1, j1, k1 int = 10, 11, 12
	// 	fmt.Printf("value of i1: %v, j1: %v, k1: %v\n", i1, j1, k1)

	// 	a1, b1, c1, d11 := 1.52, "abc", 100, 7.2+10i
	// 	fmt.Printf("value of a1: %v, type: %T\n", a1, a1)
	// 	fmt.Printf("value of b1: %v, type: %T\n", b1, b1)
	// 	fmt.Printf("value of c1: %v, type: %T\n", c1, c1)
	// 	fmt.Printf("value of d11: %v, type: %T\n", d11, d11)

	// 	var test string
	// 	var test2 string = "xyz"
	// 	test3 := "xyz2"
	// 	fmt.Printf("value of test: %v\n", test)
	// 	fmt.Printf("value of test2: %v\n", test2)
	// 	fmt.Printf("value of test3: %v\n", test3)

	// 	var c4, c2, c3 string = "a", "b", "c"
	// 	fmt.Printf("value of c4: %v\n", c4)
	// 	fmt.Printf("value of c2: %v\n", c2)
	// 	fmt.Printf("value of c3: %v\n", c3)

	// 	//comparing with ASCII values
	// 	if c4 > c2 {
	// 		fmt.Printf("greater value of %v and %v is %v\n", c4, c2, c4)
	// 	} else {
	// 		fmt.Printf("greater value of %v and %v is %v\n", c4, c2, c2)
	// 	}

	// 	//var d1, d2, d3 string = "apple", "ball", "cat"
	// 	var d1, d2, d3 string = "bat", "ball", "cat"
	// 	fmt.Printf("value of d1: %v\n", d1)
	// 	fmt.Printf("value of d2: %v\n", d2)
	// 	fmt.Printf("value of d3: %v\n", d3)

	// 	//comparing with ASCII values
	// 	if d1 > d2 {
	// 		fmt.Printf("greater value of %v and %v is %v\n", d1, d2, d1)
	// 	} else {
	// 		fmt.Printf("greater value of %v and %v is %v\n", d1, d2, d2)
	// 	}

	// 	//given ASCII value
	// 	for i := 0; i < len(d1); i++ {
	// 		fmt.Println(d1[i])
	// 	}

	// 	for i, j := range d1 {
	// 		fmt.Printf("%v - %c\n", i, j)
	// 	}
	// 	fmt.Printf("%v, length: %d\n", d1, len(d1))
	// 	fmt.Printf("%v, length: %d\n", d2, len(d2))
	// 	fmt.Printf("%c, %c, %c %d\n", d1[0], d1[1], d1[2], len(d1))

	// 	var arr [3]int                       // declaring an array default value to 0
	// 	var arr1 = [3]int{3, 5, 8}           // declaring array with define value
	// 	arr2 := [3]int{4, 5, 6}              // declaring array with define value
	// 	arr3 := [...]int{10, 11, 12, 13, 14} // declaring array with define values as size of array is inferred
	// 	fmt.Printf("%v, length: %d\n", arr, len(arr))
	// 	fmt.Printf("%v, length: %d\n", arr1, len(arr1))
	// 	fmt.Printf("%v, length: %d\n", arr2, len(arr2))
	// 	fmt.Printf("%v, length: %d\n", arr3, len(arr3))

	// 	for i, j := range arr3 {
	// 		fmt.Printf("i: %v, j: %v\n", i, j)
	// 	}

	// 	var serr = [5]string{"apple", "cat", "dog"}

	// 	for i, j := range serr {
	// 		fmt.Printf("i: %v, j: %v\n", i, j)
	// 	}
	// 	serr[3] = "orange"
	// 	serr[4] = "mango"
	// 	//serr[5] = "pear"
	// 	for _, j := range serr {
	// 		fmt.Printf("j: %v\n", j)
	// 	}

	// 	var str1 string = "variable1"
	// 	//fmt.Println(str1)

	// 	_ = str1 //we are ignoring the variable str1

	// 	fmt.Println("=========================================")
	// 	/*
	// 		write a program, create a loop of 10 numbers, create another loop of 10 strings which will
	// 		tell odd or even of number array.
	// 		Eg: n = {12, 11, 10} // this is input
	// 		s = {even, odd, even} //this should be the output of array
	// 		actual output on screan
	// 		12 - even
	// 		11 - odd
	// 		10 - even
	// 	*/
	// 	givenInput := [10]int{12, 11, 10, 22, 19, 32, 4, 30, 1, 51}
	// 	var outputStr [10]string
	// 	for i := 0; i < 10; i++ {
	// 		if givenInput[i]%2 == 0 {
	// 			outputStr[i] = "even"
	// 		} else {
	// 			outputStr[i] = "odd"
	// 		}
	// 		//fmt.Printf("%v - %v\n", givenInput[i], outputStr)
	// 	}
	// 	for i := 0; i < len(givenInput); i++ {
	// 		fmt.Printf("Given number: %v - %v\n", givenInput[i], outputStr[i])
	// 	}
	// 	fmt.Println("=========================================")

	/*
		Given a number, first bit if it is 1 it is a positive number, if first bit is 0 it is negative number
		Eg: n = 1010 = +2
		n = 0111 = -7
		example 100100
		0110110
	*/

	givenBitNum := "10010"
	firstDigit := givenBitNum[0]
	number := givenBitNum[1:]
	var n int
	fmt.Printf("Given bit number: %v\n", givenBitNum)
	//fmt.Printf("First bit number: %b\n", firstDigit)
	fmt.Printf("Rest number: %v\n", number)
	var valueSign int
	_ = valueSign

	if firstDigit == '1' {
		valueSign = +1
	} else {
		valueSign = -1
	}
	var result int
	//for i := len(number) - 1; i >= 0; i-- {

	for i := 0; i <= len(number)-1; i++ {
		if number[i] == '0' {
			n = 0
		} else {
			n = 1
		}
		fmt.Printf("%v ", n)
		//result += int(math.Pow(2.0, float64(i))) * n
		result += n * int(math.Pow(2.0, float64(i)))
	}
	fmt.Println()
	fmt.Println("Result:", result)
	fmt.Printf("Result of given number is %v\n", result+valueSign)

}
