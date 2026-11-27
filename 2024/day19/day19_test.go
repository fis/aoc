package day19

import (
	"testing"
)

var (
	towels  = []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	designs = []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}
)

func TestCountPossible(t *testing.T) {
	want1, want2 := 6, 16
	if got1, got2 := countPossible(designs, towels); got1 != want1 || got2 != want2 {
		t.Errorf("countPossible(ex) = (%d,%d), want (%d,%d)", got1, got2, want1, want2)
	}
}
