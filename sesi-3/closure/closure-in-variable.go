package main

import "fmt"

func main() {
	var eventNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(eventNumbers(numbers...))
}
