package day15

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

type result struct{ elfPower, outcome int }

var examples = []struct {
	name    string
	level   string
	results []result
}{
	{
		name: "ex1",
		level: `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`,
		results: []result{
			{elfPower: 3, outcome: 27730},
			{elfPower: 15, outcome: 4988},
		},
	},
	{
		name: "ex2",
		level: `
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######
`,
		results: []result{{elfPower: 3, outcome: 36334}},
	},
	{
		name: "ex3",
		level: `
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
`,
		results: []result{
			{elfPower: 3, outcome: 39514},
			{elfPower: 4, outcome: 31284},
		},
	},
	{
		name: "ex4",
		level: `
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
`,
		results: []result{
			{elfPower: 3, outcome: 27755},
			{elfPower: 15, outcome: 3478},
		},
	},
	{
		name: "ex5",
		level: `
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`,
		results: []result{
			{elfPower: 3, outcome: 28944},
			{elfPower: 12, outcome: 6474},
		},
	},
	{
		name: "ex6",
		level: `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
`,
		results: []result{
			{elfPower: 3, outcome: 18740},
			{elfPower: 34, outcome: 1140},
		},
	},
}

func TestCombat(t *testing.T) {
	for _, test := range examples {
		for _, res := range test.results {
			level := util.ParseLevelString(strings.TrimSpace(test.level), '#')
			got := combat(level, res.elfPower, false)
			if got != res.outcome {
				t.Errorf("combat(%s, %d) = %d, want %d", test.name, res.elfPower, got, res.outcome)
			}
		}
	}
}
