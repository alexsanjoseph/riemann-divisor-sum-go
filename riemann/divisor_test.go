package riemann_test

import (
	"math"
	"math/big"
	"reflect"

	"github.com/alexsanjoseph/riemann-divisor-sum-go/riemann"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Divisor Calculation", func() {

	Describe("Naive Divisor sum", func() {
		Context("should work on good cases", func() {
			It("Should give correct result", func() {
				output, err := riemann.DivisorSum(72)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(output).To(Equal(int64(195)))
			})

			It("Should give correct result", func() {
				output, err := riemann.DivisorSum(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(output).To(Equal(int64(1)))
			})
		})

		Context("should error out on bad cases", func() {
			It("zero value", func() {
				output, err := riemann.DivisorSum(0)
				Expect(output).To(BeZero())
				Expect(err).To(HaveOccurred())
			})

			It("negative value", func() {
				output, err := riemann.DivisorSum(-5)
				Expect(output).To(BeZero())
				Expect(err).To(HaveOccurred())
			})
		})

		It("should work on parametrized cases", func() {
			divisorSums := []int64{
				1, 3, 4, 7, 6, 12, 8, 15, 13, 18, 12, 28, 14, 24, 24, 31, 18, 39, 20, 42, 32,
				36, 24, 60, 31, 42, 40, 56, 30, 72, 32, 63, 48, 54, 48, 91, 38, 60, 56, 90, 42,
				96, 44, 84, 78, 72, 48, 124, 57, 93, 72,
			}

			for i, sum := range divisorSums {
				By("Passing each cases", func() {
					output, err := riemann.DivisorSum(int64(i) + 1)
					if err != nil {
						Fail("error should be nil")
					}
					Expect(output).To(Equal(sum))
				})
			}
		})
	})

	Describe("Prime factor divisor sum", func() {
		Context("should work on good cases", func() {
			It("For empty set", func() {
				inputArray := [][]int64{}
				output, err := riemann.PrimeFactorDivisorSum(inputArray)
				if err != nil {
					Fail("error should be nil")
				}
				Expect(output).To(Equal(*big.NewInt(1)))
			})

			It("For prime number", func() {
				inputArray := [][]int64{{7, 1}}
				output, err := riemann.PrimeFactorDivisorSum(inputArray)
				if err != nil {
					Fail("error should be nil")
				}
				Expect(output).To(Equal(*big.NewInt(8)))
			})

			It("For composite number", func() {
				inputArray := [][]int64{{2, 3}, {3, 2}}
				output, err := riemann.PrimeFactorDivisorSum(inputArray)
				if err != nil {
					Fail("error should be nil")
				}
				Expect(output).To(Equal(*big.NewInt(195)))
			})

			It("For failing case", func() {
				inputArray := [][]int64{{2, 3}}
				output, err := riemann.PrimeFactorDivisorSum(inputArray)
				if err != nil {
					Fail("error should be nil")
				}
				output2, err := riemann.DivisorSumBig(*big.NewInt(8))
				Expect(err).ShouldNot(HaveOccurred())
				Expect(output).To(Equal(output2))
			})

			It("It gives error for non prime base", func() {
				inputArray := [][]int64{{2, 3}, {4, 2}}
				_, err := riemann.PrimeFactorDivisorSum(inputArray)
				Expect(err).To(HaveOccurred())
			})

			It("It gives error for non unique base", func() {
				inputArray := [][]int64{{2, 3}, {2, 2}}
				_, err := riemann.PrimeFactorDivisorSum(inputArray)
				Expect(err).To(HaveOccurred())
			})

			It("For large number, it should work errors", func() {
				inputArray := [][]int64{{7, 10}, {29, 10}, {3, 14}}

				n := int64(1)
				for _, x := range inputArray {
					n *= int64(math.Pow(float64(x[0]), float64(x[1])))
				}

				_, err := riemann.PrimeFactorDivisorSum(inputArray)
				Expect(err).ShouldNot(HaveOccurred())
			})

			It("For large number, it should work", func() {
				inputArray := [][]int64{{17, 5}, {11, 5}, {13, 5}, {5, 15}}

				n := int64(1)
				for _, x := range inputArray {
					n *= int64(math.Pow(float64(x[0]), float64(x[1])))
				}

				_, err := riemann.PrimeFactorDivisorSum(inputArray)

				Expect(err).ShouldNot(HaveOccurred())

			})
		})

		Describe("Big Divisor sum", func() {
			Context("should work on good cases", func() {
				It("Should give correct result", func() {
					output, err := riemann.DivisorSumBig(*big.NewInt(72))
					Expect(err).ShouldNot(HaveOccurred())
					Expect(output).To(Equal(*big.NewInt(195)))
				})

				It("Should give correct result", func() {
					output, err := riemann.DivisorSumBig(*big.NewInt(1))
					Expect(err).ShouldNot(HaveOccurred())
					Expect(output).To(Equal(*big.NewInt(1)))
				})
			})

			Context("should error out on bad cases", func() {
				It("zero value", func() {
					output, err := riemann.DivisorSumBig(*big.NewInt(0))
					Expect(output).To(BeZero())
					Expect(err).To(HaveOccurred())
				})

				It("negative value", func() {
					output, err := riemann.DivisorSumBig(*big.NewInt(-5))
					Expect(output).To(BeZero())
					Expect(err).To(HaveOccurred())
				})
			})

			It("should work on parametrized cases", func() {
				divisorSums := []int64{
					1, 3, 4, 7, 6, 12, 8, 15, 13, 18, 12, 28, 14, 24, 24, 31, 18, 39, 20, 42, 32,
					36, 24, 60, 31, 42, 40, 56, 30, 72, 32, 63, 48, 54, 48, 91, 38, 60, 56, 90, 42,
					96, 44, 84, 78, 72, 48, 124, 57, 93, 72,
				}

				for i, sum := range divisorSums {
					By("Passing each cases", func() {
						output, err := riemann.DivisorSumBig(*big.NewInt(int64(i) + 1))
						if err != nil {
							Fail("error should be nil")
						}
						Expect(output).To(Equal(*big.NewInt(sum)))
					})
				}
			})
		})

		It("should work on autogenerated tests", func() {

			parameters := gopter.DefaultTestParameters()
			parameters.Rng.Seed(1235)
			parameters.MinSuccessfulTests = 10
			parameters.MaxSize = 30
			properties := gopter.NewProperties(parameters)

			properties.Property("Check Prime Factor Divisor Sum", prop.ForAll(
				func(a, b []int64) bool {

					smallerSliceLength := len(b)
					if len(a) <= len(b) {
						smallerSliceLength = len(a)
					}
					input := [][]int64{}
					n := *big.NewInt(1)

					for i := 0; i < smallerSliceLength; i++ {
						input = append(input, []int64{a[i], b[i]})
						expVal := new(big.Int).Exp(big.NewInt(a[i]), big.NewInt(b[i]), nil)
						n = *new(big.Int).Mul(expVal, &n)
					}

					resultA, err := riemann.PrimeFactorDivisorSum(input)
					if err != nil {
						return false
					}

					resultB, err := riemann.DivisorSumBig(n)
					if err != nil {
						return false
					}

					return resultB.Cmp(&resultA) == 0

				},
				gen.SliceOf(gen.Int64Range(1, 30).SuchThat(func(v interface{}) bool {
					return riemann.CheckIfPrime(v.(int64))
				}),
					reflect.TypeOf(int64(0))).SuchThat(func(v interface{}) bool {
					return riemann.CheckUniqueness(v.([]int64))
				}).WithLabel("a"),
				gen.SliceOf(gen.Int64Range(1, 5),
					reflect.TypeOf(int64(0))).WithLabel("b"),
			))
			Expect(properties.Run(gopter.ConsoleReporter(true))).To(BeTrue())

		})
	})
})
