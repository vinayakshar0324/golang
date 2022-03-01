package main

import "fmt"

func main() {
	var username string = "Vinayak"
	fmt.Println(username)
	fmt.Printf("Variables is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type: %T \n", smallVal)

	var smallFloat float64 = 255.4555556666677778
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type: %T \n", smallFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

}
