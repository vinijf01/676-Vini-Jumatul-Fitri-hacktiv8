package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/*1. Looping over string (byte-by-byte)*/
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	//kebalikan
	kota := []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(kota))

	/*2. Looping Over String(rune by rune)*/
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2)) // = 6, karena â == 2 byte

	// untuk menghitung karakter nya saja maka string harus diganti menjadi rune terlebih dahulu
	fmt.Printf("str1 byte length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 byte length => %d\n", utf8.RuneCountInString(str2))

	for i, s := range str2 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
}
