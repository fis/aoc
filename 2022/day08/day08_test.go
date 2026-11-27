package day08

import (
	"testing"
)

var ex = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestCountVisible(t *testing.T) {
	want := 21
	if got := countVisible(ex); got != want {
		t.Errorf("countVisible(ex) = %d, want %d", got, want)
	}
}

func TestFindBest(t *testing.T) {
	want := 8
	if got := findBest(ex); got != want {
		t.Errorf("findBest(ex) = %d, want %d", got, want)
	}
}
