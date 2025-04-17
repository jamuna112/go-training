package main

import "fmt"

// func main() {
	//================== Map =================//
	// var a = map[string]string{
	// 	"brand": "Ford",
	// 	"model": "Mustang",
	// 	"year":  "1964",
	// }
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", a["brand"])

	// var b = make(map[string]int) //this is empty map, cannot initialize values but can have capacity
	// b["ford"] = 5
	// b["toyota"] = 2
	// fmt.Printf("%v\n", b)
	// fmt.Printf("number of ford available: %v\n", b["ford"])

	// var name, email, age string

	// fmt.Printf("Enter name: ")
	// fmt.Scanf("%v\n", &name)
	// userDetails["name"] = name

	// fmt.Printf("Enter email: ")
	// fmt.Scanf("%v\n", &email)
	// userDetails["email"] = email

	// fmt.Printf("Enter age: ")
	// fmt.Scanf("%v\n", &age)
	// userDetails["age"] = age

	// fmt.Printf("userdetails %v\n", userDetails)

	//================== Slice =================//

	//slice of map for all the students
	// students := []map[string]string{}

	// fmt.Printf("Enter name: ")
	// fmt.Scanf("%v\n", &name)

	// fmt.Printf("Enter email: ")
	// fmt.Scanf("%v\n", &email)

	// fmt.Printf("Enter age: ")
	// fmt.Scanf("%v\n", &age)

	//map for each student field
	// student := map[string]string{
	// 	"name":  name,
	// 	"email": email,
	// 	"age":   age,
	// }
	// students = append(students, student)
	// fmt.Printf("student details %v\n", students)

	//===============Runes===============//
	// str := "hello"
	// r := []rune(str)
	// fmt.Println(str)
	// fmt.Println(r)

	// for i, v := range r {
	// 	fmt.Printf("rune at position %v is %c\n", i, v)
	// }
	// fmt.Println("=======Passing a Function as an Argument=======")
	// result1 := applyOperation(4, 10, add)
	// fmt.Printf("operation result: %v\n", result1)

	// result2 := applyOperation(10, 3, sub)
	// fmt.Printf("operation result: %v\n", result2)

	//======= Basic pointer =======
	//1. Write a program that declares an integer variable and a pointer to it.Print both the value and the address
	var num int = 5
	var p *int = &num

	_, _ = p, num

	// fmt.Printf("num value is %d, memory address: %p\n", num, &num)
	// fmt.Printf("p value: %v, memory address: %p\n", *p, p)

	//2. Create a pointer to a string and use it to change the original string's value. (pointer derefencing)
	// var str string = "Greeting"
	// var ptr *string = &str
	// fmt.Printf("Initial string value: %v, memory address: %v\n", str, &str)
	// fmt.Printf("ptr is a pointer to str value: %v, memory address: %v\n", *ptr, ptr)

	// *ptr = "Greeting pointer"
	// fmt.Printf("after change, string value: %v, memory address: %v\n", str, &str)
	// fmt.Printf("after change, ptr value: %v, memory address: %v\n", *ptr, ptr)

	/*
		3. Nil Pointers. Declare a pointer without initialization and check if it's nil.
	*/
	// var ptr *int
	// fmt.Printf("ptr value without pointing initialization: %v\n", ptr) //nil

	/*
		4. Pointer as Function Parameter.
		Write a function that takes a pointer to an integer and doubles its value.
	*/
	// num1 := 5
	// doubleValue(&num1)
	// fmt.Printf("number: %v\n", num1)

	//5. swap two numbers
	x := 10
	y := 20

	fmt.Printf("Before swap: x = %v, y = %v\n", x, y)

	swapTwoNumbers(&x, &y)
	fmt.Printf("After swap: x = %v, y = %v\n", x, y)

}

// ======= Passing a Function as an Argument =======
// func applyOperation(a int, b int, operation func(int, int) int) int {
// 	return operation(a, b)
// }

// func add(num1 int, num2 int) int {
// 	return num1 + num2
// }
// func sub(num1 int, num2 int) int {
// 	return num1 - num2
// }

/*
4. Pointer as Function Parameter.
Write a function that takes a pointer to an integer and doubles its value.
*/
// func doubleValue(n *int) {
// 	*n = *n * 5
// }

// swap two intefers using pointers
// func swapTwoNumbers(a *int, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }
