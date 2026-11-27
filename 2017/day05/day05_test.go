package day05

import (
	"testing"
)

func TestPart1(t *testing.T) {
	offsets := []int{0, 3, 0, 1, -3}
	want := 5
	if got := part1(offsets); got != want {
		t.Errorf("part1(%v) = %d, want %d", offsets, got, want)
	}
}

func TestPart2(t *testing.T) {
	offsets := []int{0, 3, 0, 1, -3}
	want := 10
	if got := part2(offsets); got != want {
		t.Errorf("part2(%v) = %d, want %d", offsets, got, want)
	}
}
