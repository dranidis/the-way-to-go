package main

/* producer-consumer problem in Go */

import (
	"fmt"
)

var done = make(chan bool)
var msgs = make(chan int)

// PROD OMIT
func produce(n int) <-chan int {
	msgs := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			msgs <- i
		}
		done <- true
	}()
	return msgs
}

// END OMIT

// CONS OMIT
func consume(ch1, ch2 <-chan int) {
	for {
		select {
		case n1 := <-ch1:
			fmt.Printf("From ch1 %d\n", n1)
		case n2 := <-ch2:
			fmt.Printf("\tFrom ch2 %d\n", n2)
		}
	}
}

// END OMIT

func main() {
	ch1 := produce(5)
	ch2 := produce(5)
	go consume(ch1, ch2)
	<-done
	<-done
}
