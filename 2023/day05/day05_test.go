package day05

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var al = parseAlmanac(util.Chunks(strings.TrimPrefix(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, "\n")))

func TestLowestSingle(t *testing.T) {
	want := 35
	if got := lowestSingle(al); got != want {
		t.Errorf("lowestSingle(al) = %d, want %d", got, want)
	}
}

func TestLowestRanged(t *testing.T) {
	want := 46
	if got := lowestRanged(al); got != want {
		t.Errorf("lowestRanged(al) = %d, want %d", got, want)
	}
}
func TestMapSeed(t *testing.T) {
	tests := []struct {
		seed int
		want int
	}{
		{79, 82},
		{14, 43},
		{55, 86},
		{13, 35},
	}
	for _, test := range tests {
		if got := al.mapSeed(test.seed); got != test.want {
			t.Errorf("al.mapSeed(%d) = %d, want %d", test.seed, got, test.want)
		}
	}
}
