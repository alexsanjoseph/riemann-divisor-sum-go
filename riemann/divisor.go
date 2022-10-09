package riemann

import (
	"errors"
	"math"
	"math/big"
)

var ErrIntegerOverflow = errors.New("integer overflow")

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
		if sum < 0 {
			return -1, ErrIntegerOverflow
		}

	}

	return sum, nil
}

func CheckIfPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}

func CheckUniqueness(n []int) bool {
	allValsSet := make(map[int]struct{})
	for _, x := range n {
		allValsSet[x] = struct{}{}
	}
	return len(allValsSet) == len(n)
}

func checkUniquePrimeFactors(PrimeFactors [][]int) bool {
	baseArray := []int{}
	for _, x := range PrimeFactors {
		baseArray = append(baseArray, x[0])
		isPrime := CheckIfPrime(x[0])
		if !isPrime {
			// fmt.Println("Base Array ", baseArray, " has values: ", x[0], " that is not prime")
			return false
		}
	}
	return CheckUniqueness(baseArray)
}

func PrimeFactorDivisorSum(PrimeFactors [][]int) (int64, error) {

	if len(PrimeFactors) < 1 {
		return 1, nil
	}

	checksPassed := checkUniquePrimeFactors(PrimeFactors)
	if !checksPassed {
		return -1, ErrIntegerOverflow
	}

	divisorSum := int64(1)
	for _, x := range PrimeFactors {
		divisorSum *= int64((math.Pow(float64(x[0]), float64(x[1]+1)) - 1) / (float64(x[0]) - 1))
		if divisorSum < 0 {
			return -1, ErrIntegerOverflow
		}
	}

	return divisorSum, nil
}
