package day20

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`, "\n")

func TestAnalyze(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want1, want2 := 1, 285
	if got1, got2 := analyze(level, 50); got1 != want1 || got2 != want2 {
		t.Errorf("analyze(ex, 50) = (%d,%d), want (%d,%d)", got1, got2, want1, want2)
	}
}
