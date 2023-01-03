package main

import "fmt"

func main() {

	/*1. type assertion*/
	var v interface{}

	v = 20

	// v = v * 9 //errror
	value, ok := v.(int) //mengembalikan ke type data aslinya

	if ok == true {
		v = value * 9
	}

	fmt.Println(v)

	/*2. map & slice */
	rs := []interface{}{1, "airel", true, "vini", true}

	rm := map[string]interface{}{
		"Name":   "Airell",
		"Status": true,
		"Age":    23,
	}

	_, _ = rs, rm
}
