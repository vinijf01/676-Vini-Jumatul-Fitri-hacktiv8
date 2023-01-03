package main

import (
	"fmt"
	"strings"
)

func main() {
	/*0. Map*/
	var person map[string]string //deklarasi

	person = map[string]string{ //inisialisasi
		"name": "Vini",
		"age":  "21",
	}

	person["address"] = "Jalan Sudirman"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	fmt.Println(strings.Repeat("#", 25))

	/*1. Map(Looping with map)*/

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
	fmt.Println(strings.Repeat("#", 25))

	/*2. Map (Deleteing value) & map (detecting a value)*/
	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	fmt.Println("Before deleting: ", person)

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value does'nt exist")
	}

	fmt.Println("after deleting: ", person)

	/*3. Map(combining slice with map)*/
	people := []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Vini", "age": "21"},
	}

	for i, value := range people {
		fmt.Printf("index: %d, name: %s, age: %s\n", i, value["name"], value["age"])
	}
}
