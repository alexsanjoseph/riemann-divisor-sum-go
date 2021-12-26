package riemann_test

import (
	"fmt"
	"testing"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
)

func TestDivisorSum(t *testing.T) {
	sum := riemann.DivisorSum(192)
	fmt.Println(sum)
	if sum != 72 {
		t.Fatal("Divisor sum doesn't match!")
	}
}
