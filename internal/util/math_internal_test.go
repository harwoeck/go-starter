package util

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompareFastAndBigImplementationBelowBreakPoint(t *testing.T) {
	// this test compares the results for the internal implementations of the
	// fast path and the slow path. This gives us more confidence in the big
	// implementation, which hasn't really any verified test vectors.

	compare := func(n uint64) {
		fmt.Println(n)
		var fast, big string

		require.NotPanics(t, func() {
			fast = triangularFast(n)
		}, "triangularFast paniced, but not expected for n=%d", n)

		require.NotPanics(t, func() {
			big = triangularBig(n)
		}, "triangularBig paniced, but not expected for n=%d", n)

		assert.Equal(t, fast, big, "fast and big implementation must produce same result until break-point at MaxUint32")
	}

	// run every case in 0-53 (as the fast implementation is verified against the
	// OEIS test vectors in this range).
	for n := uint64(0); n < 53; n++ {
		compare(n)
	}

	/* #nosec G404 */ // it's ok here
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := uint64(53); n < math.MaxUint32; n += uint64(r.Int63()) % 10000000 { // randomly increase n and test
		compare(n)
	}

	// as last example compare the biggest n possible for fast
	compare(math.MaxUint32)
}
