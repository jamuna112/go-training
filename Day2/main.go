package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Welcome to GO!")

	var num1 int = 10
	var num2 int = 7
	fmt.Printf("num1: %6v, b: %6b, o: %6o , h: %6x\n", num1, num1, num1, num1) // positive value for left alignment
	fmt.Printf("num2: %6v, b: %6b, o: %6o , h: %6x\n", num2, num2, num2, num2)

	num1++
	num2 += 5
	fmt.Printf("num1: %-6v, b: %-6b, o: %-6o , h: %-6x\n", num1, num1, num1, num1) //negative value for right alignment
	fmt.Printf("num2: %-6v, b: %-6b, o: %-6o , h: %-6x\n", num2, num2, num2, num2)

	var u8 uint8 = 255
	fmt.Printf("u8: %0v, b: %010b, o: %08o, h: %08x\n", u8, u8, u8, u8) //proper alingment for numbers

	/*
			&=	x &= 3	x = x & 3
		|=	x |= 3	x = x | 3
		^=	x ^= 3	x = x ^ 3
		>>=	x >>= 3	x = x >> 3
		<<=	x <<= 3	x = x << 3
	*/

	var num3 = 4
	var num4 = 6
	//num5 := num3 + num4
	fmt.Printf("num3: %6v, b: %6b\n", num3, num3)
	fmt.Printf("num4: %6v, b: %6b\n", num4, num4)
	//fmt.Printf("num5: %6v, b: %6b\n", num5, num5)

	// num6 := num3 & num4
	// fmt.Printf("num6: %6v, b: %6b\n", num6, num6)

	// num7 := num3 | num4
	// fmt.Printf("num6: %6v, b: %6b\n", num7, num7)

	num8 := num3 ^ num4 //xor, if one is value is zero and another value is one then result is one. Otherwise zero
	fmt.Printf("num8: %6v, b: %6b\n", num8, num8)

	num3 <<= 1
	fmt.Printf("num3: %6v, b: %6b\n", num3, num3)

	num4 >>= 1
	fmt.Printf("num4: %6v, b: %6b\n", num4, num4)
}
