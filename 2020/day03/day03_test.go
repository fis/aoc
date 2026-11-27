package day03

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const example = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestCountTrees(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(example, "\n"), '.')

	tests := []struct {
		slope util.P
		want  int
	}{
		{slope: util.P{1, 1}, want: 2},
		{slope: util.P{3, 1}, want: 7},
		{slope: util.P{5, 1}, want: 3},
		{slope: util.P{7, 1}, want: 4},
		{slope: util.P{1, 2}, want: 2},
	}
	for _, test := range tests {
		got := countTrees(level, test.slope)
		if got != test.want {
			t.Errorf("countTrees(..., %v) = %d, want %d", test.slope, got, test.want)
		}
	}
}
