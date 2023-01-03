package main

import "fmt"

func main() {
	greet("vini", 21)
}

//sebuah fungsi baru bernama fungsi greet dg 2 param yaitu name dan age

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}
