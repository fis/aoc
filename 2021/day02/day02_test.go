package day02

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = [][]string{
	{"forward", "5"},
	{"down", "5"},
	{"forward", "8"},
	{"up", "3"},
	{"down", "8"},
	{"forward", "2"},
}

func TestApplyMoves(t *testing.T) {
	want := util.P{15, 10}
	if got := applyMoves(util.P{0, 0}, parseInput(ex)); got != want {
		t.Errorf("applyMoves(0, %v) = %d, want %d", ex, got, want)
	}
}

func TestApplyMoves2(t *testing.T) {
	want := util.P{15, 60}
	if got := applyMoves2(util.P{0, 0}, parseInput(ex)); got != want {
		t.Errorf("applyMoves2(0, %v) = %d, want %d", ex, got, want)
	}
}
