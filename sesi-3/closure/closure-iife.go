//immediately-invoked function expression
package main

import "fmt"

func main() {
	var isPolindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}
		return temp == str
	}("katak")

	fmt.Println(isPolindrome)
}
