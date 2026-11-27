// Package day01 solves AoC 2025 day 1.
package day01

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2025, 1, glue.LineSolver(glue.WithParser(parseMove, solve)))
}

func solve(moves []move) ([]string, error) {
	p1, p2 := countZeros(moves)
	return glue.Ints(p1, p2), nil
}

func countZeros(moves []move) (landed, clicked int) {
	dial := 50
	for _, m := range moves {
		clicked += m.distance / 100
		dist := m.distance % 100
		switch m.dir {
		case 'L':
			if dial != 0 && dial-dist <= 0 {
				clicked++
			}
			dial = (dial + 100 - dist) % 100
		case 'R':
			if dial != 0 && dial+dist >= 100 {
				clicked++
			}
			dial = (dial + dist) % 100
		}
		if dial == 0 {
			landed++
		}
	}
	return landed, clicked
}

type move struct {
	dir      byte
	distance int
}

func parseMove(line string) (m move, err error) {
	if len(line) == 0 {
		return move{}, fmt.Errorf("no move")
	} else if line[0] != 'L' && line[0] != 'R' {
		return move{}, fmt.Errorf("not a valid direction: %c", line[0])
	}
	m.dir = line[0]
	m.distance, err = strconv.Atoi(line[1:])
	if err != nil {
		return move{}, fmt.Errorf("not a valid distance: %q", line[1:])
	}
	return m, nil
}
