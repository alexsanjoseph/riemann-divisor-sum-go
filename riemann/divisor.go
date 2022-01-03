package riemann

import (
	"errors"
	"math"
)

func DivisorSum(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("value cannot be less than or equal to zero")
	}

	limit := math.Sqrt(float64(n))
	sum := 0

	for i := 1; float64(i) <= limit; i++ {
		if n%i == 0 {
			sum += i
			if float64(i) != limit {
				sum += n / i
			}
		}

	}

	return sum, nil
}
