package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Vini", "Airell"}
	printmsg := greet("Heii", names)

	fmt.Println(printmsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}
