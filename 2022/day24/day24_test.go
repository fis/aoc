package day24

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []string{
	"#.######",
	"#>>.<^<#",
	"#.<..<<#",
	"#>v.><>#",
	"#<^v^^>#",
	"######.#",
}

func TestFindPath(t *testing.T) {
	lvl := parseLevel(ex)
	tests := []struct {
		t0         int
		start, end util.P
		want       int
	}{
		{0, util.P{0, -1}, util.P{5, 4}, 18},
		{18, util.P{5, 4}, util.P{0, -1}, 23},
		{18 + 23, util.P{0, -1}, util.P{5, 4}, 13},
	}
	for _, test := range tests {
		if got := findPath(lvl, test.t0, test.start, test.end); got != test.want {
			t.Errorf("findPath(ex, %d, %v, %v) = %d, want %d", test.t0, test.start, test.end, got, test.want)
		}
	}
}
