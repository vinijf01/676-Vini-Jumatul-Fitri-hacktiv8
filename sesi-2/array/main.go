package main

import (
	"fmt"
	"strings"
)

func main() {
	/*1. Array*/

	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var kata = [3]string{"Airell", "Nanda", "Vini"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", kata)

	/*2. Array(Modify Element Through Index)*/
	var fruits = [3]string{"apel", "anggur", "manggo"}
	fruits[0] = "apple"
	fruits[1] = "grape"
	fruits[2] = "mango"

	fmt.Println("%#v\n", fruits)

	/*3. Array(Loop through elements)*/
	//cara 1 : range loop
	for i, v := range fruits {
		fmt.Printf("Index : %d, Value : %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))
	//cara 2 :
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index : %d, Value : %s\n", i, fruits[i])
	}

	/*4. Array(Multidimensional Array)*/
	balances := [2][3]int{{5, 4, 6}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
