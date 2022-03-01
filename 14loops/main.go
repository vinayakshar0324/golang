package main

import (
	"fmt"
)

func main() {
	fmt.Println("loops in Go")

	days := []string{"Sunday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	for i := range days {
		fmt.Println(days[i])
	}

	// for _, day := range days {
	// 	fmt.Println("index is %v and value is %v\n", day)
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			break
		}

		fmt.Println("value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("jumping")

}
