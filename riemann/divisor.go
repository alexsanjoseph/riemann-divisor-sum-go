package riemann

import (
	"errors"
	"math"
	"math/big"
)

var ErrCheckFailed = errors.New("array not unique or prime")
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

func CheckIfPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func CheckUniqueness(n []int64) bool {
	allValsSet := make(map[int64]struct{})
	for _, x := range n {
		allValsSet[x] = struct{}{}
	}
	return len(allValsSet) == len(n)
}

func checkUniquePrimeFactors(PrimeFactors [][]int64) bool {
	baseArray := []int64{}
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

func PrimeFactorDivisorSum(PrimeFactors [][]int64) (big.Int, error) {

	if len(PrimeFactors) < 1 {
		return *big.NewInt(1), nil
	}

	checksPassed := checkUniquePrimeFactors(PrimeFactors)
	if !checksPassed {
		return *big.NewInt(0), ErrCheckFailed
	}

	// divisorSum := int64(1)
	divisorSum := big.NewInt(1)
	for _, x := range PrimeFactors {
		z0 := new(big.Int).Exp(big.NewInt(x[0]), big.NewInt(x[1]+1), nil)
		z1 := new(big.Int).Mul(divisorSum, new(big.Int).Sub(z0, big.NewInt(1)))
		z2 := new(big.Int).Sub(big.NewInt(x[0]), big.NewInt(1))
		z3 := new(big.Int).Div(z1, z2)

		divisorSum = z3
	}

	return *divisorSum, nil
}

func DivisorSumBig(n big.Int) (big.Int, error) {
	if n.Sign() <= 0 {
		return *big.NewInt(0), errors.New("value cannot be less than or equal to zero")
	}

	limit := new(big.Int).Sqrt(&n)
	sum := big.NewInt(0)

	for i := big.NewInt(1); i.Cmp(limit) < 1; i = new(big.Int).Add(i, big.NewInt(1)) {

		if new(big.Int).Mod(&n, i).Cmp(big.NewInt(0)) == 0 {
			sum = new(big.Int).Add(sum, i)
			if new(big.Int).Mul(i, i).Cmp(&n) != 0 {
				sum = new(big.Int).Add(sum, new(big.Int).Div(&n, i))
			}
		}

	}

	return *sum, nil
}
