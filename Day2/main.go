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

	num8 := num3 ^ num4 //xor, if one value is zero and another value is one then result is one. Otherwise zero
	fmt.Printf("num8: %6v, b: %6b\n", num8, num8)

	num3 <<= 1
	fmt.Printf("num3: %6v, b: %6b\n", num3, num3)

	num4 >>= 1
	fmt.Printf("num4: %6v, b: %6b\n", num4, num4)

	//Practice bitwise assignment operators
	//Bitwise AND assignment
	value1 := 5 // 0101
	value2 := 2 // 0010
	value3 := value1 & value2
	fmt.Printf("value 3 is %v, binary is: %b\n", value3, value3) //0000

	//Bitwise OR assignment
	value4 := value1 | value2
	fmt.Printf("value 4 is %v, binary is: %b\n", value4, value4) //0111

	//Bitwise XOR assignment
	value5 := value1 ^ value2
	fmt.Printf("value 5 is %v, binary is: %b\n", value5, value5) //0111

	//right shift assignment
	value6 := 9                                                  //1001
	value6 >>= 2                                                 //right shift by 2 position
	fmt.Printf("value 6 is %v, binary is: %b\n", value6, value6) //10

}
