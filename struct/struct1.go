package main

import "fmt"

/*
	Creating pet record system that stores and displays the pet information using struct
	1. create Pet struct
	2. create a slice to hold pets
	3. add pets, create a function which will store pets.

*/

type Pet struct {
	Name  string
	Age   int
	Color string
}

func addPets(pet Pet, pets []Pet) []Pet {
	p := Pet{Name: pet.Name, Age: pet.Age, Color: pet.Color}
	pets = append(pets, p)

	return pets
}

func listOfPets(pets []Pet) {
	fmt.Println("=============================")
	fmt.Println("List of pets: ")
	for i, v := range pets {
		fmt.Printf("%v. Pet name: %s, age: %d, color: %s\n", (i + 1), v.Name, v.Age, v.Color)
	}
	fmt.Println("=============================")
}

func main() {
	petSlice := []Pet{}
	var numOfPets, petAge int
	var petName, petColor string
	fmt.Print("How many pets you want to store? ")
	fmt.Scanf("%v", &numOfPets)

	for i := 0; i < numOfPets; i++ {
		fmt.Print("Enter pet name: ")
		fmt.Scanf("%v", &petName)
		fmt.Print("Enter pet age: ")
		fmt.Scanf("%v", &petAge)
		fmt.Print("Enter pet color: ")
		fmt.Scanf("%v", &petColor)

		pet := Pet{petName, petAge, petColor}
		petSlice = addPets(pet, petSlice)

	}
	listOfPets(petSlice)

}
