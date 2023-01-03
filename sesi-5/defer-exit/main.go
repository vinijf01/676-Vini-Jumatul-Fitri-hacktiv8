package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println("defer function starts to execute")
	// fmt.Println("Hai Everyone")
	// fmt.Println("Welcome back to Go learning ccenter")
	callDeferFunc()
	fmt.Println("hai Everyone!!")

	defer fmt.Println("invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
