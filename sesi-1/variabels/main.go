package main

import "fmt"

func main() {
	/*1. Variabel With Data Type*/

	var name string
	var age int = 21

	name = "Vini"
	//re-assign value
	age = 22

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)

	/*2. Variable Without Data Type
	disebut juga dengan type inference
	- %T : untuk jenis type data
	- %v : untuk value dari variable
	- %s : untuk tipe data string
	- %d : untuk tipe data integer
	- %t : untuk type data bool
	*/

	var name1 = "Vini1"
	var age1 = 21

	fmt.Printf("%T, %v \n", name1, age1)
	fmt.Printf("%v, %T \n", name1, age1)

	//SHORT DECLRATION
	name2 := "Airel"
	age2 := 23

	fmt.Printf("%v, %v \n", name2, age2)

	/*3. Multiple Variable Declration*/

	var student1, student2, student3 = "satu", 2, "tiga"

	fmt.Println(student1, student2, student3)

	/*4. Underscore variable*/

	var firstVariable string

	var name3, age3, address3 = "Airell", 23, "Jalan sudirman"

	_, _, _, _ = firstVariable, name3, age3, address3

}
