package riemann

type RiemannDivisorSum struct {
	N            int64
	DivisorSum   int64
	WitnessValue float64
}

type SummaryStats struct {
	LargestComputedN    RiemannDivisorSum
	LargestWitnessValue RiemannDivisorSum
}

type DivisorDb interface {
	Load() []RiemannDivisorSum
	Upsert([]RiemannDivisorSum)
	Summarize() SummaryStats
}
