package day14

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestAddSand(t *testing.T) {
	paths, err := fn.MapE(ex, parsePath)
	if err != nil {
		t.Fatal(err)
	}
	bmp, sourceX := buildBitmap(paths)
	want := 24
	if got := addSand(bmp, sourceX); got != want {
		t.Errorf("addSand(ex, %d) = %d, want %d", sourceX, got, want)
	}
}

func TestAddShadow(t *testing.T) {
	paths, err := fn.MapE(ex, parsePath)
	if err != nil {
		t.Fatal(err)
	}
	bmp, sourceX := buildBitmap(paths)
	want := 93
	if got := addShadow(bmp, sourceX); got != want {
		t.Errorf("addShadow(ex, %d) = %d, want %d", sourceX, got, want)
	}
}
