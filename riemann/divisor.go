package riemann

import (
	"errors"
	"math"
)

func DivisorSum(n int64) (int64, error) {
	if n <= 0 {
		return 0, errors.New("value cannot be less than or equal to zero")
	}

	limit := math.Sqrt(float64(n))
	sum := int64(0)

	for i := int64(1); float64(i) <= limit; i++ {
		if n%i == 0 {
			sum += i
			if float64(i) != limit {
				sum += n / i
			}
		}

	}

	return sum, nil
}
