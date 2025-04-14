package main

import "fmt"

func main() {
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
	str := "hello"
	r := []rune(str)
	fmt.Println(str)
	fmt.Println(r)

	for i, v := range r {
		fmt.Printf("rune at position %v is %c\n", i, v)
	}

}
