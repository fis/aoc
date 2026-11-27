package day21

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

var ex = strings.TrimPrefix(`
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
`, "\n")

func TestCountReachable(t *testing.T) {
	steps := []int{1, 2, 3, 6, 10}
	want := []int{2, 4, 6, 16, 50}
	l := util.ParseFixedLevel([]byte(ex))
	if got := countReachable(l, steps); !cmp.Equal(got, want) {
		t.Errorf("countReachable(ex, %v) = %v, want %v", steps, got, want)
	}
}

func TestCountReachableDynamic(t *testing.T) {
	steps := []int{1, 2, 3, 6, 10, 50, 100}
	want := []int{2, 4, 6, 16, 50, 1594, 6536}
	l := util.ParseFixedLevel([]byte(ex))
	if got := countReachableDynamic(l, steps); !cmp.Equal(got, want) {
		t.Errorf("countReachableDynamic(ex, %v) = %v, want %v", steps, got, want)
	}
}
