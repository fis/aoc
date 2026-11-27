package day11

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`, "\n")

func TestTotalDistance(t *testing.T) {
	tests := []struct {
		factor int
		want   int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}
	galaxies := findGalaxies(util.Lines(ex))
	for _, test := range tests {
		expanded := expandSpace(galaxies, test.factor)
		if got := totalDistance(expanded); got != test.want {
			t.Errorf("totalDistance(ex * %d) = %d, want %d", test.factor, got, test.want)
		}

	}
}
