package day13

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestShortestPath(t *testing.T) {
	key, from, to, want := 10, util.P{1, 1}, util.P{7, 4}, 11
	if got := shortestPath(key, from, to); got != want {
		t.Errorf("shortestPath(%d, %v, %v) = %d, want %d", key, from, to, got, want)
	}
}
