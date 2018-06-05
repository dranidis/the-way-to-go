package factorial

import (
	"testing"
)

// TEST OMIT
func TestFactorialZero(t *testing.T) {
	actual := Factorial(0)
	expected := 1
	if expected != actual {
		t.Errorf("Expected %d, got %d\n", expected, actual)
	}
}

// END OMIT
