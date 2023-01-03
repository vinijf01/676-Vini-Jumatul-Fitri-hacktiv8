package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "manggo", "durian", "rambutan"}
	var wg sync.WaitGroup
	for index, fruit := range fruits {
		wg.Add(1) // untuk menambah counter dr waitgroup (counter berguna untuk memberitahu jumlah go routine yang harus ditunggu)
		go printFruit(index, fruit, &wg)
	}

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
