package day11

import (
	"strings"
	"testing"
)

func TestDistances(t *testing.T) {
	tests := []struct {
		steps string
		want1 int
		want2 int
	}{
		{steps: "ne,ne,ne", want1: 3, want2: 3},
		{steps: "ne,ne,sw,sw", want1: 0, want2: 2},
		{steps: "ne,ne,s,s", want1: 2, want2: 2},
		{steps: "se,sw,se,sw,sw", want1: 3, want2: 3},
	}
	for _, test := range tests {
		if got1, got2, err := distances(strings.Split(test.steps, ",")); err != nil {
			t.Errorf("part1(%s): %v", test.steps, err)
		} else if got1 != test.want1 || got2 != test.want2 {
			t.Errorf("part1(%s) = (%d, %d), want (%d, %d)", test.steps, got1, got2, test.want1, test.want2)
		}
	}
}
