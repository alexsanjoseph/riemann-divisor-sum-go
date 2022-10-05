package riemann

func PartitionsOfN(n int) [][]int {
	output := [][]int{{n}}
	arrayToAdd := make([]int, n)
	splitPoint := 0
	arrayToAdd[splitPoint] = n
	for i := n - 1; i > 0; i-- {
		newParts := PartitionsOfN(n - i)
		for _, part := range newParts {
			if part[0] <= i {
				output = append(output, append([]int{i}, part...))
			}
		}
	}
	return output
}

func cachedPartitionsOfN(n int, cache map[int][][]int) [][]int {
	possibleOutput, ok := cache[n]
	if ok {
		return possibleOutput
	}

	output := [][]int{{n}}
	arrayToAdd := make([]int, n)
	splitPoint := 0
	arrayToAdd[splitPoint] = n
	for i := n - 1; i > 0; i-- {
		newParts := cachedPartitionsOfN(n-i, cache)
		cache[n-i] = newParts
		for _, part := range newParts {
			if part[0] <= i {
				output = append(output, append([]int{i}, part...))
			}
		}
	}
	return output
}

func MemoizedPartitionsOfN(n int) [][]int {
	cache := make(map[int][][]int)
	return cachedPartitionsOfN(n, cache)
}
