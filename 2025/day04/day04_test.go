package day04

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`, "\n")

func TestRemovePaper(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want1, want2 := 13, 43
	if got1, got2 := removePaper(level); got1 != want1 && got2 != want2 {
		t.Errorf("removePaper(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
