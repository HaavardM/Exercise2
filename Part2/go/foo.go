// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		case c := <-control:
			switch c {
			case GetNumber:
				number <- i
			case Exit:
				break
			}
		case n := <-add_number:
			i += n
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
		}
	}
}

func incrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- 1
	}
	//TODO: signal that the goroutine is finished
	close(finished)
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		add_number <- -1
	}
	//TODO: signal that the goroutine is finished
	close(finished)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	addNumber := make(chan int)
	control := make(chan int)
	number := make(chan int)
	incFinished := make(chan bool)
	decrFinished := make(chan bool)

	// TODO: Spawn the required goroutines
	go number_server(addNumber, control, number)
	go incrementing(addNumber, incFinished)
	go decrementing(addNumber, decrFinished)

	// TODO: block on finished from both "worker" goroutines
	<-incFinished
	<-decrFinished

	control <- GetNumber
	Println("The magic number is:", <-number)
	control <- Exit
}
