package riemann

import "fmt"

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

func (ess *ExhaustiveSearchState) Value() int64 {
	return ess.n
}

func (ess *ExhaustiveSearchState) GetNextBatch(batchSize int64) []SearchState {
	output := []SearchState{}
	startingVal := ess.Value()
	for i := int64(1); i <= batchSize; i++ {
		output = append(output, SearchState(&ExhaustiveSearchState{startingVal + i}))
	}
	return output
}

func (ess *ExhaustiveSearchState) GetLatestSearchState() SearchState {

	ss := SearchState(&ExhaustiveSearchState{
		n: 10000,
	})

	return ss
}
