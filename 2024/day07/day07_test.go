package day07

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestSumPossible(t *testing.T) {
	eqs, _ := fn.MapE(ex, parseEquation)
	want := 3749
	if got := sumPossible(eqs); got != want {
		t.Errorf("sumPossible(ex) = %d, want %d", got, want)
	}
}

func TestSumPossible3(t *testing.T) {
	eqs, _ := fn.MapE(ex, parseEquation)
	want := 11387
	if got := sumPossible3(eqs); got != want {
		t.Errorf("sumPossible3(ex) = %d, want %d", got, want)
	}
}
