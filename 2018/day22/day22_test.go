package day22

import "testing"

func TestTotalRisk(t *testing.T) {
	want := 114
	got := newCave(510, 10, 10).totalRisk()
	if got != want {
		t.Errorf("cave(510, 10, 10).totalRisk() = %d, want %d", got, want)
	}
}

func TestShortestPath(t *testing.T) {
	want := 45
	got := newCave(510, 10, 10).shortestPath()
	if got != want {
		t.Errorf("cave(510, 10, 10).shortestPath() = %d, want %d", got, want)
	}
}
