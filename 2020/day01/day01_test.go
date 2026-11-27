package day01

import (
	"testing"
)

var input = []int{1721, 979, 366, 299, 675, 1456}

func TestPart1(t *testing.T) {
	want := 514579
	got := part1(input)
	if got != want {
		t.Errorf("part1 = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 241861950
	got := part2(input)
	if got != want {
		t.Errorf("part2 = %d, want %d", got, want)
	}
}
