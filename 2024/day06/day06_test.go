package day06

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`, "\n")

func TestWalkOff(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 41
	if got := walkOff(level); got != want {
		t.Errorf("walkOff(ex) = %d, want %d", got, want)
	}
}

func TestCountLoops(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 6
	if got := countLoops(level); got != want {
		t.Errorf("countLoops(ex) = %d, want %d", got, want)
	}

}
