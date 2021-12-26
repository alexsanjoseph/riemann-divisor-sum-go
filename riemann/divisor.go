package riemann

import (
	"errors"
)

var DIVISOR_CACHE = make(map[int]int)

func DivisorSumNaive(n int) (int, error) {
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

func divisorSumCached(n int) (int, error) {

	result, ok := DIVISOR_CACHE[n]

	if ok {
		return result, nil
	}

	sum, error := DivisorSum(n)
	DIVISOR_CACHE[n] = sum
	return sum, error
}

func DivisorSum(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("value cannot be less than or equal to zero")
	}

	signSequence := []int{1, 1, -1, -1}

	sum := 0

	for i := 0; true; i++ {
		sign := signSequence[i%4]
		x := i/2 + 1

		if i%2 == 1 {
			x = -x
		}
		argDiff := int((3*x*x - x) / 2)

		if n-argDiff < 0 {
			break
		}

		if n-argDiff == 0 {
			sum += sign * n
			break
		}
		output, err := divisorSumCached(n - argDiff)
		if err != nil {
			panic("Error shouldn't come here!")
		}
		sum += sign * output
	}

	return sum, nil
}
