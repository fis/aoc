// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package day10 solves AoC 2023 day 10.
package day10

import (
	"math"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 10, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(l *util.Level) ([]string, error) {
	lp := traceLoop(l)
	p1 := len(lp.tiles) / 2
	p2 := enclosedArea(l, lp)
	return glue.Ints(p1, p2), nil
}

type loop struct {
	start util.P
	dir   util.P
	tiles []util.P
}

func traceLoop(l *util.Level) loop {
	x, y, ok := l.Find('S')
	if !ok {
		panic("nowhere to start from")
	}
	start := util.P{x, y}
	l.Set(x, y, 'X')

	dx, dy := 0, 0
	for _, d := range [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		if connected(l.At(x+d[0], y+d[1]), d[0], d[1]) {
			dx, dy = d[0], d[1]
			break
		}
	}
	if dx == 0 && dy == 0 {
		panic("nowhere to go")
	}
	dir := util.P{dx, dy}

	tiles := []util.P{{x, y}}
	for {
		x, y = x+dx, y+dy
		t := l.At(x, y)
		if t == 'X' {
			break
		}
		l.Set(x, y, 'X')
		tiles = append(tiles, util.P{x, y})
		dx, dy = turn(t, dx, dy)
	}

	return loop{start: start, dir: dir, tiles: tiles}
}

func enclosedArea(l *util.Level, lp loop) (area int) {
	N := len(lp.tiles)
	p, q := findInside(lp.tiles)
	for i, tp := range lp.tiles {
		pp, np := lp.tiles[(i+N-1)%N], lp.tiles[(i+1)%N]
		dx, dy := tp.X-pp.X, tp.Y-pp.Y
		area += fillEmpty(l, tp.X+p*dy, tp.Y+q*dx)
		dx, dy = np.X-tp.X, np.Y-tp.Y
		area += fillEmpty(l, tp.X+p*dy, tp.Y+q*dx)
	}
	return area
}

func findInside(tiles []util.P) (p, q int) {
	// See: https://en.wikipedia.org/wiki/Curve_orientation#Orientation_of_a_simple_polygon

	N := len(tiles)

	ci, cx, cy := -1, math.MaxInt, math.MaxInt // index of one corner of the convex hull
	for i, tp := range tiles {
		if tp.Y < cy || tp.Y == cy && tp.X < cx {
			ci, cx, cy = i, tp.X, tp.Y
		}
	}

	a, b, c := tiles[(ci+N-1)%N], tiles[ci], tiles[(ci+1)%N]
	det := (b.X-a.X)*(c.Y-a.Y) - (c.X-a.X)*(b.Y-a.Y)
	if det < 0 {
		return 1, -1
	} else {
		return -1, 1
	}
}

func fillEmpty(l *util.Level, x, y int) (area int) {
	if l.At(x, y) == 'X' {
		return 0
	}
	l.Set(x, y, 'X')
	area = 1
	q := []util.P{{x, y}}
	for len(q) > 0 {
		p := q[len(q)-1]
		q = q[:len(q)-1]
		for _, n := range p.Neigh() {
			if l.At(n.X, n.Y) != 'X' {
				l.Set(n.X, n.Y, 'X')
				area++
				q = append(q, n)
			}
		}
	}
	return area
}

func connected(target byte, dx, dy int) bool {
	switch target {
	case '|':
		return dy != 0
	case '-':
		return dx != 0
	case 'L':
		return dx == -1 || dy == 1
	case 'J':
		return dx == 1 || dy == 1
	case '7':
		return dx == 1 || dy == -1
	case 'F':
		return dx == -1 || dy == -1
	}
	return false
}

func turn(target byte, dx, dy int) (nx, ny int) {
	switch target {
	case '|', '-':
		return dx, dy
	case 'L', '7':
		if dx != 0 {
			return turnRight(dx, dy)
		} else {
			return turnLeft(dx, dy)
		}
	case 'J', 'F':
		if dx != 0 {
			return turnLeft(dx, dy)
		} else {
			return turnRight(dx, dy)
		}
	}
	panic("you've done a bad turn")
}

func turnLeft(dx, dy int) (nx, ny int) {
	return dy, -dx
}

func turnRight(dx, dy int) (nx, ny int) {
	return -dy, dx
}
