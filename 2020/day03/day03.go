// Package day03 solves AoC 2020 day 3.
package day03

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2020, 3, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(level *util.Level) ([]string, error) {
	trees := countTrees(level, util.P{3, 1})

	allTrees := trees
	for _, slope := range []util.P{{1, 1}, {5, 1}, {7, 1}, {1, 2}} {
		allTrees *= countTrees(level, slope)
	}

	return glue.Ints(trees, allTrees), nil
}

func countTrees(level *util.Level, slope util.P) int {
	trees := 0
	_, max := level.Bounds()
	for x, y, w := 0, 0, max.X+1; y <= max.Y; x, y = (x+slope.X)%w, y+slope.Y {
		if level.At(x, y) == '#' {
			trees++
		}
	}
	return trees
}
