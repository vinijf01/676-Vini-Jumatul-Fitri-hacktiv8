package main

import (
	"fmt"
	"strings"
)

func main() {
	//1
	studentList := print("Airell", "Vini", "Nanda", "Schannle", "Marco")

	fmt.Printf("%v\n", studentList)

	//2
	numberLists := []int{1, 2, 3, 4, 5, 6}
	result := sum(numberLists...)

	fmt.Println("Result : ", result)

	//3
	profile("Airel", "pasta", "ayam geprek", "ikan ros", "sate padang")
}

//1
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

//2
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

//3
func profile(name string, favfoods ...string) {
	margeFoods := strings.Join(favfoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", margeFoods)
}
