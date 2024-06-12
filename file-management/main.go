package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file - Rohan Bhatia"

	file, err := os.Create("./myfile.txt")

	checkNilError(err)

	length, _ := io.WriteString(file, content)

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Printf("Text data inside file is: %v\n", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
