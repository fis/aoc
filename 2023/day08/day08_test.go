package day08

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestCountSteps(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ex1", ex1, 2},
		{"ex2", ex2, 6},
	}
	for _, test := range tests {
		if g, dirs, err := parseMaps(util.Chunks(test.input)); err != nil {
			t.Errorf("parseMaps(%s): %v", test.name, err)
		} else if got := countSteps(g, g.v("AAA"), g.v("ZZZ"), dirs); got != test.want {
			t.Errorf("countSteps(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}

func TestCountGhostSteps(t *testing.T) {
	want := 6
	if g, dirs, err := parseMaps(util.Chunks(ex3)); err != nil {
		t.Errorf("parseMaps(ex3): %v", err)
	} else if got := countGhostSteps(g, dirs); got != want {
		t.Errorf("countSteps(ex3) = %d, want %d", got, want)
	}
}
