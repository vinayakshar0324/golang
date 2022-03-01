package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["GO"] = "Golang"
	language["PY"] = "Python"

	fmt.Println("List of all languages:", language)
	fmt.Println("short from", language["JS"])

	delete(language, "PY")

	fmt.Println(language)

	//loops

	for key, value := range language {
		fmt.Printf("for Key %v, value is %v\n", key, value)
	}

}
