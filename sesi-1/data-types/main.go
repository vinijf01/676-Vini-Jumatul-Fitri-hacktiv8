package main

import "fmt"

func main() {
	/*1. Number (integer)*/

	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("type data first %T\n", first)
	fmt.Printf("type data second %T\n", second)

	/*1. Number (Float)*/

	var decimalnumber float32 = 3.63

	fmt.Printf("deciman Number: %f\n", decimalnumber)
	fmt.Printf("deciman Number: %0.3f\n", decimalnumber)

	/*2. Bool */

	var condition bool = false
	fmt.Printf("is it permitted? %t \n", condition)

	/*3. String*/

	message := `Selamat Datanng di "Hactiv8"
	.Salam kenal :).
	Mari belajar "Scalable Web Service With Go".`
	fmt.Print(message)
}
