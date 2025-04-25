package errorhandling

import "fmt"

/*
write a program to handle defer and recover to catch the panic and handle it.
1. create array with few elements
2. create function to access array using index, if index is out of bounds, trigger a panic
3. use defer and recover() to catch the panic and handle it
4. print a message and continue execution

*/

func AccessingArrayElement(arr []int, index int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)

		}
	}()

	fmt.Printf("accessing index %d of the array\n", index)
	fmt.Printf("Array element: %d\n", arr[index])

}
