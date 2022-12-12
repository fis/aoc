// Copyright 2022 Google LLC
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

// Package day12 solves AoC 2022 day 12.
package day12

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	m, start, end := readMap(lines)
	p1 := shortestPath(m, start, end)
	p2 := scenicPath(m, end)
	return glue.Ints(p1, p2), nil
}

func shortestPath(m [][]byte, start, end util.P) int {
	w, h := len(m[0]), len(m)
	type path struct {
		p util.P
		d int
	}
	q := []path{{start, 0}}
	seen := util.MakeFixedBitmap2D(w, h)
	seen.Set(0, 0)
	for len(q) > 0 {
		at := q[0]
		q = q[1:]
		for _, n := range at.p.Neigh() {
			if n.X >= 0 && n.X < w && n.Y >= 0 && n.Y < h && m[n.Y][n.X] <= m[at.p.Y][at.p.X]+1 && !seen.Get(n.X, n.Y) {
				nd := at.d + 1
				if n == end {
					return nd
				}
				seen.Set(n.X, n.Y)
				q = append(q, path{n, nd})
			}
		}
	}
	return -1
}

func scenicPath(m [][]byte, end util.P) int {
	w, h := len(m[0]), len(m)
	type path struct {
		p util.P
		d int
	}
	q := []path{{end, 0}}
	seen := util.MakeFixedBitmap2D(w, h)
	seen.Set(0, 0)
	for len(q) > 0 {
		at := q[0]
		q = q[1:]
		for _, n := range at.p.Neigh() {
			if n.X >= 0 && n.X < w && n.Y >= 0 && n.Y < h && m[n.Y][n.X] >= m[at.p.Y][at.p.X]-1 && !seen.Get(n.X, n.Y) {
				nd := at.d + 1
				if m[n.Y][n.X] == 'a' {
					return nd
				}
				seen.Set(n.X, n.Y)
				q = append(q, path{n, nd})
			}
		}
	}
	return -1
}

func readMap(lines []string) (m [][]byte, start, end util.P) {
	m = fn.Map(lines, func(line string) []byte { return []byte(line) })
	for y, row := range m {
		for x, b := range row {
			switch b {
			case 'S':
				m[y][x] = 'a'
				start = util.P{x, y}
			case 'E':
				m[y][x] = 'z'
				end = util.P{x, y}
			}
		}
	}
	return m, start, end
}
