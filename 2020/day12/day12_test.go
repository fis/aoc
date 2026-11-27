package day12

import (
	"testing"
)

var ex = []string{"F10", "N3", "F7", "R90", "F11"}

func TestTurtle(t *testing.T) {
	x := newTurtle()
	x.move(parseInput(ex))
	want := 25
	got := x.distance()
	if got != want {
		t.Errorf("distance = %d, want %d", got, want)
	}
}

func TestShip(t *testing.T) {
	x := newShip()
	x.move(parseInput(ex))
	want := 286
	got := x.distance()
	if got != want {
		t.Errorf("distance = %d, want %d", got, want)
	}
}
