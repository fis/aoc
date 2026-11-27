// Package day09 solves AoC 2018 day 9.
package day09

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 9, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	var players, marbles int
	if _, err := fmt.Sscanf(lines[0], "%d players; last marble is worth %d points", &players, &marbles); err != nil {
		return nil, err
	}

	p1 := game(players, marbles)
	p2 := game(players, 100*marbles)

	return glue.Ints(p1, p2), nil
}

func game(players, marbles int) (maxScore int) {
	ccw := make([]int, marbles+1)
	cw := make([]int, marbles+1)
	m := 0
	// cw[0] = 0, ccw[0] = 0

	scores := make([]int, players)
	for turn := 1; turn <= marbles; turn++ {
		if turn%23 != 0 {
			m = cw[m]
			ccw[turn], cw[turn] = m, cw[m]
			ccw[cw[m]], cw[m] = turn, turn
			m = turn
		} else {
			pl := turn % players
			scores[pl] += turn
			m = ccw[ccw[ccw[ccw[ccw[ccw[ccw[m]]]]]]]
			scores[pl] += m
			nm := cw[m]
			cw[ccw[m]], ccw[cw[m]] = cw[m], ccw[m]
			m = nm
		}
	}

	maxScore = scores[0]
	for _, s := range scores[1:] {
		if s > maxScore {
			maxScore = s
		}
	}
	return maxScore
}
