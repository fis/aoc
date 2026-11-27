package day01

import (
	"testing"
)

func TestIncreases(t *testing.T) {
	depths := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want1, want2 := 7, 5
	if got1, got2 := increases(depths); got1 != want1 || got2 != want2 {
		t.Errorf("increases(%v) = (%d, %d), want (%d, %d)", depths, got1, got2, want1, want2)
	}
}
