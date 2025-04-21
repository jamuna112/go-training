package main

//Student grade tracker
/*
	Store a list of students (name + grade).
	Add a new student.
	Update a student’s grade using both:
	A pointer to a student.
	A slice of students.
	Print all student data.

	Plan:
	- for student details - name and grade, create Student struct
	- to store a list of student, create empty slice
	- to update student's grade using pointer, create func pointing to student
	- to update student's grade using slice, pass slice parameter
	- Print student data, using loop
*/

type Student struct {
	name  string
	grade int
}

// add student using slice by value
func addStudent(std *[]Student, name string, grade int) {
	student := Student{name: name, grade: grade}

	*std = append(*std, student)
}

// update grade using pointer
/*
NOTE: If I don't use pointer to student, then update is happen locally meaning only inside of this function
if, I call this function in main, change won't take effect. To make the change, we would need to work with
pointer
*/
func updateGradePtr(std *Student, newGrade int) {
	std.grade = newGrade
}

// update using slice index
func updateGradeSlice(std []Student, index int, newGrade int) {
	std[index].grade = newGrade
}

// func main() {

// 	student := []Student{
// 		{name: "Joe", grade: 80},
// 		{name: "Frank", grade: 60},
// 	}
// 	fmt.Printf("Initial student: %v, memory address: %p\n", student, &student)
// 	addStudent(&student, "Jassy", 90)
// 	fmt.Printf("adding student: %v, memory address: %p\n", student, &student)

// 	//update func using pointer
// 	updateGradePtr(&student[0], 70)
// 	fmt.Printf("updating student's grade using pointer: %v, memory address: %p\n", student, &student)

// 	//update func using slice index
// 	updateGradeSlice(student, 1, 40)
// 	fmt.Printf("updating student's grade using slice index : %v, memory address: %p\n", student, &student)

// }
