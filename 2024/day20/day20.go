// Package day20 solves AoC 2024 day 20.
package day20

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 20, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	part1, part2 := analyze(level, 100)
	return glue.Ints(part1, part2), nil
}

const (
	cheatDist = 20
	chunkBits = 3
)

type pathStep struct {
	p util.P
	d int
}

func analyze(level *util.FixedLevel, minCheat int) (shortCheats, longCheats int) {
	path := make([]util.P, 0, level.W*level.H/2)

	x, y, _ := level.Find('S')
	prev, next := util.P{x, y}, util.P{}
	for level.At(x, y) != 'E' {
		path = append(path, util.P{x, y})
		for _, next = range (util.P{x, y}).Neigh() {
			if nc := level.At(next.X, next.Y); next != prev && (nc == '.' || nc == 'E') {
				break
			}
		}
		x, y, prev = next.X, next.Y, util.P{x, y}
	}
	path = append(path, util.P{x, y})

	c0x := (1 - level.H) >> chunkBits
	c1x := (level.W - 1) >> chunkBits
	c1y := (level.W + level.H - 2) >> chunkBits
	chunks := util.Make2D[[]pathStep](c1x-c0x+1, c1y+1)
	for d, p := range path {
		cx, cy := chunkCoord(p.X, p.Y)
		chunks[cy][cx-c0x] = append(chunks[cy][cx-c0x], pathStep{p: p, d: d})
	}

	for i := 0; i < len(path)-minCheat; i++ {
		s := path[i]
		s0x, s0y := chunkCoord(s.X-cheatDist, s.Y)
		s1x, s1y := chunkCoord(s.X+cheatDist, s.Y)
		s0x, s0y, s1x, s1y = max(s0x, c0x), max(s0y, 0), min(s1x, c1x), min(s1y, c1y)
		for cy := s0y; cy <= s1y; cy++ {
			for cx := s0x; cx <= s1x; cx++ {
				chunk := chunks[cy][cx-c0x]
				j := startPos(chunk, i+minCheat)
				for ; j < len(chunk); j++ {
					e := chunk[j]
					d := util.DistM(s, e.p)
					if cheat := (e.d - i) - d; cheat >= minCheat {
						if d == 2 {
							shortCheats++
						}
						if d <= cheatDist {
							longCheats++
						}
					}
				}
			}
		}
	}

	return shortCheats, longCheats
}

func startPos(chunk []pathStep, minD int) int {
	n := len(chunk)
	a, b := 0, n
	for a < b {
		h := (a + b) >> 1
		if chunk[h].d < minD {
			a = h + 1
		} else {
			b = h
		}
	}
	return a
}

func chunkCoord(x, y int) (cx, cy int) {
	xp, yp := x-y, x+y
	return xp >> chunkBits, yp >> chunkBits
}
