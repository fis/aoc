package day13

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`, "\n")

func TestSumReflections(t *testing.T) {
	ps := fn.Map(util.Chunks(ex), parsePattern)
	tests := []struct {
		name string
		find func(p pattern) (rx, ry int)
		want int
	}{
		{"findReflection", findReflection, 405},
		{"findSmudged", findSmudged, 400},
	}
	for _, test := range tests {
		if got := sumReflections(ps, test.find); got != test.want {
			t.Errorf("sumReflections(ex, %s) = %d, want %d", test.name, got, test.want)
		}

	}
}
