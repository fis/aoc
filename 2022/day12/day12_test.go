package day12

import (
	"testing"
)

var ex = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestShortestPath(t *testing.T) {
	m, start, end := readMap(ex)
	want := 31
	if got := shortestPath(m, start, end); got != want {
		t.Errorf("shortestPath(ex) = %d, want %d", got, want)
	}
}

func TestScenicPath(t *testing.T) {
	m, _, end := readMap(ex)
	want := 29
	if got := scenicPath(m, end); got != want {
		t.Errorf("scenicPath(ex) = %d, want %d", got, want)
	}
}
