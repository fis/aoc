package day07

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestCountAncestors(t *testing.T) {
	bag := "shiny gold"
	want := 4
	if g, err := parseRules(util.Lines(ex1)); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if got := countAncestors(g, bag); got != want {
		t.Errorf("countAncestors(%s) = %d, want %d", bag, got, want)
	}
}

func TestCountDescendants(t *testing.T) {
	bag := "shiny gold"
	tests := []struct {
		name  string
		rules string
		want  int
	}{
		{"ex1", ex1, 32},
		{"ex2", ex2, 126},
	}
	for _, test := range tests {
		if g, err := parseRules(util.Lines(test.rules)); err != nil {
			t.Errorf("parseRules(%s): %v", test.name, err)
		} else if got := countDescendants(g, bag); got != test.want {
			t.Errorf("countDescendants(%s, %s) = %d, want %d", test.name, bag, got, test.want)
		}
	}
}
