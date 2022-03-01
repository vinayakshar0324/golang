package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of FruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 934
	highScore[2] = 434
	highScore[3] = 334
	//highScore[4] = 564

	highScore = append(highScore, 578, 900, 0324)

	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove slice based on index

	var courses = []string{"React", "javascript", "Ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
