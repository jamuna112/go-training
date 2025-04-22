package main

import "fmt"

/*
error:
Go uses error as a type to represent errors.
The standard way is to return an error value from a function and check if it is nil.

panic: panic stops the ordinary flow of control and begins panicking.
Useful for unexpected or unrecoverable situations.

defer:
defer is used to postpone the execution of a function until the surrounding function returns.
Commonly used to clean up resources like closing files, unlocking mutexes, etc.

recover:
recover is used to regain control of a panicking goroutine.
Only works inside a defer function
*/

// func divOperation() {

// 	for i := 0; i < 2; i++ {
// 		for j := 5; j < 7; j++ {
// 			fmt.Printf("Passing two numbers, n1: %d, n2: %v\n", j, i)
// 			remainder, err := division(j, i)
// 			if err != nil {
// 				fmt.Println("error occured ", err)
// 			} else {
// 				fmt.Printf("Remainder is: %d\n", remainder)
// 			}
// 		}
// 	}

// }

// func division(num1, num2 int) (int, error) {

// 	var remainder int
// 	if num2 == 0 {
// 		return 0, errors.New("div by zero")
// 	}
// 	remainder = num1 % num2
// 	return remainder, nil
// }

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 7 {
			//panic("number divisible by 7")
			fmt.Println("Seven")
		}
	}

	one()
	fmt.Println("Back in to main function")
}

func one() {
	fmt.Println("I have entered function 1")
	defer fmt.Println("In function one, this is second line") //LIFO, defer will be called after all the activities are run
	two()

}

func two() {
	fmt.Println("I have entered function 2")
	//panic("I am in function two")
	fmt.Println("In function two, this is second line")
	fmt.Println("In function two, this is third line")
	three()
}

func three() {
	fmt.Println("I have entered function 3")
	fmt.Println("In function three, this is second line")
	fmt.Println("In function three, this is third line")
	defer handlePanic()
	panic("In function one, this is third line")

}
func handlePanic() {

	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
