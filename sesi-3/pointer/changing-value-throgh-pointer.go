package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (Value): ", firstPerson)
	fmt.Println("firstPerson (memory address): ", &firstPerson)
	fmt.Println("secondPerson (Value): ", *secondPerson)
	fmt.Println("secondPerson (memory address): ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (Value): ", firstPerson)
	fmt.Println("firstPerson (memory address): ", &firstPerson)
	fmt.Println("secondPerson (Value): ", *secondPerson)
	fmt.Println("secondPerson (memory address): ", secondPerson)
}
