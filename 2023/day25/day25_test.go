package day25

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestFindCutSizes(t *testing.T) {
	g := parseGraph(util.Lines(ex))
	wantL, wantR := 9, 6
	if l, r := findCutSizes(g); l != wantL || r != wantR {
		t.Errorf("findCut(ex) = (%d, %d), want (%d, %d)", l, r, wantL, wantR)
	}
}
