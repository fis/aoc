package day02

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestFindPossible(t *testing.T) {
	games, err := fn.MapE(ex, parseGame)
	if err != nil {
		t.Fatal(err)
	}
	cubes := cubeCount{colorRed: 12, colorGreen: 13, colorBlue: 14}
	want := 8
	if got := findPossible(games, cubes); got != want {
		t.Errorf("findPossible(ex, %v) = %d, want %d", cubes, got, want)
	}
}

func TestFindPower(t *testing.T) {
	games, err := fn.MapE(ex, parseGame)
	if err != nil {
		t.Fatal(err)
	}
	want := 2286
	if got := findPower(games); got != want {
		t.Errorf("findPower(ex) = %d, want %d", got, want)
	}
}
