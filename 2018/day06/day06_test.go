package day06

import (
	"testing"

	"github.com/fis/aoc/util"
)

var example = []util.P{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}}

func TestMaxArea(t *testing.T) {
	want := 17
	got := maxArea(example)
	if got != want {
		t.Errorf("maxArea = %d, want %d", got, want)
	}
}

func TestSafeArea(t *testing.T) {
	want := 16
	got := safeArea(example, 32)
	if got != want {
		t.Errorf("safeArea = %d, want %d", got, want)
	}
}
