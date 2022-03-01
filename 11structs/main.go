package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in GO")
	//no inheritance in Golang; no super no parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
