package day24

import (
	"strings"
	"testing"
)

var ex = strings.Split(strings.TrimSpace(`
....#
#..#.
#..##
..#..
#....
`), "\n")

func TestFindRepeating(t *testing.T) {
	initial := parseState(ex)
	want := state(2129920)
	got := findRepeating(initial)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountBugs(t *testing.T) {
	initial := parseState(ex)
	got := countBugs(initial, 10)
	if got != 99 {
		t.Errorf("got %d, want 99", got)
	}
}
