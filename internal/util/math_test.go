package util_test

import (
	"math"
	"testing"

	"allaboutapps.dev/aw/go-starter/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinAndMapInt(t *testing.T) {
	max := math.MaxInt32
	min := math.MinInt32
	assert.Equal(t, max, util.MaxInt(max, min))
	assert.Equal(t, max, util.MaxInt(min, max))
	assert.Equal(t, min, util.MinInt(max, min))
	assert.Equal(t, min, util.MinInt(min, max))
	assert.Equal(t, min, util.MaxInt(min, min))
	assert.Equal(t, max, util.MinInt(max, max))
}

func TestTriangularNumber(t *testing.T) {
	cases := map[uint64]string{
		// Official OEIS test cases from https://oeis.org/A000217/list
		0:  "0",
		1:  "1",
		2:  "3",
		3:  "6",
		4:  "10",
		5:  "15",
		6:  "21",
		7:  "28",
		8:  "36",
		9:  "45",
		10: "55",
		11: "66",
		12: "78",
		13: "91",
		14: "105",
		15: "120",
		16: "136",
		17: "153",
		18: "171",
		19: "190",
		20: "210",
		21: "231",
		22: "253",
		23: "276",
		24: "300",
		25: "325",
		26: "351",
		27: "378",
		28: "406",
		29: "435",
		30: "465",
		31: "496",
		32: "528",
		33: "561",
		34: "595",
		35: "630",
		36: "666",
		37: "703",
		38: "741",
		39: "780",
		40: "820",
		41: "861",
		42: "903",
		43: "946",
		44: "990",
		45: "1035",
		46: "1081",
		47: "1128",
		48: "1176",
		49: "1225",
		50: "1275",
		51: "1326",
		52: "1378",
		53: "1431",
		// additional checks to verify uint64 overflow scenarios (for big.Int)
		math.MaxUint32:     "9223372034707292160", // "last" n for fast route
		math.MaxUint32 + 1: "9223372039002259456", // "first" n for big.Int route
		math.MaxInt64:      "42535295865117307928310139910543638528",
		math.MaxUint64:     "170141183460469231722463931679029329920",
	}

	for n, expect := range cases {
		var got string

		require.NotPanics(t, func() {
			got = util.TriangularNumber(n)
		}, "test case paniced, but not expected for n=%d", n)

		assert.Equal(t, expect, got)
	}
}
