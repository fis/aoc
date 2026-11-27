package day23

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestLongestPath(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	g, _, _ := deconstruct(l)
	want := 94
	if got := longestPath(g); got != want {
		t.Errorf("longestPath(ex) = %d, want %d", got, want)
	}
}

func TestUnsafeLongestPath(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	g, startV, endV := deconstruct(l)
	want := 154
	if got := unsafeLongestPath(g, startV, endV); got != want {
		t.Errorf("evenLongestPath(ex) = %d, want %d", got, want)
	}
}
