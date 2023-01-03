package main

import "fmt"

func main() {
	/*1. Condition (temporary variable)*/

	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("Kamu belum noleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	/*2. Switch*/
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	/*3. Switch with relational operators*/

	var nilai = 2
	switch {
	case nilai == 8:
		fmt.Println("perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}

	/*4. Switch (fallthrough keyword)*/
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("Perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
		fallthrough
	case score2 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don't have a good score yet")
		}
	}

	/*5. Nested Conditions*/

	var score3 = 10

	if score > 7 {
		switch score3 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score3 == 5 {
			fmt.Println("not bad")
		} else if score3 == 3 {
			fmt.Println("Keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
