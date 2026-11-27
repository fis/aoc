package day01

import (
	"slices"
	"testing"
)

var (
	exA = slices.Sorted(slices.Values([]int{3, 4, 2, 1, 3, 3}))
	exB = slices.Sorted(slices.Values([]int{4, 3, 5, 3, 9, 3}))
)

func TestTotalDistance(t *testing.T) {
	want := 11
	got := totalDistance(exA, exB)
	if got != want {
		t.Errorf("totalDistance(exA, exB) = %d, want %d", got, want)
	}
}

func TestSimilarity(t *testing.T) {
	want := 31
	got := similarity(exA, exB)
	if got != want {
		t.Errorf("similarity(exA, exB) = %d, want %d", got, want)
	}
}
