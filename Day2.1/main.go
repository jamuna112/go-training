package main

import "fmt"

func main() {

	var binaryNum = 0b1010 //camel case
	var HexNum = 0x1A      //pascal case
	var octal_num = 012    //snake case

	fmt.Printf("binary num: %v\n", binaryNum)
	fmt.Printf("hexa num: %v\n", HexNum)
	fmt.Printf("octal num: %v\n", octal_num)

	const NUMBER int = 5 //can't change the value of const

	//conditions
	value1 := 8
	value2 := 11

	if value1 > value2 {
		fmt.Println("Value 1 is greater")
	} else if value1 == value2 {
		fmt.Println("both are equal")
	} else {
		fmt.Println("value 2 is greater")
	}

	value3 := 15
	fmt.Println("Comparing three values", value1, value2, value3)
	if value1 > value2 && value1 > value3 {
		fmt.Println("Value 1 is greater")
	} else if value2 > value3 && value2 > value1 {
		fmt.Println("Value 2 is greater")
	} else if value3 > value1 && value3 > value1 {
		fmt.Println("Value 3 is greater")
	}

	if value1 > value2 {
		if value1 > value3 {
			fmt.Println("Value 1 is greater")
		} else {
			fmt.Println("Value 3 is greater")
		}
	} else {
		if value2 > value3 {
			fmt.Println("Value 2 is greater")
		} else {
			fmt.Println("Value 3 is greater")
		}
	}

	//if two/three values are same out of given three value, it should print and should print the greater value

	var b bool
	fmt.Println("value of b", b)
	b = value1 > value2
	fmt.Printf("value1: %v, value2: %v, value1 > value2: %v\n", value1, value2, b)
	b = value1 < value2
	fmt.Printf("value1: %v, value2: %v, value1 < value2: %v\n", value1, value2, b)

	var year int = 1999
	//var year int = 2000
	//var year int = 1900
	var yearLogic = year % 4
	fmt.Println("year logic", yearLogic)

	// if year%4 == 0 {
	// 	fmt.Println("It's a leap year")
	// } else {
	// 	fmt.Println("It's not a leap year")
	// }

	//if left condition is false program wil not check right condition in case of &&
	//if left condition is true program will not check right condition in case of ||
	if year%100 == 0 && year%400 != 0 {
		fmt.Println("It's not a leap year", year)
	} else if year%100 != 0 && year%4 == 0 {
		fmt.Println("It's a leap year", year)
	} else if year%400 == 0 {
		fmt.Println("It's a leap year", year)
	}

	//If year % 4 = 0 AND year % 100!= 0 OR year%400 == 0
	year = 1900
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("It's a leap year", year)
	} else {
		fmt.Println("It's not a leap year", year)
	}

	//switch statement
	var day int = 9

	switch day {
	case 0:
		fmt.Println("It is Sunday")
	case 1:
		fmt.Println("It is Monday")
	case 2:
		fmt.Println("It is Tuesday")
	case 3:
		fmt.Println("It is Wednesday")
	case 4:
		fmt.Println("It is Thursday")
	case 5:
		fmt.Println("It is Friday")
	case 6:
		fmt.Println("It is Saturday")
	default:
		fmt.Println("It is not a week day")

	}
	day = 5
	switch day {
	case 0, 6:
		fmt.Println("It is Weekend")
	case 1, 2, 3, 4, 5:
		fmt.Println("It is Week day")
	default:
		fmt.Println("Something went wrong")

	}

	//fizzbuzz
	/*
		if a number is divisible by 3, print value fizz
		if number is divisible by 5, print value buzz
		if number id divisible by 3 and 5, print fizz buzz
	*/

	/*
		Find a perfect square, given a number print if it is perfect square or not eg:
		25 is perfect square of 5 as 5 * 5 is 25
		26, 30.. are not perfect square
		36 is perfect square
	*/
}
