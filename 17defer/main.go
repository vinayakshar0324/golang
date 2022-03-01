package main

import "fmt"

func main() {

	defer fmt.Println("helloworld")
	fmt.Println("defers in go")
	myDefer()

}

// LIFO

func myDefer() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
}
