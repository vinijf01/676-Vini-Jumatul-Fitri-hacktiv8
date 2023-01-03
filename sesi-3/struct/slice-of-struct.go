package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var people = []Person{
		{name: "Vini", age: 21},
		{name: "Airell", age: 23},
		{name: "Mailo", age: 23},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
