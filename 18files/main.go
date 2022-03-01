package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")

	content := "This needs to go in file"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("/Users/bimerboy/Desktop/golang/18files/mylcogofile.txt")
}

func readFile(filename, string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text Data inside the file is \n", string(databyte))
}
