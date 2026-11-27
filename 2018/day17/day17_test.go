package day17

import (
	"testing"

	"github.com/fis/aoc/util"
)

var example = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
`

func TestFill(t *testing.T) {
	lines := util.Lines(example)
	level, minY, maxY := readScan(lines)
	fill(level, util.P{500, 0}, make(map[util.P]struct{}))
	want1, want2 := 57, 29
	got1, got2 := measureWater(level, minY, maxY)
	if got1 != want1 || got2 != want2 {
		t.Errorf("measureWater = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
