package main

import (
	"fmt"
	"strconv"
)

func main() {

	number, err := strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
}
