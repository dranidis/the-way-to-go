package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CODE OMIT
func foo1() {
	for {
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("foo1")
	}
}

func main() {

	// GO OMIT
	go foo1() // HL

	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("main")
	}
}

// END OMIT
