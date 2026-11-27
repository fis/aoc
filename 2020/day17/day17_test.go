package day17

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var example = `
.#.
..#
###
`

func TestCycle3(t *testing.T) {
	state, min, max := loadLevel3(util.ParseLevelString(strings.TrimSpace(example), '.'))
	for range 6 {
		state, min, max = cycle3(state, min, max)
	}
	want := 112
	got := len(state)
	if got != want {
		t.Errorf("cycle3*6 -> %d, want %d", got, want)
	}
}

func TestCycle4(t *testing.T) {
	state, min, max := loadLevel4(util.ParseLevelString(strings.TrimSpace(example), '.'))
	for range 6 {
		state, min, max = cycle4(state, min, max)
	}
	want := 848
	got := len(state)
	if got != want {
		t.Errorf("cycle4*6 -> %d, want %d", got, want)
	}
}
