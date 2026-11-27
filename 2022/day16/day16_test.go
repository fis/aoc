package day16

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func TestFindOne(t *testing.T) {
	lines, err := util.ReadLines("../../testdata/2022/day16.txt")
	if err != nil {
		t.Fatal(err)
	}
	scan, _ := fn.MapE(lines, ParseValveScan)
	sum := Preprocess(scan)
	want := 1789 // 2496
	if got := findOne(sum, 30); got != want {
		t.Errorf("findOne(day16, 30) = %d, want %d", got, want)
	}
}

func TestFindTwo(t *testing.T) {
	lines, err := util.ReadLines("../../testdata/2022/day16.txt")
	if err != nil {
		t.Fatal(err)
	}
	scan, _ := fn.MapE(lines, ParseValveScan)
	sum := Preprocess(scan)
	want := 2496
	if got := findTwo(sum, 26); got != want {
		t.Errorf("findTwo(day16, 26) = %d, want %d", got, want)
	}
}
