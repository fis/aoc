package day17

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestFindShots(t *testing.T) {
	Tmin, Tmax := util.P{20, -10}, util.P{30, -5}
	want1, want2 := 45, 112
	if got1, got2 := findShots(Tmin, Tmax); got1 != want1 || got2 != want2 {
		t.Errorf("bestShot(%v, %v) = (%d, %d), want (%d, %d)", Tmin, Tmax, got1, got2, want1, want2)
	}
}
