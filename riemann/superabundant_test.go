package riemann_test

import (
	"fmt"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func GeneratePartitionTests(i int, expectedPartitions [][]int) {
	Describe("Finds partition for an integer", func() {
		It("", func() {
			output := riemann.PartitionsOfN(i + 1)
			Expect(output).To(Equal(expectedPartitions))
		})
	})
}

func GeneratePartitionCountTests(i int, expectedPartitionCount int) {
	Describe("Finds partition count for an integer", func() {
		It("", func() {
			output := riemann.PartitionsOfN(i + 1)
			Expect(len(output)).To(Equal(expectedPartitionCount))
		})
	})
}

var _ = Describe("For partition counts", func() {

	expectedPartitions := [][][]int{
		{{1}},
		{{2}, {1, 1}},
		{{3}, {2, 1}, {1, 1, 1}},
		{{4}, {3, 1}, {2, 2}, {2, 1, 1}, {1, 1, 1, 1}},
		{{5}, {4, 1}, {3, 2}, {3, 1, 1}, {2, 2, 1}, {2, 1, 1, 1}, {1, 1, 1, 1, 1}},
		{{6}, {5, 1}, {4, 2}, {4, 1, 1}, {3, 3}, {3, 2, 1}, {3, 1, 1, 1}, {2, 2, 2}, {2, 2, 1, 1}, {2, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1}},
		{{7}, {6, 1}, {5, 2}, {5, 1, 1}, {4, 3}, {4, 2, 1}, {4, 1, 1, 1}, {3, 3, 1}, {3, 2, 2}, {3, 2, 1, 1}, {3, 1, 1, 1, 1},
			{2, 2, 2, 1}, {2, 2, 1, 1, 1}, {2, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1}},
		{{8}, {7, 1}, {6, 2}, {6, 1, 1}, {5, 3}, {5, 2, 1}, {5, 1, 1, 1}, {4, 4}, {4, 3, 1}, {4, 2, 2}, {4, 2, 1, 1}, {4, 1, 1, 1, 1},
			{3, 3, 2}, {3, 3, 1, 1}, {3, 2, 2, 1}, {3, 2, 1, 1, 1}, {3, 1, 1, 1, 1, 1}, {2, 2, 2, 2}, {2, 2, 2, 1, 1}, {2, 2, 1, 1, 1, 1},
			{2, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	expectedPartitionsCounts := []int{
		1, 2, 3, 5, 7,
		11, 15, 22, 30, 42,
		56, 77, 101, 135, 176,
		231, 297, 385, 490, 627,
		792, 1002, 1255, 1575, 1958,
		2436, 3010, 3718, 4565,
	}

	for i := 0; i < len(expectedPartitions); i++ {
		GeneratePartitionTests(i, expectedPartitions[i])
	}

	for i := 0; i < 20; i++ {
		fmt.Println(expectedPartitionsCounts[i])
		GeneratePartitionCountTests(i, expectedPartitionsCounts[i])
	}

})
