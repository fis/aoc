package day11

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

func TestFixedPoint(t *testing.T) {
	tests := []struct {
		name      string
		mapper    func(*util.Level) [][8]int
		tolerance int
		want      int
	}{
		{name: "near", mapper: nearMap, tolerance: 4, want: 37},
		{name: "far", mapper: farMap, tolerance: 5, want: 26},
	}
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
	for _, test := range tests {
		got := fixedPoint(level, test.mapper, test.tolerance)
		if got != test.want {
			t.Errorf("fixedPoint(%s, %d) = %d, want %d", test.name, test.tolerance, got, test.want)
		}
	}
}
