package main

import "fmt"

// Shape OMIT
type Shape interface {
	Area() float64
}

// PrintArea OMIT
func PrintArea(s Shape) {
	fmt.Printf("Shape's area is %.2f\n", s.Area())
}

func (s Square) PrintArea() { fmt.Printf("Area is %f\n", s.Area()) }

//START OMIT
// Square OMIT
type Square struct {
	x float64
}

// Area OMIT
func (s Square) Area() float64 { return s.x * s.x }

type Rectangle struct {
	Square // embedded
	y      float64
}

func (r Rectangle) Area() float64 { return r.x * r.y } // OVERRIDE

func (s Square) Perimeter() float64 { return s.x + s.Area()/s.x }

func main() {
	s := Square{5}
	PrintArea(s)
	// OMIT

	//fmt.Printf("Perimeter of %v:%T is %.2f\n", s, s, s.Perimeter())

	//r := Rectangle{Square{5}, 3}
	//fmt.Printf("Sides of rectangle are: %.2f and %.2f\n", r.x, r.y)
	//PrintArea(r)
	//fmt.Printf("Perimeter of %v:%T is %.2f!\n", r, r, r.Perimeter())
}

// END OMIT
