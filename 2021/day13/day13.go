// Package day13 solves AoC 2021 day 13.
package day13

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 13, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(?:(\d+),(\d+)|fold along ([xy])=(\d+)|)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	points, folds := parseInput(lines)
	folds[0].apply(points)
	p1 := countVisible(points)
	for _, fold := range folds[1:] {
		fold.apply(points)
	}
	p2 := printPoints(points)
	return append(glue.Ints(p1), p2...), nil
}

type foldSpec struct {
	vertical bool
	at       int
}

func (fold foldSpec) apply(points []util.P) {
	if fold.vertical {
		for i := range points {
			if x := points[i].X; x > fold.at {
				points[i].X = fold.at - (x - fold.at)
			}
		}
	} else {
		for i := range points {
			if y := points[i].Y; y > fold.at {
				points[i].Y = fold.at - (y - fold.at)
			}
		}
	}
}

func countVisible(points []util.P) int {
	visible := map[util.P]struct{}{}
	for _, p := range points {
		visible[p] = struct{}{}
	}
	return len(visible)
}

func printPoints(points []util.P) []string {
	level := util.SparseLevel(util.P{0, 0}, ' ')
	for _, p := range points {
		level.Set(p.X, p.Y, '#')
	}
	buf := strings.Builder{}
	level.Write(&buf)
	return util.Lines(buf.String())
}

func parseInput(lines [][]string) (points []util.P, folds []foldSpec) {
	for _, line := range lines {
		if line[0] != "" {
			x, _ := strconv.Atoi(line[0])
			y, _ := strconv.Atoi(line[1])
			points = append(points, util.P{x, y})
		} else if line[2] != "" {
			at, _ := strconv.Atoi(line[3])
			folds = append(folds, foldSpec{vertical: line[2][0] == 'x', at: at})
		}
	}
	return points, folds
}
