// Package day18 solves AoC 2023 day 18.
package day18

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = [][]string{
	{"R", "6", "70c710"},
	{"D", "5", "0dc571"},
	{"L", "2", "5713f0"},
	{"D", "2", "d2c081"},
	{"R", "2", "59c680"},
	{"D", "2", "411b91"},
	{"L", "5", "8ceee2"},
	{"U", "2", "caa173"},
	{"L", "1", "1b58a2"},
	{"U", "2", "caa171"},
	{"R", "2", "7807d2"},
	{"U", "3", "a77fa3"},
	{"L", "2", "015232"},
	{"U", "2", "7a21e3"},
}

func TestPolygonArea(t *testing.T) {
	tests := []struct {
		name string
		f    func(instruction) (direction, int)
		want int
	}{
		{"small", instruction.small, 62},
		{"big", instruction.big, 952408144115},
	}
	plan, err := fn.MapE(ex, parseDigPlan)
	if err != nil {
		t.Fatalf("parseDigPlan(ex): %v", err)
	}
	for _, test := range tests {
		if got := polygonArea(traceLoop(plan, test.f)); got != test.want {
			t.Errorf("polygonArea(ex, %s) = %d, want %d", test.name, got, test.want)
		}
	}
}
