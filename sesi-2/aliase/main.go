package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte

	b = a // no error, karena byte memiliki tipe data yang sama dnegan uint8
	_ = b

	//membuat aliase sendiri

	type second = uint
	var hour second = 3600
	fmt.Printf("Hour type: %T\n", hour)
}
