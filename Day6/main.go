package main

import "fmt"

func main() {
	var i int
	var pi *int
	pi = &i

	fmt.Printf("Initial value of i:\n")
	fmt.Printf("Value of i: %v, address of i: %p\n", i, &i)
	fmt.Printf("Value of *pi: %v, address pointed by pi: %p\n", *pi, pi)
	i = 10

	fmt.Printf("When value of i is 10:\n")
	fmt.Printf("Value of i: %v, address of i: %p\n", i, &i)
	fmt.Printf("Value of *pi: %v, address pointed by pi: %p\n", *pi, pi)

	fmt.Printf("When value of *pi is 100:\n")
	*pi = 100
	fmt.Printf("Value of i: %v, address of i: %p\n", i, &i)
	fmt.Printf("Value of *pi: %v, address pointed by pi: %p\n", *pi, pi)

	arr1 := [10]int{2, 4, 5, 7, 1, 3, 9, 23, 6, 17}
	fmt.Printf("%v ", arr1)
	parr1 := &arr1
	//address of arr1 and parr1
	fmt.Printf("\naddress of arr1: %p ", &arr1)
	fmt.Printf("\naddress of parr1: %p ", parr1)

	fmt.Printf("\nvalue of arr1: %v ", arr1)
	fmt.Printf("\nvalue of parr1: %v ", *parr1)

	//print index of arr1[1]
	fmt.Printf("\nvalue of arr1[1]: %v \n", arr1[1])
	fmt.Printf("value of parr1[1]: %v \n", (*parr1)[1])

	array2 := []int{32, 45, 12}
	pointerArray2 := &array2

	//address of arr1 and parr1
	fmt.Printf("\naddress of array2: %p ", &array2)
	fmt.Printf("\naddress of pointerArray2: %p ", pointerArray2)

	fmt.Printf("\nvalue of array2: %v ", array2)
	fmt.Printf("\nvalue of pointerArray2: %v ", pointerArray2)

	//pointer to arr1
	(*parr1)[1] = 100
	fmt.Printf("value of arr1: %v\n", arr1)
	//pointer to array2
	(*pointerArray2)[1] = 200
	fmt.Printf("value of array2: %v\n", array2)
	fmt.Printf("====== declaring new variable n1 for arr1: ======\n")
	n1 := arr1
	fmt.Printf("value of arr1: %v, data type: %T\n", arr1, arr1)
	fmt.Printf("value of n1: %v, data type: %T\n", n1, n1)

	fmt.Printf("====== declaring new variable n2 for slice array2: ======\n")
	n2 := array2
	fmt.Printf("value of array2: %v, data type: %T\n", array2, array2)
	fmt.Printf("value of n2: %v, data type: %T\n", n2, n2)

	n1[0] = 56
	n2[0] = 99
	fmt.Printf("====== updating n1 and n2 index 0: ======\n")
	fmt.Printf("value of arr1: %v, data type: %T\n", arr1, arr1)
	fmt.Printf("value of n1: %v, data type: %T\n", n1, n1)
	fmt.Printf("value of array2: %v, data type: %T\n", array2, array2)
	fmt.Printf("value of n2: %v, data type: %T\n", n2, n2)

	array2[2] = 121
	arr1[2] = 131
	fmt.Printf("====== updating array2 and arr1 index 2: ======\n")
	fmt.Printf("value of arr1: %v, data type: %T\n", arr1, arr1)
	fmt.Printf("value of n1: %v, data type: %T\n", n1, n1)
	fmt.Printf("value of array2: %v, data type: %T\n", array2, array2)
	fmt.Printf("value of n2: %v, data type: %T\n", n2, n2)

	var number int
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &number)
	fmt.Printf("number: %v\n", number)

}
