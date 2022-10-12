package riemann

import (
	"fmt"
	"math"
	"math/big"
)

//===================================================

type ExhaustiveSearchState struct {
	n int64
}

func NewExhaustiveSearchState(n int64) *ExhaustiveSearchState {
	ess := ExhaustiveSearchState{}
	ess.n = n
	return &ess
}

func (ess *ExhaustiveSearchState) Serialize() string {
	return fmt.Sprint(ess.n)
}

func (ess *ExhaustiveSearchState) Value() string {
	return fmt.Sprint(ess.n)
}

func (ess *ExhaustiveSearchState) GetNextBatch(batchSize int64) []SearchState {
	output := []SearchState{}
	startingVal := ess.n
	for i := int64(1); i <= batchSize; i++ {
		output = append(output, NewExhaustiveSearchState(startingVal+i))
	}
	return output
}

func (ess *ExhaustiveSearchState) ComputeRiemannDivisorSum() RiemannDivisorSum {
	i := ess.n
	ds, err := DivisorSum(i)
	if err != nil {
		panic("Divisor Sum cannot be found")
	}
	wv := WitnessValue(i, ds)
	return RiemannDivisorSum{N: i, WitnessValue: wv, DivisorSum: ds}
}

//===================================================

type SuperabundantSearchState struct {
	level        int
	indexInLevel int64
	value        []int
}

func NewSuperAbundantSearchState(level int, indexInLevel int64, value []int) *SuperabundantSearchState {
	sass := SuperabundantSearchState{}
	sass.level = level
	sass.indexInLevel = indexInLevel
	sass.value = value
	return &sass
}

func (sass *SuperabundantSearchState) Serialize() string {
	return fmt.Sprintf("%d, %d", sass.level, sass.indexInLevel)
}

func (sass *SuperabundantSearchState) Value() string {
	return fmt.Sprint(sass.value)
}

func (sass *SuperabundantSearchState) GetNextBatch(batchSize int64) []SearchState {
	output := []SearchState{}
	currentLevel := sass.level
	currentIndexInLevel := sass.indexInLevel + 1
	for len(output) <= int(batchSize) {
		partitions := PartitionsOfN(int(currentLevel))

		if currentIndexInLevel > int64(len(partitions)) {
			panic("index level is illegal")
		}

		partitionsToAdd := partitions[currentIndexInLevel:]
		for i, partition := range partitionsToAdd {
			output = append(output, SearchState(NewSuperAbundantSearchState(currentLevel, currentIndexInLevel+int64(i), partition)))
		}

		currentLevel += 1
		currentIndexInLevel = 0
	}
	return output[:batchSize]
}

func FindNFromPrimeFactors(PrimeFactors [][]int64) big.Int {
	n := *big.NewInt(1)
	for i := 0; i < len(PrimeFactors); i++ {
		expVal := new(big.Int).Exp(big.NewInt(PrimeFactors[i][0]), big.NewInt(PrimeFactors[i][1]), nil)
		n = *new(big.Int).Mul(expVal, &n)
	}
	return n
}

func (sass *SuperabundantSearchState) ComputeRiemannDivisorSum() RiemannDivisorSum {

	primeFactors := [][]int64{}
	primes := FirstNPrimes(len(sass.value))
	for i, x := range sass.value {
		primeFactors = append(primeFactors, []int64{int64(primes[i]), int64(x)})
	}

	divSum, err := PrimeFactorDivisorSum(primeFactors)
	if err != nil {
		panic("Divisor Sum cannot be found")
	}

	n := FindNFromPrimeFactors(primeFactors)
	nFloat, _ := new(big.Float).SetInt(&n).Float64()

	denom := nFloat * math.Log(math.Log(nFloat))

	num, _ := new(big.Float).SetInt(&divSum).Float64()

	wv := num / denom
	fmt.Println("Level:", sass.level, ", Index:", sass.indexInLevel)
	fmt.Println("N:", n, ", WV:", wv)

	return RiemannDivisorSum{N: n.Int64(), WitnessValue: wv, DivisorSum: divSum.Int64()} // TODO: hack!
}
