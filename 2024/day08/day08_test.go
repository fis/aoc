package day08

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`, "\n")

func TestCountAntinodes(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	antennae := findAntennae(level)
	want := 14
	if got := countAntinodes(level, antennae); got != want {
		t.Errorf("countAntinodes(ex) = %d, want %d", got, want)
	}
}

func TestCountAllAntinodes(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	antennae := findAntennae(level)
	want := 34
	if got := countAllAntinodes(level, antennae); got != want {
		t.Errorf("countAllAntinodes(ex) = %d, want %d", got, want)
	}
}
