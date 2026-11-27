package day10

import (
	"testing"
)

func TestPart1(t *testing.T) {
	N := 5
	lengths := []byte{3, 4, 1, 5}
	want := 12
	if got := part1(N, lengths); got != want {
		t.Errorf("part1(%d, %v) = %d, want %d", N, lengths, got, want)
	}
}
