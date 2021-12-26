package riemann

import (
	"fmt"
	"testing"
)

func TestDivisorSum(t *testing.T) {
	sum := DivisorSum(192)
	fmt.Println(sum)
	if sum != 72 {
		t.Fatal("Divisor sum doesn't match!")
	}
}
