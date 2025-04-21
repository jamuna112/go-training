package main

/*
	Practice problem, compare slice and pointers
*/

// func main() {
// 	sliceNum := []int{1, 2, 3}
// 	fmt.Printf("Given Slice: %v, memory add: %p\n", sliceNum, &sliceNum)
// 	modifyUsingSlice(sliceNum)
// 	fmt.Printf("Slice after modify slice: %v, memory add: %p\n", sliceNum, &sliceNum)

// 	x := 20
// 	fmt.Printf("Given X: %v, memory add: %p\n", x, &x)
// 	modifyUsingPointer(&x)
// 	fmt.Printf("X after modify pointer: %v, memory add: %p\n", x, &x)

// 	person := Person{name: "Original"}
// 	fmt.Printf("Initial person name: %v, memory add: %p\n", person, &person)

// 	changeNamePtr(&person)
// 	fmt.Printf("person name via pointer: %v, memory add: %p\n", person, &person)

// 	person2 := []Person{{name: "Original"}}
// 	fmt.Printf("person name via slice: %v, memory add: %p\n", person2, &person2)
// 	changeNameSlice(person2)
// 	fmt.Printf("person name via slice: %v, memory add: %p\n", person2[0].name, &person2)

// 	s := []int{1, 2, 3}
// 	fmt.Printf("initial value of s: %v, memory address: %p\n", s, &s)
// 	growWithPointer(&s)
// 	fmt.Printf("Modified slice s: %v, memory address: %p\n", s, &s)

// }

func modifyUsingSlice(s []int) {
	s[0] = 100
}

func modifyUsingPointer(p *int) {
	*p = 200
}

// pass slice vs pointer to struct
type Person struct {
	name string
}

// pass by pointer
func changeNamePtr(p *Person) {
	p.name = "Joe"
}

// pass by slice of struct
func changeNameSlice(s []Person) {
	s[0].name = "Sam"
}

// Pass a pointer to slice (real change) note: if it's not pointer value won't update
func growWithPointer(ps *[]int) {
	*ps = append(*ps, 99)
}

/*
Didn't see much difference in slice and in pointer. Both update the data directly in the memory address.
- slice can grow in size with append func wheareas pointer cannot grow, unless pointing to slice
*/
