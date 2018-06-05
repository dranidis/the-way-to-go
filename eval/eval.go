package main

import "fmt"
// M OMIT
func Eval(f func(int) int, x int) int {
	return f(x)
}

func square(x int) int {
	return x * x
}

func main() {
	fmt.Println(Eval(square, 2))

	increase := func(x int) int { return x+1}

	fmt.Println(Eval(increase, 0))

    fmt.Println(Eval(func(x int) int {return x * x * x}, 2))
}
// END OMIT