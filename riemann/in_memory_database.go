package riemann

type InMemoryDivisorDb struct {
	Data map[int64]RiemannDivisorSum
}

func (imdb InMemoryDivisorDb) Load() []RiemannDivisorSum {
	output := []RiemannDivisorSum{}
	for _, value := range imdb.Data {
		output = append(output, value)
	}
	return output
}

func (imdb InMemoryDivisorDb) Upsert(rds []RiemannDivisorSum) {
	for _, value := range rds {
		imdb.Data[value.N] = value
	}
}

func (imdb InMemoryDivisorDb) Summarize() SummaryStats {
	if len(imdb.Data) == 0 {
		return SummaryStats{
			LargestWitnessValue: RiemannDivisorSum{},
			LargestComputedN:    RiemannDivisorSum{},
		}
	}
	largest_computed_n := RiemannDivisorSum{N: -1}
	largest_witness_value := RiemannDivisorSum{WitnessValue: -1}

	for _, rds := range imdb.Data {
		if rds.N > largest_computed_n.N {
			largest_computed_n = rds
		}
		if rds.WitnessValue > largest_witness_value.WitnessValue {
			largest_witness_value = rds
		}
	}
	return SummaryStats{
		LargestWitnessValue: largest_witness_value,
		LargestComputedN:    largest_computed_n,
	}

}
