package errorhandling

import (
	"fmt"
	"os"
)

/*
	write a program that opens a file, reads it, and processes the contents. At the end file should always
	close even if something goes wrong (will use defer keyword)

	plan:
	1. open a file using os.Open() function
	2. if opening fails, return
	3. if successful, defer file.close()
	4. read and print the file
	5. now, close() is called so that we exit properly
*/

func ReadFile(filename string) {

	//opening a file
	file, err := os.Open(filename)

	if err != nil {
		fmt.Printf("error reading file %s %v\n", filename, err)
		return
	}

	//defer the clean up
	defer func() {
		fmt.Println("Closing the file..")
		file.Close()
	}()

	//processing
	fmt.Println("successfully read data from file")

	panic("something happen while file processing") //defer will still run
}
