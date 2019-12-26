package day12

import (
	"strings"
	"testing"
)

const (
	ex1 = `
<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`
	ex2 = `
<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`
)

func TestPart1(t *testing.T) {
	tests := []struct {
		comment string
		initial string
		steps   int
		want    int
	}{
		{
			comment: "example 1",
			initial: ex1,
			steps:   10,
			want:    179,
		},
		{
			comment: "example 2",
			initial: ex2,
			steps:   100,
			want:    1940,
		},
	}
	for _, test := range tests {
		state := parseState(strings.Split(strings.TrimSpace(test.initial), "\n"))
		run(state, test.steps)
		got := totalEnergy(state)
		if got != test.want {
			t.Errorf("%s: after %d steps, got %d, want %d", test.comment, test.steps, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		comment string
		initial string
		want    int
	}{
		{
			comment: "example 1",
			initial: ex1,
			want:    2772,
		},
		{
			comment: "example 2",
			initial: ex2,
			want:    4686774924,
		},
	}
	for _, test := range tests {
		state := parseState(strings.Split(strings.TrimSpace(test.initial), "\n"))
		got := cycle(state)
		if got != test.want {
			t.Errorf("%s: got %d, want %d", test.comment, got, test.want)
		}
	}
}
