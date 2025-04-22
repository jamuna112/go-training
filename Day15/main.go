package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
Program for file handling
*/

func main() {
	var fileName string = "testfile.txt"
	fmt.Printf("reading a file:\n")
	readFile(fileName)

	fmt.Printf("\ncreating a file:\n")
	createFile(fileName)

	fmt.Printf("\nWrite a file:\n")
	writeFile(fileName, "Re changing the text")

	fmt.Printf("\nOpen and read a file:\n")
	openAndReadFile(fileName)

	fmt.Printf("\ndeleting a file:\n")
	deleteFile(fileName)

}
func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("file reading error ", err)
		return
	}
	fmt.Println("file content")
	fmt.Println(string(data))

}

func createFile(filename string) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("error creating file", err)
		return
	}
	defer file.Close()
	fmt.Print("File created ", filename)
}

func writeFile(filename string, text string) {
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		fmt.Println("error writing in file", err)
		return
	}
	fmt.Println("content return in the file ", filename)
}

func openAndReadFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("file reading error ", err)
		return
	}
	fmt.Printf("read %d bytes, content %s\n", count, string(data[:count]))
}

func deleteFile(fileName string) {
	error := os.Remove(fileName)

	if error != nil {
		fmt.Println("error deleting file", error)
		return
	}
	fmt.Println("Successfully deleted")
}
