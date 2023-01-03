package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Vini", age: 22}, division: "development engineer"},
		{person: Person{name: "Amanda", age: 23}, division: "marketing"},
		{person: Person{name: "Airell", age: 22}, division: "finance"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
