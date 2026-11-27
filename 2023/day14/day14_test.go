package day14

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

var ex = strings.TrimPrefix(`
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`, "\n")

func TestTotalLoad(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	slideNorth(l)
	want := 136
	if got := totalLoad(l); got != want {
		t.Errorf("totalLoad(slideNorth(ex)) = %d, want %d", got, want)
	}
}

func TestSlideNorth(t *testing.T) {
	want := util.ParseFixedLevel([]byte(strings.TrimPrefix(`
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`, "\n")))
	got := util.ParseFixedLevel([]byte(ex))
	slideNorth(got)
	if diff := cmp.Diff(printLevel(want), printLevel(got)); diff != "" {
		t.Errorf("slideNorth(ex) mismatch (-want +got):\n%s", diff)
	}
}

func printLevel(l *util.FixedLevel) string {
	var buf strings.Builder
	for y := range l.H {
		fmt.Fprintf(&buf, "%s\n", l.Row(y))
	}
	return buf.String()
}
