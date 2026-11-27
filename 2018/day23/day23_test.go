package day23

import "testing"

func TestInRange(t *testing.T) {
	bots := []nanobot{
		{p3{0, 0, 0}, 4},
		{p3{1, 0, 0}, 1},
		{p3{4, 0, 0}, 3},
		{p3{0, 2, 0}, 1},
		{p3{0, 5, 0}, 3},
		{p3{0, 0, 3}, 1},
		{p3{1, 1, 1}, 1},
		{p3{1, 1, 2}, 1},
		{p3{1, 3, 1}, 1},
	}
	want := 7
	got := inRange(bots)
	if got != want {
		t.Errorf("inRange(%v) = %d, want %d", bots, got, want)
	}
}

func TestBestPos(t *testing.T) {
	bots := []nanobot{
		{p3{10, 12, 12}, 2},
		{p3{12, 14, 12}, 2},
		{p3{16, 12, 12}, 4},
		{p3{14, 14, 14}, 6},
		{p3{50, 50, 50}, 200},
		{p3{10, 10, 10}, 5},
	}
	want := 36
	got := bestPos(bots)
	if got != want {
		t.Errorf("bestPos(%v) = %d, want %d", bots, got, want)
	}
}
