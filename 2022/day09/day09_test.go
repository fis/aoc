package day09

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex1 = []move{
	{dirRight, 4},
	{dirUp, 4},
	{dirLeft, 3},
	{dirDown, 1},
	{dirRight, 4},
	{dirDown, 1},
	{dirLeft, 5},
	{dirRight, 2},
}

var ex2 = []move{
	{dirRight, 5},
	{dirUp, 8},
	{dirLeft, 8},
	{dirDown, 3},
	{dirRight, 17},
	{dirDown, 10},
	{dirLeft, 25},
	{dirUp, 20},
}

func TestMeasureTail(t *testing.T) {
	want := 13
	if got := measureTailMap(ex1); got != want {
		t.Errorf("measureTail(ex1) = %d, want %d", got, want)
	}
}

func BenchmarkMeasureTail(b *testing.B) {
	lines, err := util.ReadLines("../../testdata/2022/day09.txt")
	if err != nil {
		b.Fatal(err)
	}
	moves, err := fn.MapE(lines, parseMove)
	if err != nil {
		b.Fatal(err)
	}
	algos := []struct {
		name string
		f    func(moves []move) int
	}{
		{"map", measureTailMap},
		{"bitmap", measureTail},
	}
	want := 6391
	for _, algo := range algos {
		b.Run("algo="+algo.name, func(b *testing.B) {
			for range b.N {
				if got := algo.f(moves); got != want {
					b.Errorf("%s(ex) = %d, want %d", algo.name, got, want)
				}
			}
		})
	}
}

func TestMeasureLongTail(t *testing.T) {
	tests := []struct {
		name  string
		moves []move
		want  int
	}{
		{"ex1", ex1, 1},
		{"ex2", ex2, 36},
	}
	for _, test := range tests {
		if got := measureLongTail(test.moves); got != test.want {
			t.Errorf("measureLongTail(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}

func BenchmarkMeasureLongTail(b *testing.B) {
	lines, err := util.ReadLines("../../testdata/2022/day09.txt")
	if err != nil {
		b.Fatal(err)
	}
	moves, err := fn.MapE(lines, parseMove)
	if err != nil {
		b.Fatal(err)
	}
	algos := []struct {
		name string
		f    func(moves []move) int
	}{
		{"map", measureLongTailMap},
		{"bitmap", measureLongTail},
	}
	want := 2593
	for _, algo := range algos {
		b.Run("algo="+algo.name, func(b *testing.B) {
			for range b.N {
				if got := algo.f(moves); got != want {
					b.Errorf("%s(ex) = %d, want %d", algo.name, got, want)
				}
			}
		})
	}
}
