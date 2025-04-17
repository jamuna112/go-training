package main

import "fmt"

func main() {
	// arr := [4]int{3, 4, 5, 6}

	// map1 := map[string]int{
	// 	"three": 3,
	// 	"four":  4,
	// 	"five":  5,
	// 	"six":   6,
	// }
	// fmt.Printf("arr data type: %T\n", arr)
	// fmt.Printf("arr: %v\n", arr)
	// fmt.Printf("map data type: %T\n", map1)
	// for i, v := range map1 {
	// 	fmt.Printf("map key: %v - value: %v\n", i, v)

	// }

	// /*
	// 	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	// */
	// map2 := make(map[string]int)
	// map2["one"] = 1
	// map2["two"] = 2

	// fmt.Printf("map2 data type: %T\n", map2)
	// for i, k := range map2 {
	// 	fmt.Printf("%v - %v\n", i, k)
	// }

	// arr2 := make([]int, 6, 8)
	// arr2[0] = 9
	// arr2[1] = 8
	// arr2[2] = 7

	// fmt.Printf("arr2: %v, data type: %T\n", arr2, arr2)

	// newArr := []int{7, 4, 9, 2, 0}
	// newSlice := newArr

	// fmt.Printf("new array: %v, length: %v, capacity: %v\n", newArr, len(newArr), cap(newArr))
	// fmt.Printf("new slice: %v, length: %v, capacity: %v\n", newSlice, len(newSlice), cap(newSlice))
	// fmt.Printf("new array of 0 index: %v, %p\n", newArr[0], &newArr[0])
	// fmt.Printf("new array of 1 index: %v, %p\n", newArr[1], &newArr[1])

	// fmt.Printf("new slice of 0 index: %v, %p\n", newSlice[0], &newSlice[0])
	// fmt.Printf("new slice of 1 index: %v, %p\n", newSlice[1], &newSlice[1])

	// newSlice[0] = 100
	// fmt.Printf("new array: %v, length: %v, capacity: %v\n", newArr, len(newArr), cap(newArr))
	// fmt.Printf("new slice: %v, length: %v, capacity: %v\n", newSlice, len(newSlice), cap(newSlice))

	// fmt.Printf("new array of 0 index: %v, %p\n", newArr[0], &newArr[0])
	// fmt.Printf("new array of 1 index: %v, %p\n", newArr[1], &newArr[1])

	// fmt.Printf("new slice of 0 index: %v, %p\n", newSlice[0], &newSlice[0])
	// fmt.Printf("new slice of 1 index: %v, %p\n", newSlice[1], &newSlice[1])

	// newSlice = append(newSlice, 200)

	// fmt.Printf("new array: %v, length: %v, capacity: %v\n", newArr, len(newArr), cap(newArr))
	// fmt.Printf("new slice: %v, length: %v, capacity: %v\n", newSlice, len(newSlice), cap(newSlice))

	// fmt.Printf("new array of 0 index: %v, %p\n", newArr[0], &newArr[0])
	// fmt.Printf("new slice of 0 index: %v, %p\n", newSlice[0], &newSlice[0])

	// var n *int = new(int)
	// var n1 *int
	// fmt.Println("n:", *n)
	// fmt.Println("n:", n1)

	var number1, number2 int
	var num1 *int = new(int)
	fmt.Println("Enter first number")
	fmt.Scanf("%v", num1)
	fmt.Println("Enter second number")
	fmt.Scanf("%v", &number2)

	fmt.Printf("First num: %v, second num: %v\n", number1, number2)
	fmt.Printf("First num: %v, second num: %v\n", *num1, number2)

	fmt.Println("value of num2", number2)
	doubleNum(number2)
	fmt.Println("value of num2 after func call", number2)

	fmt.Println("value of num2", number2)
	doubleNumForPointer(&number2)
	fmt.Println("value of num2 after func call", number2)

}

/*
write a func where you will pass a number and double that number
*/

func doubleNum(num int) {
	fmt.Println("in fumc receive ", num)
	num = num * 2
	fmt.Println("in func generated", num)

}

func doubleNumForPointer(num *int) {
	fmt.Println("in fumc receive ", *num)
	*num = *num * 2
	fmt.Println("in func generated", *num)

}
