package day09

import (
	"slices"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []util.P{
	{7, 1},
	{11, 1},
	{11, 7},
	{9, 7},
	{9, 5},
	{2, 5},
	{2, 3},
	{7, 3},
}

func TestLargestRect(t *testing.T) {
	want := 50
	if got := largestRect(ex); got != want {
		t.Errorf("largestRect(ex) = %d, want %d", got, want)
	}
}

func TestLargestValid(t *testing.T) {
	want := 24
	if got := largestValid(ex); got != want {
		t.Errorf("largestValid(ex) = %d, want %d", got, want)
	}
}

func TestLargestValidReversed(t *testing.T) {
	rev := slices.Clone(ex)
	slices.Reverse(rev)
	want := 24
	if got := largestValid(rev); got != want {
		t.Errorf("largestValid(rev) = %d, want %d", got, want)
	}
}
