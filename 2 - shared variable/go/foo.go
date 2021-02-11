// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"runtime"
)

func number_server(CH_addNr <-chan int, CH_subNr <-chan int, CH_read chan<- int) {
	var number = 0
	// This for-select pattern is one you will become familiar with...
	for {
		select {
		case addingNr := <-CH_addNr:
			number = number + addingNr
		case subtractingNr := <-CH_subNr:
			number = number - subtractingNr
		case CH_read <- number:
		}
	}
}

func incrementer(add chan<- int, CH_finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add <- 1
	}
	CH_finished <- true
}

func decrementer(sub chan<- int, CH_finished chan<- bool) {
	for j := 0; j < 1000000+1; j++ {
		sub <- 1
	}
	CH_finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	CH_addNr := make(chan int)
	CH_subNr := make(chan int)
	CH_read := make(chan int)
	CH_finished := make(chan bool)

	go incrementer(CH_addNr, CH_finished)
	go decrementer(CH_subNr, CH_finished)
	go number_server(CH_addNr, CH_subNr, CH_read)

	for i := 0; i < 2; i++ {
		<-CH_finished
	}

	fmt.Println("The magic number is:", <-CH_read)
}
