// Copyright 2020 Google LLC
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

// Package day22 solves AoC 2018 day 22.
package day22

import (
	"container/heap"
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 22, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	var depth, tX, tY int
	if len(lines) != 2 {
		return nil, fmt.Errorf("expected 2 lines, got %d", len(lines))
	} else if _, err := fmt.Sscanf(lines[0], "depth: %d", &depth); err != nil {
		return nil, fmt.Errorf("bad depth line: %s: %w", lines[0], err)
	} else if _, err := fmt.Sscanf(lines[1], "target: %d,%d", &tX, &tY); err != nil {
		return nil, fmt.Errorf("bad target line: %s: %w", lines[1], err)
	}
	cave := newCave(depth, tX, tY)

	part1 := cave.totalRisk()
	part2 := cave.shortestPath()

	return glue.Ints(part1, part2), nil
}

type cave struct {
	depth        int
	tX, tY       int
	erosionCache map[util.P]int
}

type rType int

const (
	rocky  rType = 0
	wet    rType = 1
	narrow rType = 2
)

func newCave(depth, tX, tY int) *cave {
	return &cave{depth: depth, tX: tX, tY: tY, erosionCache: make(map[util.P]int)}
}

func (c *cave) erosion(x, y int) int {
	switch {
	case x == 0 && y == 0, x == c.tX && y == c.tY:
		return c.depth % 20183
	case y == 0:
		return (16807*x + c.depth) % 20183
	case x == 0:
		return (48271*y + c.depth) % 20183
	default:
		if cached, ok := c.erosionCache[util.P{x, y}]; ok {
			return cached
		}
		e := (c.erosion(x-1, y)*c.erosion(x, y-1) + c.depth) % 20183
		c.erosionCache[util.P{x, y}] = e
		return e
	}
}

func (c *cave) rType(x, y int) rType {
	return rType(c.erosion(x, y) % 3)
}

func (c *cave) totalRisk() (risk int) {
	for y := 0; y <= c.tY; y++ {
		for x := 0; x <= c.tX; x++ {
			risk += int(c.rType(x, y))
		}
	}
	return risk
}

type tool int

const (
	noTool tool = iota
	climbingGear
	torch
)

func (t tool) compatible(r rType) bool {
	switch r {
	case rocky:
		return t == climbingGear || t == torch
	case wet:
		return t == noTool || t == climbingGear
	case narrow:
		return t == noTool || t == torch
	}
	return false
}

type state struct {
	at    util.P
	equip tool
}

type path struct {
	s  state
	d  int
	hd int
}

type pathq []path

func (c *cave) shortestPath() int {
	from := state{at: util.P{0, 0}, equip: torch}
	to := state{at: util.P{c.tX, c.tY}, equip: torch}
	dist := map[state]int{from: 0}
	fringe := pathq{{s: from, d: 0, hd: util.DistM(from.at, to.at)}}
	for len(fringe) > 0 {
		p := heap.Pop(&fringe).(path)
		if p.s == to {
			return p.d
		}
		if od := dist[p.s]; od < p.d {
			continue // path no longer relevant
		}
		for _, q := range p.s.at.Neigh() {
			if q.X < 0 || q.Y < 0 || !p.s.equip.compatible(c.rType(q.X, q.Y)) {
				continue
			}
			qp := path{s: state{at: q, equip: p.s.equip}, d: p.d + 1, hd: p.d + 1 + util.DistM(q, to.at)}
			if od, ok := dist[qp.s]; ok && od <= qp.d {
				continue
			}
			dist[qp.s] = qp.d
			heap.Push(&fringe, qp)
		}
		for t := noTool; t <= torch; t++ {
			if !t.compatible(c.rType(p.s.at.X, p.s.at.Y)) {
				continue
			}
			tp := path{s: state{at: p.s.at, equip: t}, d: p.d + 7, hd: p.d + 7 + util.DistM(p.s.at, to.at)}
			if od, ok := dist[tp.s]; ok && od <= tp.d {
				continue
			}
			dist[tp.s] = tp.d
			heap.Push(&fringe, tp)
		}
	}
	return -1 // can't get there from here
}

func (q pathq) Len() int           { return len(q) }
func (q pathq) Less(i, j int) bool { return q[i].hd < q[j].hd }
func (q pathq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *pathq) Push(x interface{}) {
	*q = append(*q, x.(path))
}

func (q *pathq) Pop() interface{} {
	old, n := *q, len(*q)
	path := old[n-1]
	*q = old[0 : n-1]
	return path
}
