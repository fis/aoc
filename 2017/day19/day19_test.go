package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
    |
    |  +--+
    A  |  C
F---|----E|--+
    |  |  |  D
    +B-+  +--+
`

func TestTracePath(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want1, want2 := "ABCDEF", 38
	if got1, got2 := tracePath(level); got1 != want1 || got2 != want2 {
		t.Errorf("tracePath = (%s, %d), want (%s, %d)", got1, got2, want1, want2)
	}
}
