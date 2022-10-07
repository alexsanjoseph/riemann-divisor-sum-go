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

func PartitionsOfKun(n int) [][]int {
	arrayToAdd := make([]int, n)
	splitPoint := 0
	output := [][]int{}
	arrayToAdd[splitPoint] = n

	for {
		newArray := []int{}
		for _, piece := range arrayToAdd {
			if piece != 0 {
				newArray = append(newArray, piece)
			}
		}
		output = append(output, newArray)

		rightOfNonOne := 0

		for splitPoint >= 0 && arrayToAdd[splitPoint] == 1 {
			rightOfNonOne += 1
			splitPoint -= 1
		}

		if splitPoint < 0 {
			break
		}

		arrayToAdd[splitPoint] -= 1
		amountToSplit := rightOfNonOne + 1

		for amountToSplit > arrayToAdd[splitPoint] {
			arrayToAdd[splitPoint+1] = arrayToAdd[splitPoint]
			amountToSplit -= arrayToAdd[splitPoint]
			splitPoint += 1
		}
		arrayToAdd[splitPoint+1] = amountToSplit
		splitPoint += 1
	}
	return output
}
