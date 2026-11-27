// Package day03 solves AoC 2018 day 3.
package day03

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 3, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	claims, err := parseClaims(lines)
	if err != nil {
		return nil, err
	}

	overlap, valid := totalOverlap(claims)

	return glue.Ints(overlap, valid), nil
}

type claim struct {
	ID        int
	Pos, Size util.P
}

func parseClaims(lines []string) (out []claim, err error) {
	for _, line := range lines {
		c, err := parseClaim(line)
		if err != nil {
			return nil, fmt.Errorf("bad claim: %q: %w", line, err)
		}
		out = append(out, c)
	}
	return out, nil
}

func parseClaim(line string) (c claim, err error) {
	_, err = fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.ID, &c.Pos.X, &c.Pos.Y, &c.Size.X, &c.Size.Y)
	return c, err
}

func totalOverlap(claims []claim) (overlap int, valid int) {
	sq := make(map[util.P]int)
	for _, c := range claims {
		for y := c.Pos.Y; y < c.Pos.Y+c.Size.Y; y++ {
			for x := c.Pos.X; x < c.Pos.X+c.Size.X; x++ {
				p := util.P{x, y}
				sq[p] = sq[p] + 1
			}
		}
	}
	for _, c := range sq {
		if c > 1 {
			overlap++
		}
	}
next:
	for _, c := range claims {
		for y := c.Pos.Y; y < c.Pos.Y+c.Size.Y; y++ {
			for x := c.Pos.X; x < c.Pos.X+c.Size.X; x++ {
				if sq[util.P{x, y}] > 1 {
					continue next
				}
			}
		}
		valid = c.ID
		break
	}
	return overlap, valid
}
