package main

import (
	"fmt"
)

//1. create array of books, accept values from user and display

// /*
// - store all the student details in an array
// - ask user how many books they want to store
// - take the input and loop through getting all the details
// - add those details to array
// - now display those details, by looping through the array
// */
type Books struct {
	Name         string
	Author       string
	ReleasedYear int
	Cost         float32
}

// func main() {
// 	var size, year int
// 	var cost float32
// 	fmt.Printf("How many books you want to store: ")
// 	fmt.Scanf("%v", &size)

// 	bookSlice := make([]Books, 0, size)

// 	reader := bufio.NewReader(os.Stdin)

// 	for i := 0; i < size; i++ {
// 		fmt.Printf("Enter book name: ")
// 		name, _ := reader.ReadString('\n')
// 		name = strings.TrimSpace(name)

// 		fmt.Printf("Enter book author: ")
// 		authorName, _ := reader.ReadString('\n')
// 		authorName = strings.TrimSpace(authorName)

// 		fmt.Printf("Enter book released year: ")
// 		yearStr, _ := reader.ReadString('\n')
// 		yearStr = strings.TrimSpace(yearStr)
// 		year, _ = strconv.Atoi(yearStr)

// 		fmt.Printf("Enter book cost: ")
// 		costStr, _ := reader.ReadString('\n')
// 		costStr = strings.TrimSpace(costStr)
// 		costFloat, _ := strconv.ParseFloat(costStr, 32)
// 		cost = float32(costFloat)

// 		book := Books{
// 			Name:         name,
// 			Author:       authorName,
// 			ReleasedYear: year,
// 			Cost:         cost,
// 		}

// 		bookSlice = append(bookSlice, book)

// 	}

// 	for i, k := range bookSlice {
// 		fmt.Printf("%d - %v\n", (i + 1), k)
// 	}

// 	// // 	//===============Fib series=============//
// 	// // 	// fmt.Println("=========================")
// 	// // 	// var n int
// 	// // 	// fmt.Print("Enter number n:")
// 	// // 	// fmt.Scanf("%v", &n)

// 	// // 	// fmt.Println("Fibonacci series: ")
// 	// // 	// for i := 0; i < n; i++ {
// 	// // 	// 	fmt.Print(fibonacciSeries(i), " ")
// 	// // 	// }
// 	// // 	//====================================//
// 	comp1 := Computer{name: "Sony"}
// 	fmt.Printf("\nInital computer name: %v, memory address: %p\n", comp1, &comp1)

// 	comp1.changeName("LG")
// 	fmt.Printf("after update computer name: %v, memory address: %p\n", comp1, &comp1)

// 	comp1.updateName("Mac")
// 	fmt.Printf("change: %v\n", comp1)
// 	fmt.Printf("value receive, name: %v, memory address: %p\n", comp1, &comp1)

// }

// 2. fibonacci series on recursion
// func fibonacciSeries(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fibonacciSeries(n-1) + fibonacciSeries(n-2)

// }

// methods in golang - methods
// Pointer receiver
type Computer struct {
	name string
}

// pointer receiver method
func (comp *Computer) changeName(newName string) {
	comp.name = newName
}

// value receiver method
func (comp Computer) updateName(newName string) {
	comp.name = newName
	fmt.Printf("change name: %v, memory address: %p\n", comp.name, &comp)
}
