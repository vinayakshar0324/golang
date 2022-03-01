package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web request in GO")
	PerformGetRequest()
	PerformJsonPostRequest()
	PerformFormRequest()
}

func PerformGetRequest() {
	const myUrl = "https://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)

	fmt.Println(string(content))

}

func PerformJsonPostRequest() {
	const myUrl = "https://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName": "Golang",
			"price":"400",
			"platform":"lco.in"

		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "vinayak")
	data.Add("lastName", "sharma")
	data.Add("email", "vinayak@dev.com")
	data.Add("role", "admin")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
