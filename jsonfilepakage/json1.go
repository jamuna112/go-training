package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func readJSON(filename string) []Person {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return nil
	}

	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	fmt.Println("Read JSON:", people)
	return people
}

func addJSONRecord(filename string, newPerson Person) {
	people := readJSON(filename)
	people = append(people, newPerson)

	fmt.Println("people data", people)

	data, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	os.WriteFile(filename, data, 0644)
	fmt.Println("Person added:", newPerson)
}

func modifyJSONRecord(filename string, id int, updated Person) {
	people := readJSON(filename)
	for i, p := range people {
		if p.Id == id {
			people[i] = updated
			break
		}
	}

	data, _ := json.MarshalIndent(people, "", "  ")
	os.WriteFile(filename, data, 0644)
	fmt.Println("Person updated:", updated)
}

func deletePerson(filename string, id int) {
	people := readJSON(filename)
	var per []Person

	for _, p := range people {
		if p.Id != id {
			per = append(per, p)
		}

	}

	data, err := json.MarshalIndent(per, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	os.WriteFile(filename, data, 0644)
	fmt.Println("person deleted")
}

func main() {
	fmt.Printf("Programming with json file:")
	filename := "person.json"
	per := readJSON(filename)
	fmt.Println(per)

	//p := Pet{Name: pet.Name, Age: pet.Age, Color: pet.Color}
	p := Person{1, "Peter", 30}

	_ = p

	//addJSONRecord(filename, p)

	//update := Person{2, "Bobby", 35}

	//modifyJSONRecord(filename, 2, update)

	deletePerson(filename, 2)

}

/*
storing this file
[
  {"id": 1, "name": "Alice", "age": 30},
  {"id": 2, "name": "Bob", "age": 25}
]
*/
