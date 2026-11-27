package day06

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoopLength(t *testing.T) {
	banks := []byte{0, 2, 7, 0}
	want1, want2 := 5, 4
	if got := loopLength(banks); got != want1 {
		t.Errorf("loopLength (1st run) = %d, want %d", got, want1)
	}
	if got := loopLength(banks); got != want2 {
		t.Errorf("loopLength (2nd run) = %d, want %d", got, want2)
	}
}

func TestBalance(t *testing.T) {
	steps := [][]byte{
		{0, 2, 7, 0},
		{2, 4, 1, 2},
		{3, 1, 2, 3},
		{0, 2, 3, 4},
		{1, 3, 4, 1},
		{2, 4, 1, 2},
	}
	for i := range len(steps) - 1 {
		banks, want := steps[i], steps[i+1]
		got := append(banks[:0:0], banks...)
		balance(got)
		if !cmp.Equal(got, want) {
			t.Errorf("balance(%v) -> %v, want %v", banks, got, want)
		}
	}
}
