package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	myChannel := make(chan int)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	go func(){ch chan int, wg *sync.WaitGroup}(wg.Done()))


}
