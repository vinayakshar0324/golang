package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev/3000/learn?coursename=reacts&paymentid=ghhsdjaflk"

func main() {
	fmt.Println("Handling Urls in Go")
	fmt.Println(myUrl)

	//parsing

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("the type query params are: %T\n", qParams)

	fmt.Println(qParams["coursename"])

	for _, val := range qParams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/learn",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
