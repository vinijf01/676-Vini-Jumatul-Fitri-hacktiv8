package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine start sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data int the channel")
	}(c1)

	fmt.Println("main goroutine sleeps fo 2 second")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine start receiving data")
	d := <-c1
	fmt.Println("main goroutine recived d ata:", d)

	close(c1)
	time.Sleep(time.Second)
}
