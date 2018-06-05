package main

import "fmt"

func Add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	add5 := Add(5)

	fmt.Println(add5(4))
    fmt.Println(Add(5)(4))
}