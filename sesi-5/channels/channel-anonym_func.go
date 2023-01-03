package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}
}

func print(c chan string) {
	fmt.Println(<-c)
}
