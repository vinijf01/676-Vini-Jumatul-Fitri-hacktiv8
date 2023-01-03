package main

import "fmt"

func main() {
	/*1. Constant*/

	const full_name string = "Airel Jordan"

	fmt.Printf("Hello %s", full_name)

	/*2. Operators (Arithmetic Operators)*/

	var value = 2 + 2*3
	fmt.Println(value)

	/*2. Operators (Relational Operators)*/

	var firstConditional bool = 2 < 3
	var secondConditional bool = "joey" == "Joey"
	var thridConditional bool = 10 != 2.3
	var fourtgConditional bool = 11 <= 11

	fmt.Println("first condition:", firstConditional)
	fmt.Println("second condition:", secondConditional)
	fmt.Println("third condtiion:", thridConditional)
	fmt.Println("fourth condition:", fourtgConditional)

	/*2. Operator (logical Operators)*/

	var right = true
	var wrong = false

	var wrongAndRigth = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRigth)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)

}
