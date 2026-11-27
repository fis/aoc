package day22

import (
	"testing"
)

var ex = []string{
	"        ...#",
	"        .#..",
	"        #...",
	"        ....",
	"...#.......#",
	"........#...",
	"..#....#....",
	"..........#.",
	"        ...#....",
	"        .....#..",
	"        .#......",
	"        ......#.",
	"",
	"10R5L5R10L4R5L5",
}

func TestDecode(t *testing.T) {
	lvl, path, err := parseInput(ex)
	if err != nil {
		t.Fatalf("parseInput(ex): %v", err)
	}
	want := 6032
	if got := decode(lvl, path); got != want {
		t.Errorf("decode(ex) = %d, want %d", got, want)
	}
}

func TestDecodeCube(t *testing.T) {
	lvl, path, err := parseInput(ex)
	if err != nil {
		t.Fatalf("parseInput(ex): %v", err)
	}
	c, err := fold(lvl, 4)
	if err != nil {
		t.Fatalf("fold(ex, 4): %v", err)
	}
	want := 5031
	if got := decodeCube(c, path); got != want {
		t.Errorf("decodeCube(ex) = %d, want %d", got, want)
	}
}
