package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	var employee1 = Employee{}

	employee1.person.name = "Vini"
	employee1.person.age = 22
	employee1.division = "Development Engineer"

	fmt.Printf("%+v", employee1)
}
