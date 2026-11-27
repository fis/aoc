package day04

import "testing"

var ex = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestCountMatches(t *testing.T) {
	want := 18
	if got := countMatches(ex, "XMAS"); got != want {
		t.Errorf("countMatches(ex, XMAS) = %d, want %d", got, want)
	}
}

func TestCountXmas(t *testing.T) {
	want := 9
	if got := countXmas(ex); got != want {
		t.Errorf("countXmas(ex) = %d, want %d", got, want)
	}
}
