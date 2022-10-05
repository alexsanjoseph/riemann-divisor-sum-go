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
