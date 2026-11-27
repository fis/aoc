// Package day03 solves AoC 2019 day 3.
package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2019, 3, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 2 {
		return nil, fmt.Errorf("unexpected amount of lines: %v", lines)
	}
	p1, p2 := compute(lines[0], lines[1])
	return glue.Ints(p1, p2), nil
}

func compute(w1, w2 string) (closest, best int) {
	m := make(map[util.P]int)
	walk(w1, func(x, y, s1 int) {
		p := util.P{x, y}
		if _, ok := m[p]; ok {
			return
		}
		m[p] = s1
	})
	closest, best = -1, -1
	walk(w2, func(x, y, s2 int) {
		s1, ok := m[util.P{x, y}]
		if !ok {
			return
		}
		td := ix.Abs(x) + ix.Abs(y)
		if closest == -1 || td < closest {
			closest = td
		}
		ts := s1 + s2
		if best == -1 || ts < best {
			best = ts
		}
	})
	return closest, best
}

var dir = map[byte]util.P{
	'L': {-1, 0},
	'R': {1, 0},
	'U': {0, -1},
	'D': {0, 1},
}

func walk(w string, cb func(x, y, s int)) {
	x, y, s := 0, 0, 0
	for _, op := range strings.Split(w, ",") {
		d := dir[op[0]]
		n, _ := strconv.Atoi(op[1:])
		for range n {
			x, y, s = x+d.X, y+d.Y, s+1
			cb(x, y, s)
		}
	}
}
