package main

import "fmt"

func main() {
	msg := "later"
	defer fmt.Println(msg) // HL

	msg = "See you"

	fmt.Println(msg)
}