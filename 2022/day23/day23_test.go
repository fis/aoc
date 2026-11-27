package day23

import (
	"testing"
)

var (
	ex1 = []string{
		".....",
		"..##.",
		"..#..",
		".....",
		"..##.",
		".....",
	}
	ex2 = []string{
		"....#..",
		"..###.#",
		"#...#.#",
		".#...##",
		"#.###..",
		"##.#.##",
		".#..#..",
	}
)

func TestSimulate(t *testing.T) {
	tests := []struct {
		name         string
		lines        []string
		want1, want2 int
	}{
		{"ex1", ex1, 0, 4},
		{"ex2", ex2, 110, 20},
	}
	for _, test := range tests {
		if got1, got2 := simulate(test.lines); got1 != test.want1 || got2 != test.want2 {
			t.Errorf("simulate(%s) = (%d, %d), want (%d, %d)", test.name, got1, got2, test.want1, test.want2)
		}
	}
}
