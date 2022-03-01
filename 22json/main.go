package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"_"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Json in Go")
}

func EncodeJson() {
	lcoCourse := []course{
		{"ReactJs", 299, "lco.in", "abc12345", []string{"webdev", "js"}},
		{"Mern", 399, "lco.in", "ab12345", []string{"stack", "js"}},
		{"Python", 499, "lco.in", "bc12345", nil},
	}

	//package as json data

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename: "ReactJs",
		"Price": 299,
		"website": "LearnCodeOnline.in"
		"tags": ["web-dev", "js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

}
