package day17

import (
	"testing"
)

func TestPart1(t *testing.T) {
	skip := 3
	want := 638
	if got := part1(skip); got != want {
		t.Errorf("part1(%d) = %d, want %d", skip, got, want)
	}
}
