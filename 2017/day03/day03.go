// Package day03 solves AoC 2017 day 3.
package day03

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2017, 3, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	if len(input) != 1 {
		return nil, fmt.Errorf("expecting only 1 integer of input, got %d", len(input))
	}
	p1 := part1(input[0])
	p2 := part2(input[0])
	return glue.Ints(p1, p2), nil
}

func part1(square int) int {
	if square == 1 {
		return 0
	}
	square--
	d := ix.Sqrt(square)
	if d%2 == 0 {
		d--
	}
	square -= d * d
	r := (d + 1) / 2
	return r + ix.Abs(square%(2*r)-(r-1))
}

func part2(limit int) int {
	grid := map[util.P]int{{0, 0}: 1}
	r := 1
	for {
		p := util.P{X: r, Y: -r}
		for _, d := range []util.P{{0, 1}, {-1, 0}, {0, -1}, {1, 0}} {
			for range 2 * r {
				p = p.Add(d)
				v := 0
				for _, n := range p.Neigh8() {
					v += grid[n]
				}
				if v > limit {
					return v
				}
				grid[p] = v
			}
		}
		r++
	}
}
