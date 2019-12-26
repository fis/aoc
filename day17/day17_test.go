package day17

import (
	"strings"
	"testing"

	"github.com/fis/aoc2019-go/util"
)

const ex = `
..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..
`

func TestCrosses(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
	got := crosses(level)
	if got != 76 {
		t.Errorf("got %d, want 76", got)
	}
}
