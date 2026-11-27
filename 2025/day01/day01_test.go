package day01

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func TestCountZeros(t *testing.T) {
	moves, err := fn.MapE(ex, parseMove)
	if err != nil {
		t.Fatal(err)
	}
	want1, want2 := 3, 6
	if got1, got2 := countZeros(moves); got1 != want1 || got2 != want2 {
		t.Errorf("countZeros(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
