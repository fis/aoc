package day10

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
)

const (
	ex1 = `
.#..#
.....
#####
....#
...##
`
	ex2 = `
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`
	ex3 = `
#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
`
	ex4 = `
.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
`
	ex5 = `
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
`
)

func TestPart1(t *testing.T) {
	tests := []struct {
		comment string
		level   string
		wantAt  util.P
		wantVis int
	}{
		{
			comment: "small example",
			level:   ex1,
			wantAt:  util.P{3, 4},
			wantVis: 8,
		},
		{
			comment: "medium example 1",
			level:   ex2,
			wantAt:  util.P{5, 8},
			wantVis: 33,
		},
		{
			comment: "medium example 2",
			level:   ex3,
			wantAt:  util.P{1, 2},
			wantVis: 35,
		},
		{
			comment: "medium example 3",
			level:   ex4,
			wantAt:  util.P{6, 3},
			wantVis: 41,
		},
		{
			comment: "big example",
			level:   ex5,
			wantAt:  util.P{11, 13},
			wantVis: 210,
		},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), '.')
		at, vis := findBest(level)
		if at != test.wantAt || vis != test.wantVis {
			t.Errorf("%s: got (%v, %d), want (%v, %d)", test.comment, at, vis, test.wantAt, test.wantVis)
		}
	}
}

func TestPart2(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex5, "\n"), '.')
	at, _ := findBest(level)
	nth := findNth(at, 200, level)
	if at != (util.P{11, 13}) || nth != (util.P{8, 2}) {
		t.Errorf("got (%v, %v), want ({11 3} {8 2})", at, nth)
	}
}
