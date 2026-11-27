package day23

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := 57
	want := 3025
	if got := part1(n); got != want {
		t.Errorf("part1(%d) = %d, want %d", n, got, want)
	}
}

func TestCountComposite(t *testing.T) {
	low, high, skip := 105700, 122700, 17
	want := 915
	if got := countComposite(low, high, skip); got != want {
		t.Errorf("countComposite(%d, %d, %d) = %d, want %d", low, high, skip, got, want)
	}
}
