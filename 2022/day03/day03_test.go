package day03

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestPart1(t *testing.T) {
	sacks, _ := fn.MapE(ex, pack)
	want := 157
	if got := part1(sacks); got != want {
		t.Errorf("part1(%v) = %d, want %d", sacks, got, want)
	}
}

func TestPart2(t *testing.T) {
	sacks, _ := fn.MapE(ex, pack)
	want := 70
	if got := part2(sacks); got != want {
		t.Errorf("part2(%v) = %d, want %d", sacks, got, want)
	}
}
