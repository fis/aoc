package day01

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		steps string
		want  int
	}{
		{steps: "R2, L3", want: 5},
		{steps: "R2, R2, R2", want: 2},
		{steps: "R5, L5, R5, R3", want: 12},
	}
	for _, test := range tests {
		steps := strings.Split(test.steps, ", ")
		got, _ := find(steps)
		if got != test.want {
			t.Errorf("find(%v) = (%d, ?), want (%d, ?)", steps, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	steps := []string{"R8", "R4", "R4", "R8"}
	want := 4
	_, got := find(steps)
	if got != want {
		t.Errorf("find(%v) = (?, %d), want (?, %d)", steps, got, want)
	}
}
