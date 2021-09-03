package util

import (
	"math"
	"math/big"
	"strconv"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TriangularNumber implements an efficient computation of the object count
// inside a equilateral triangle (which conveniently matches the prefix sum of
// natural numbers) by calculating the binomial coefficient of (n+1, 2).
// Internally the fastest possible solution for the provided `n` is chosen
// (either performing calculations over an uint64 or math/big.Int). The result
// is formatted to it's base(10) equivalent string and returned.
//
// https://oeis.org/A000217
func TriangularNumber(n uint64) string {
	// Until Sqrt(MaxUint64)-1 (which is equal to MaxUint32) we can safely use
	// uint64 during calculations without any overflows, therefore reduce
	// allocations and improve performance for smaller `n`s
	if n <= math.MaxUint32 {
		return triangularFast(n)
	}

	// from n > MaxUint32 we need to use Go's big.Int to avoid integer
	// overflows
	return triangularBig(n)
}

func triangularFast(n uint64) string {
	// (n * (n + 1)) / 2
	n *= (n + 1)
	n /= 2
	return strconv.FormatUint(n, 10)
}

func triangularBig(n uint64) string {
	// (n * (n + 1)) / 2
	nBig := new(big.Int).SetUint64(n)
	nBigPlus1 := new(big.Int).Add(nBig, big.NewInt(1))
	nBig.Mul(nBig, nBigPlus1)
	nBig.Div(nBig, big.NewInt(2))
	return nBig.Text(10)
}
