package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//anonymous struct tanpa pengisian field
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Vini", age: 22}
	employee1.division = "Development engineer"

	//anonymous struct dengan pengisian field
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "airell", age: 23},
		division: "marketing",
	}

	fmt.Println("Employee1: %+v\n", employee1)
	fmt.Println("Employee2: %+v\n", employee2)
}
