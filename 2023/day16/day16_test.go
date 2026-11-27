package day16

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`, "\n")

func TestCountEnergized(t *testing.T) {
	tests := []struct {
		x, y, dx, dy, di int
		want             int
	}{
		{0, 0, 1, 0, 0, 46},
		{3, 0, 0, 1, 1, 51},
	}
	l := util.ParseFixedLevel([]byte(ex))
	for _, test := range tests {
		if got := countEnergized(l, test.x, test.y, test.dx, test.dy, test.di); got != test.want {
			t.Errorf("countEnergized(ex, %d, %d, ...) = %d, want %d", test.x, test.y, got, test.want)
		}
	}
}
