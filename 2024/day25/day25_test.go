package day25

import (
	"testing"
)

var ex = []string{
	"#####",
	".####",
	".####",
	".####",
	".#.#.",
	".#...",
	".....",
	"",
	"#####",
	"##.##",
	".#.##",
	"...##",
	"...#.",
	"...#.",
	".....",
	"",
	".....",
	"#....",
	"#....",
	"#...#",
	"#.#.#",
	"#.###",
	"#####",
	"",
	".....",
	".....",
	"#.#..",
	"###..",
	"###.#",
	"###.#",
	"#####",
	"",
	".....",
	".....",
	".....",
	"#....",
	"#.#..",
	"#.#.#",
	"#####",
}

func TestCountPossible(t *testing.T) {
	locks, keys := parseInput(ex)
	want := 3
	if got := countPossible(locks, keys); got != want {
		t.Errorf("countPossible(ex) = %d, want %d", got, want)
	}
}
