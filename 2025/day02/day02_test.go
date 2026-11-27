package day02

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = [][2]int{
	{11, 22},
	{95, 115},
	{998, 1012},
	{1188511880, 1188511890},
	{222220, 222224},
	{1698522, 1698528},
	{446443, 446449},
	{38593856, 38593862},
	{565653, 565659},
	{824824821, 824824827},
	{2121212118, 2121212124},
}

func TestCountInvalid(t *testing.T) {
	want1, want2 := 1227775554, 4174379265
	if got1, got2 := countInvalid(ex); got1 != want1 || got2 != want2 {
		t.Errorf("countInvalid(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}

func TestCountInvalidFast(t *testing.T) {
	want1, want2 := 1227775554, 4174379265
	if got1, got2 := countInvalidFast(ex); got1 != want1 || got2 != want2 {
		t.Errorf("countInvalidFast(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}

func BenchmarkCountInvalid(b *testing.B) {
	lines, err := util.ReadLines("../../testdata/2025/day02.txt")
	if err != nil {
		b.Fatal(err)
	}
	ranges, err := parseRanges(lines[0])
	if err != nil {
		b.Fatal(err)
	}
	algos := []struct {
		name string
		f    func(ranges [][2]int) (p1, p2 int)
	}{
		{name: "basic", f: countInvalid},
		{name: "fast", f: countInvalidFast},
	}
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for b.Loop() {
				if got1, got2 := algo.f(ranges); got1 != 23701357374 || got2 != 34284458938 {
					b.Errorf("%s = (%d, %d), want (23701357374, 34284458938)", algo.name, got1, got2)
				}
			}
		})
	}
}
