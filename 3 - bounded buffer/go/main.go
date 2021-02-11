package main

import (
	"fmt"
	"time"
)

func producer(val chan<- int) {

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		val <- i
	}

}

func consumer(val <-chan int) {

	time.Sleep(1 * time.Second)
	for {
		i := <-val
		fmt.Printf("[consumer]: %d\n", i)
		time.Sleep(50 * time.Millisecond)
	}

}

func main() {

	// TODO: make a bounded buffer
	buffer := make(chan int, 5) //makes the buffer size the same as defined in the c code.

	go consumer(buffer)
	go producer(buffer)

	select {}
}
