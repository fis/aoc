// Package day02 solves AoC 2021 day 2.
package day02

import (
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 2, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(forward|down|up) (\d+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	moves := parseInput(lines)
	p1 := applyMoves(util.P{0, 0}, moves)
	p2 := applyMoves2(util.P{0, 0}, moves)
	return glue.Ints(p1.X*p1.Y, p2.X*p2.Y), nil
}

func applyMoves(from util.P, moves []util.P) util.P {
	for _, move := range moves {
		from = from.Add(move)
	}
	return from
}

func applyMoves2(from util.P, moves []util.P) util.P {
	aim := 0
	for _, move := range moves {
		if move.Y == 0 {
			from.X += move.X
			from.Y += aim * move.X
		} else {
			aim += move.Y
		}
	}
	return from
}

func parseInput(lines [][]string) (moves []util.P) {
	commands := map[string]util.P{
		"forward": {1, 0},
		"down":    {0, 1},
		"up":      {0, -1},
	}
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1])
		moves = append(moves, commands[line[0]].Scale(n))
	}
	return moves
}
