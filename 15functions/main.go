package main

import "fmt"

//main acts as an entry point of your program
func main() {
	fmt.Println("functions in Golang")
	greeter()
	greeterTwo()

	result := adder(3, 4)

	fmt.Println("The result is:", result)

	proRes := proAdder(2, 5, 8, 7)
	fmt.Println("Pro Result", proRes)

	//you are not allowed to write functions inside the function
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total

}

func greeterTwo() {
	fmt.Println("greeter two")
}

func greeter() {
	fmt.Println("hello from golang")
}
