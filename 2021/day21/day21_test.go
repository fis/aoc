package day21

import "testing"

func TestPart1(t *testing.T) {
	want := 739785
	if got := part1(4, 8); got != want {
		t.Errorf("part1(4, 8) = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 444356092776315
	if got := part2(4, 8); got != want {
		t.Errorf("part2(4, 8) = %d, want %d", got, want)
	}
}
