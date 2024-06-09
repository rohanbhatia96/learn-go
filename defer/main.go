package main

import "fmt"

func main() {
	defer fmt.Println("I'm Rohan")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
	fmt.Println("Last statement")
}

func myDefer() {
	fmt.Println("myDefer function has been called")
	defer fmt.Println("")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
