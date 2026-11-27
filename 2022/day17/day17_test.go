package day17

import (
	"testing"
)

var ex = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestDropRocks(t *testing.T) {
	want := 3068
	got := dropRocks(ex, 2022)
	if got != want {
		t.Errorf("dropRocks(ex, 2022) = %d, want %d", got, want)
	}
}

func TestAnalyzeRocks(t *testing.T) {
	want := 1514285714288
	got := analyzeRocks(ex, 1000000000000)
	if got != want {
		t.Errorf("analyzeRocks(ex, 1000000000000) = %d, want %d", got, want)
	}
}
