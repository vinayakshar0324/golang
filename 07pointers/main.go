package main

import "fmt"

func main() {
	fmt.Println("Welcome to Class of Pointers")

	// var ptr *int

	// fmt.Println("the value of pointer is:", ptr)

	myNumber := 34

	var ptr = &myNumber

	fmt.Println("the value of pointer is:", ptr)
	fmt.Println("the value of actual pointer is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The New Value is:", myNumber)

}
