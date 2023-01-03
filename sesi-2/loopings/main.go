package main

import "fmt"

func main() {
	/*1. Looping (First way of loopong)*/

	for i := 0; i < 3; i++ {
		fmt.Println("Angka 1: ", i)
	}

	/*2. Loopings (Second way of looping)*/
	var j = 0

	for j < 3 {
		fmt.Println("Angka 2: ", j)
		j++
	}

	/*3. Loopings (third way of looping) use key word _break_ */
	var k = 0

	for {
		fmt.Println("Angka 3: ", k)

		k++
		if k == 3 {
			break
		}
	}

	/*4. loopings break and continue keyword*/
	for l := 1; l <= 10; l++ {
		if l%2 == 1 {
			continue
		}

		if l > 8 {
			break
		}

		fmt.Println("Angka 4:", l)
	}

	/*5. Loopings  Nested Looping*/
	for m := 0; m < 5; m++ {
		for n := m; n < 5; n++ {
			fmt.Print(n, " ")
		}

		fmt.Println()
	}

	/*6. Loopings (Label)*/
outerloop:
	for a := 0; a < 3; a++ {
		fmt.Println("Perulangan ke - ", a+1)
		for b := 0; b < 3; b++ {
			if a == 2 {
				break outerloop
			}
			fmt.Print(b, " ")
		}
		fmt.Print("\n")
	}

}
