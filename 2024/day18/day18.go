// Copyright 2024 Google LLC
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

// Package day18 solves AoC 2024 day 18.
package day18

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 18, glue.LineSolver(glue.WithParser(util.ParseP, solve)))
}

func solve(incoming []util.P) ([]string, error) {
	level := simulate(71, incoming[:1024])
	path := findPath(level, nil)
	part1 := len(path) - 1
	part2 := findBreak(level, path, incoming[1024:])
	return []string{
		strconv.Itoa(part1),
		fmt.Sprintf("%d,%d", part2.X, part2.Y),
	}, nil
}

func simulate(size int, incoming []util.P) *util.FixedLevel {
	level := util.EmptyFixedLevel(size, size, '.')
	for _, p := range incoming {
		level.Set(p.X, p.Y, '#')
	}
	return level
}

func findPath(level *util.FixedLevel, path []util.P) []util.P {
	type link struct {
		prev util.P
		dist int
	}
	state := util.Make2D[link](level.W, level.H)
	state[0][0] = link{dist: 1}

	q := util.MakeQueue[util.P](128)
	q.Push(util.P{0, 0})

	tx, ty := level.W-1, level.H-1
loop:
	for !q.Empty() {
		p := q.Pop()
		d := state[p.Y][p.X].dist
		for _, n := range p.Neigh() {
			if !level.InBounds(n.X, n.Y) || level.At(n.X, n.Y) != '.' || state[n.Y][n.X].dist != 0 {
				continue
			}
			state[n.Y][n.X] = link{prev: p, dist: d + 1}
			if n.X == tx && n.Y == ty {
				break loop
			}
			q.Push(n)
		}
	}
	end := state[ty][tx]
	if end.dist == 0 {
		return nil
	}

	if end.dist > len(path) {
		path = append(path, make([]util.P, end.dist-len(path))...)
	} else {
		path = path[:end.dist]
	}
	for i, p := end.dist-1, (util.P{tx, ty}); i >= 0; i-- {
		path[i] = p
		p = state[p.Y][p.X].prev
	}
	return path
}

func findBreak(level *util.FixedLevel, path, incoming []util.P) util.P {
	markPath(level, path, 'O')
	for _, p := range incoming {
		if level.At(p.X, p.Y) != 'O' {
			level.Set(p.X, p.Y, '#')
			continue
		}
		markPath(level, path, '.')
		level.Set(p.X, p.Y, '#')
		path = findPath(level, path)
		if len(path) == 0 {
			return p // cut
		}
		markPath(level, path, 'O')
	}
	return util.P{-1, -1} // no cuts
}

func markPath(level *util.FixedLevel, path []util.P, mark byte) {
	for _, p := range path {
		level.Set(p.X, p.Y, mark)
	}
}
