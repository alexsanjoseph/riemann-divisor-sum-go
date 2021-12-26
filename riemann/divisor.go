package riemann

import "errors"

func DivisorSum(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("value cannot be less than or equal to zero")
	}
	sum := n
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum, nil
}
