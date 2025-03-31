package main // this is the main package

import "fmt" // fmt package for formattin

/*
	main is the entry point for any golang program

functions can be imported by refering their package
*/
func main() {
	fmt.Println("Hello World!")

	var i int      //initial value of i is 0 if not defined
	var j uint = 5 //assigned int value
	k := 6.0       // value is infered
	var l = 7      // value is infered

	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)

	fmt.Printf("i: %v, type is %T\n", i, i)
	fmt.Printf("j: %v, type is %T\n", j, j)
	fmt.Printf("k: %v, type is %T\n", k, k)
	fmt.Printf("l: %v, type is %T\n", l, l)

	//try
	var u8 uint8 = 7    // memory allocation is 2^8, 0 - 2^8-1 (2*2*2*2*2*2*2*2-1= 255)
	var u16 uint16 = 9  // memory allocation is 2^16, 0 - 2^16-1
	var u32 uint32 = 1  // memory allocation is 2^32, 0 - 2^32-1
	var u64 uint64 = 19 // memory allocation is 2^64, 0 - 2^64-1

	fmt.Printf("u8: %v, type is %T\nu16: %v, type is %T\nu32: %v, type is %T\nu64: %v, type is %T\n", u8, u8, u16, u16, u32, u32, u64, u64)

	//tried other uint data types
	u16 = 65535 // greater than 65535 will throw overflow error
	fmt.Printf("u18: %v, type is %T\n", u16, u16)

	u32 = 4294967295 // greater than 4294967295 will throw overflow error
	fmt.Printf("u32: %v, type is %T\n", u32, u32)

	u64 = 1844674407370955161 // greater than 18446744073709551615 will throw overflow error
	fmt.Printf("u64: %v, type is %T\n", u64, u64)

	var i8 int8 = -7   // memory allocation is 2^8, 0 - (2^8/2) -1 (-127 to +127)
	var i16 int16 = 9  // memory allocation is 2^16, 0 - 2^16-1
	var i32 int32 = 1  // memory allocation is 2^32, 0 - 2^32-1
	var i64 int64 = 19 // memory allocation is 2^64, 0 - 2^64-1

	fmt.Printf("\ni8: %v, type is %T\ni16: %v, type is %T\ni32: %v, type is %T\ni64: %v, type is %T\n", i8, i8, i16, i16, i32, i32, i64, i64)

	//tried other int data types
	i8 = -128 //-128 to 127
	fmt.Printf("i8: %v, type is %T\n", i8, i8)

	i16 = -32768 //-32768 to 32767
	fmt.Printf("i16: %v, type is %T\n", i16, i16)

	i32 = -2147483648 //-2147483648 to 2147483647
	fmt.Printf("i32: %v, type is %T\n", i32, i32)

	i64 = -9223372036 //-92233720368 to 92233720367
	fmt.Printf("i64: %v, type is %T\n", i64, i64)
}
