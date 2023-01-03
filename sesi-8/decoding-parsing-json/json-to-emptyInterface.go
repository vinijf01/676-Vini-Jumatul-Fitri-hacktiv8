package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
	{
		"full_name": "Airell Jordan",
		"email":"airel@gmail.com",
		"age":23
	}
	`
	var temp interface{}

	var err = json.Unmarshal([]byte(jsonString), &temp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var result = temp.(map[string]interface{})

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("emai:", result["email"])
	fmt.Println("age:", result["age"])
}
