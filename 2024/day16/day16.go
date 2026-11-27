// Package day16 solves AoC 2024 day 16.
package day16

import (
	"math"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 16, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	part1, part2 := bestPath(level)
	return glue.Ints(part1, part2), nil
}

func bestPath(level *util.FixedLevel) (score, area int) {
	type state struct {
		x, y      int
		d, dx, dy int
	}
	type link struct {
		x, y, d int
	}
	type path struct {
		s int
		p []link
	}

	sx, sy, _ := level.Find('S')

	fringe := util.NewBucketQ[state](1024)
	fringe.Push(0, state{x: sx, y: sy, d: 0, dx: 1, dy: 0})
	prev := make([][][4]path, level.H)
	for y := range prev {
		prev[y] = make([][4]path, level.W)
	}
	score = math.MaxInt
	for fringe.Len() > 0 {
		s, p := fringe.Pop()
		if s >= score {
			break
		}
		next := []struct {
			s, d, dx, dy int
		}{
			{s + 1, p.d, p.dx, p.dy},
			{s + 1001, (p.d + 1) & 3, -p.dy, p.dx},
			{s + 1001, (p.d - 1) & 3, p.dy, -p.dx},
		}
		for _, n := range next {
			nx, ny := p.x+n.dx, p.y+n.dy
			nc := level.At(nx, ny)
			if nc != '.' && nc != 'E' {
				continue
			}
			np := prev[ny][nx][n.d]
			if np.s != 0 && np.s < n.s {
				continue // been in this state cheaper
			}
			if np.s != 0 && np.s > n.s {
				prev[ny][nx][n.d].p = nil // reset: got a new best
			}
			prev[ny][nx][n.d].p = appendNew(prev[ny][nx][n.d].p, link{p.x, p.y, p.d})
			prev[ny][nx][n.d].s = n.s
			if nc == 'E' {
				if n.s < score {
					score = n.s
				}
				continue
			}
			fringe.Push(n.s, state{x: nx, y: ny, d: n.d, dx: n.dx, dy: n.dy})
		}
	}

	// TODO: instead of a []link, consider [4]bool
	ex, ey, _ := level.Find('E')
	paths := []link{{ex, ey, 0}, {ex, ey, 1}, {ex, ey, 2}, {ex, ey, 3}}
	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		level.Set(p.x, p.y, 'O')
		paths = append(paths, prev[p.y][p.x][p.d].p...)
		prev[p.y][p.x][p.d].p = nil
	}

	for y := range level.H {
		for x := range level.W {
			if level.At(x, y) == 'O' {
				area++
			}
		}
	}

	return score, area
}

func appendNew[T comparable](s []T, e T) []T {
	for _, v := range s {
		if v == e {
			return s
		}
	}
	return append(s, e)
}
