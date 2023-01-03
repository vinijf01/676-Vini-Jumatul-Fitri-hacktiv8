package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee1 = Employee{name: "Vini", age: 22, division: "Development engineer"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

	fmt.Println(strings.Repeat("#", 30))

	employee2.name = "Amanda"

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)
}
