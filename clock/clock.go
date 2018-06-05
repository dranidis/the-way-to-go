package main

import (
	"fmt"
	"time"
)

// CODE OMIT

func clock(tick chan int) {
	for {
		// WRITE OMIT
		tick <- 0 // Send blocks till a receiver receives
		// END OMIT
		fmt.Println("TICK")
	}
}

func main() {
	tick := make(chan int)

	// GO OMIT
	go clock(tick) // HL
	// END OMIT

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		// READ OMIT
		<-tick // Receive from channel
		// END OMIT
	}
}

// ENDCODE OMIT
