package main

import (
	"fmt"
	"errors"
)

// ERROR OMIT
type DivisionByZero float64

func (e DivisionByZero) Error() string {
	return "Cannot divide by zero: " + fmt.Sprint(float64(e))
}

// END OMIT

// DIV OMIT
func Division(i, j float64) (float64, error) {
	if j == 0.0 {
		return i, DivisionByZero(j)
	}
	return i / j, nil
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("square root of negative number is not defined")
	}
	return f, nil // TODO
}
// END OMIT

func main() {
	// CHECK OMIT
	for div := 2; div > -1; div-- {
		res, err := Division(3.0, float64(div)) // HL
		if err != nil { // HL
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%f\n", res)
		}
	}
	res, err := Sqrt(-1.0) // HL
	if err != nil { // HL
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%f\n", res)
	}
	// END OMIT
}
