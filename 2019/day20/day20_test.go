package day20

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
)

const (
	ex1 = `
         A
         A
  #######.#########
  #######.........#
  #######.#######.#
  #######.#######.#
  #######.#######.#
  #####  B    ###.#
BC...##  C    ###.#
  ##.##       ###.#
  ##...DE  F  ###.#
  #####    G  ###.#
  #########.#####.#
DE..#######...###.#
  #.#########.###.#
FG..#########.....#
  ###########.#####
             Z
             Z
`
	ex2 = `
                   A
                   A
  #################.#############
  #.#...#...................#.#.#
  #.#.#.###.###.###.#########.#.#
  #.#.#.......#...#.....#.#.#...#
  #.#########.###.#####.#.#.###.#
  #.............#.#.....#.......#
  ###.###########.###.#####.#.#.#
  #.....#        A   C    #.#.#.#
  #######        S   P    #####.#
  #.#...#                 #......VT
  #.#.#.#                 #.#####
  #...#.#               YN....#.#
  #.###.#                 #####.#
DI....#.#                 #.....#
  #####.#                 #.###.#
ZZ......#               QG....#..AS
  ###.###                 #######
JO..#.#.#                 #.....#
  #.#.#.#                 ###.#.#
  #...#..DI             BU....#..LF
  #####.#                 #.#####
YN......#               VT..#....QG
  #.###.#                 #.###.#
  #.#...#                 #.....#
  ###.###    J L     J    #.#.###
  #.....#    O F     P    #.#...#
  #.###.#####.#.#####.#####.###.#
  #...#.#.#...#.....#.....#.#...#
  #.#####.###.###.#.#.#########.#
  #...#.#.....#...#.#.#.#.....#.#
  #.###.#####.###.###.#.#.#######
  #.#.........#...#.............#
  #########.###.###.#############
           B   J   C
           U   P   P
`
	ex3 = `
             Z L X W       C
             Z P Q B       K
  ###########.#.#.#.#######.###############
  #...#.......#.#.......#.#.......#.#.#...#
  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###
  #.#...#.#.#...#.#.#...#...#...#.#.......#
  #.###.#######.###.###.#.###.###.#.#######
  #...#.......#.#...#...#.............#...#
  #.#########.#######.#.#######.#######.###
  #...#.#    F       R I       Z    #.#.#.#
  #.###.#    D       E C       H    #.#.#.#
  #.#...#                           #...#.#
  #.###.#                           #.###.#
  #.#....OA                       WB..#.#..ZH
  #.###.#                           #.#.#.#
CJ......#                           #.....#
  #######                           #######
  #.#....CK                         #......IC
  #.###.#                           #.###.#
  #.....#                           #...#.#
  ###.###                           #.#.#.#
XF....#.#                         RF..#.#.#
  #####.#                           #######
  #......CJ                       NM..#...#
  ###.#.#                           #.###.#
RE....#.#                           #......RF
  ###.###        X   X       L      #.#.#.#
  #.....#        F   Q       P      #.#.#.#
  ###.###########.###.#######.#########.###
  #.....#...#.....#.......#...#.....#.#...#
  #####.#.###.#######.#######.###.###.#.#.#
  #.......#.......#.#.#.#.#...#...#...#.#.#
  #####.###.#####.#.#.#.#.###.###.#.###.###
  #.......#.....#.#...#...............#...#
  #############.#.#.###.###################
               A O F   N
               A A D   M
`
)

func TestShortest(t *testing.T) {
	tests := []struct {
		comment string
		level   string
		want    int
	}{
		{
			comment: "example 1",
			level:   ex1,
			want:    23,
		},
		{
			comment: "example 2",
			level:   ex2,
			want:    58,
		},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), ' ')
		dist := distances(level)
		got := shortest(label{name: "AA", outer: true}, label{name: "ZZ", outer: true}, dist)
		if got != test.want {
			t.Errorf("%s: got %d, want %d", test.comment, got, test.want)
		}
	}
}

func TestRecursive(t *testing.T) {
	tests := []struct {
		comment string
		level   string
		want    int
	}{
		{
			comment: "example 1",
			level:   ex1,
			want:    26,
		},
		{
			comment: "example 3",
			level:   ex3,
			want:    396,
		},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), ' ')
		dist := distances(level)
		got := recursive(label{name: "AA", outer: true}, label{name: "ZZ", outer: true}, dist)
		if got != test.want {
			t.Errorf("%s: got %d, want %d", test.comment, got, test.want)
		}
	}
}

func TestDistances(t *testing.T) {
	tests := []struct {
		comment string
		level   string
		want    int
	}{
		{
			comment: "example 1",
			level:   ex1,
			want:    26,
		},
		{
			comment: "example 2",
			level:   ex2,
			want:    0,
		},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.level, "\n"), ' ')
		dist := distances(level)
		got := dist[label{name: "AA", outer: true}][label{name: "ZZ", outer: true}].d
		if got != test.want {
			t.Errorf("%s: got %d, want %d", test.comment, got, test.want)
		}
	}
}
