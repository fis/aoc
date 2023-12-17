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

// Package day17 solves AoC 2023 day 17.
package day17

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 17, glue.FixedLevelSolver(solve))
}

func solve(l *util.FixedLevel) ([]string, error) {
	p1 := heatLoss(l)
	p2 := ultraHeatLoss(l)
	return glue.Ints(p1, p2), nil
}

func heatLoss(l *util.FixedLevel) int {
	w, h := l.W, l.H

	type state struct {
		pos    util.P // position
		di, dc int    // index and count of most recent moving direction
	}
	q := util.NewBucketQ[state](16)
	seen := make([]byte, w*h*4)

	q.Push(0, state{pos: util.P{0, 0}, di: -1, dc: 0})
	for q.Len() > 0 {
		pd, p := q.Pop()
		for di, n := range p.pos.Neigh() {
			if di == p.di^1 || (di == p.di && p.dc == 3) {
				continue // no U turns, no more than 3 steps in same direction
			}
			if n.X < 0 || n.X >= w || n.Y < 0 || n.Y >= h {
				continue // out of bounds
			}
			nd := pd + int(l.At(n.X, n.Y)-'0')
			if n.X == w-1 && n.Y == h-1 {
				return nd
			}
			ndc := 1
			if di == p.di {
				ndc = p.dc + 1
			}
			si := (n.Y*w+n.X)*4 + di
			if sc := seen[si]; sc > 0 && sc <= byte(ndc) {
				continue // been there like this already
			}
			seen[si] = byte(ndc)
			q.Push(nd, state{pos: n, di: di, dc: ndc})
		}
	}

	return -1
}

func ultraHeatLoss(l *util.FixedLevel) int {
	w, h := l.W, l.H

	type state struct {
		pos    util.P // position
		di, dc int    // index and count of most recent moving direction
	}
	q := util.NewBucketQ[state](16)
	seen := make([]uint16, w*h*4)

	q.Push(0, state{pos: util.P{0, 0}, di: -1, dc: 10})
	for q.Len() > 0 {
		pd, p := q.Pop()
		for di, n := range p.pos.Neigh() {
			if di == p.di^1 || (di == p.di && p.dc == 10) {
				continue // no U turns, no more than 10 steps in same direction
			}
			if di&2 != p.di&2 && p.dc < 4 {
				continue // no turning before 4 steps
			}
			if n.X < 0 || n.X >= w || n.Y < 0 || n.Y >= h {
				continue // out of bounds
			}
			nd := pd + int(l.At(n.X, n.Y)-'0')
			ndc := 1
			if di == p.di {
				ndc = p.dc + 1
			}
			if n.X == w-1 && n.Y == h-1 && ndc >= 4 {
				return nd
			}
			si := (n.Y*w+n.X)*4 + di
			if seen[si]&(1<<ndc) != 0 {
				continue // been there like this already
			}
			seen[si] |= 1 << ndc
			q.Push(nd, state{pos: n, di: di, dc: ndc})
		}
	}

	return -1
}
