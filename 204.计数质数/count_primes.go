package _04_计数质数

import (
	"math"
)

func countPrimes(n int) int {
	var c int
	for i := 2; i < n; i++ {
		if isPrimes(i) {
			c++
		}

	}

	return c
}

func isPrimes(n int) bool {
	max := int(math.Sqrt(float64(n)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
