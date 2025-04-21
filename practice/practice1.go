package main

// import "fmt"

// // import (
// // 	"fmt"
// // 	"strconv"
// // )

// // type newString string
// // type newInt int

// // type Books struct {
// // 	Name   string
// // 	Title  string
// // 	Author string
// // 	Cost   int
// // }

// // func main() {

// // 	var name string
// // 	var rollNum int
// // 	var name2 newString
// // 	var roll2 newInt

// // 	userDetails(&name, &rollNum)
// // 	msg := displayDetails(name, rollNum)
// // 	fmt.Printf("User details: %v\n", msg)

// // 	fmt.Printf("Please enter name: ")
// // 	fmt.Scanf("%s", &name2)

// // 	fmt.Printf("Please enter roll number: ")
// // 	fmt.Scanf("%d", &roll2)

// // 	fmt.Printf("name: %s, roll number: %d\n", name2, roll2)

// // 	book := Books{Name: "Java", Title: "JV", Author: "xyz", Cost: 12}
// // 	fmt.Printf("book details: %v\n", book)

// // 	book2 := Books{"Java", "JV", "xyz", 12}
// // 	fmt.Printf("book details: %v\n", book2.Name)

// // }

// // func userDetails(name *string, rollNum *int) {
// // 	fmt.Printf("Please enter name: ")
// // 	fmt.Scanf("%s", name)

// // 	fmt.Printf("Please enter roll number: ")
// // 	fmt.Scanf("%d", rollNum)
// // }

// // func displayDetails(name string, rollNum int) string {
// // 	fmt.Printf("name: %s, roll number: %d\n", name, rollNum)
// // 	return "Hello " + name + " your roll number is " + strconv.Itoa(rollNum)
// // }

// /*
// 1. create array of books, accept values from user and display
// 2. search for methods in golang - methods
// 3. difference between slice and pointer? read and practice, and make ppt
// 4. When to use make? advantange and disadvantage of using make. practice and make ppt
// 5. matrix multiplication program passing as a reference, make ppt
// 6. Insertion sort
// 7. fibonacci series on recursion // 0 , 1, 1, 2, 3, 5, 8
// */
// func main() {
// 	var num int
// 	fmt.Printf("Please enter number for fib number: ")
// 	fmt.Scanf("%d", &num)

// 	output := fib(num)
// 	fmt.Println("fib series:")
// 	fmt.Printf("%v, ", output)

// 	/*
// 		once user get input, fib function should return the sum of last two digit
// 		expected output: 0, 1, 1, 2, 3, 5, 8, 13, 21
// 	*/
// }
// func fib(n int) int {
// 	var a int = 0
// 	var b int = 1
// 	var c int

// 	for i := 0; i < n; i++ { //
// 		fmt.Printf("%v, ", a)
// 		c = a + b //1, 2, 3
// 		a = b     //1, 1,
// 		b = c     //1, 2
// 	}
// 	return c
// }
