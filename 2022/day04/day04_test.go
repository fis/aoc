package day04

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

func TestPart1(t *testing.T) {
	pairs, _ := fn.MapE(ex, parsePair)
	want := 2
	if got := part1(pairs); got != want {
		t.Errorf("part1(%v) = %d, want %d", pairs, got, want)
	}
}

func TestPart2(t *testing.T) {
	pairs, _ := fn.MapE(ex, parsePair)
	want := 4
	if got := part2(pairs); got != want {
		t.Errorf("part2(%v) = %d, want %d", pairs, got, want)
	}
}
