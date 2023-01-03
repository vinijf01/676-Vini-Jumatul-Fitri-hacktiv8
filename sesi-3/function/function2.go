package main

import "fmt"

func main() {
	greet("vini", "jalan sudirman")
}

//sebuah fungsi baru bernama fungsi greet dg 2 param yaitu name dan address yang memiliki tipe data sama

func greet(name, address string) {
	fmt.Println("Hello there! My name is ", name)
	fmt.Println("I live in", address)
}
