package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{}
	employee1.name = "Vini"
	employee1.age = 22
	employee1.division = "development engineer"

	var employee2 = Employee{name: "Airell", age: 23, division: "marketing"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
