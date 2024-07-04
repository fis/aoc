// Copyright 2019 Google LLC
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

// Package day10 solves AoC 2019 day 10.
package day10

import (
	"math"
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2019, 10, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(level *util.Level) ([]string, error) {
	at, vis := findBest(level)
	nth := findNth(at, 200, level)
	return glue.Ints(vis, nth.X*100+nth.Y), nil
}

func findBest(level *util.Level) (util.P, int) {
	bestX, bestY, best := 0, 0, -1
	for from := range level.All() {
		visible := findVisible(from.X, from.Y, level)
		if len(visible) > best {
			bestX, bestY, best = from.X, from.Y, len(visible)
		}
	}
	return util.P{bestX, bestY}, best
}

func findNth(from util.P, n int, level *util.Level) util.P {
	for {
		vis := findVisible(from.X, from.Y, level)
		if len(vis) == 0 {
			panic("out of asteroids :(")
		}
		if n <= len(vis) {
			sort.Slice(vis, func(i, j int) bool { return angle(from, vis[i]) < angle(from, vis[j]) })
			return vis[n-1]
		}
		for _, p := range vis {
			level.Set(p.X, p.Y, '.')
		}
		n -= len(vis)
	}
}

func findVisible(fromX, fromY int, level *util.Level) []util.P {
	var visible []util.P
	for to := range level.All() {
		dx, dy := to.X-fromX, to.Y-fromY
		if dx == 0 && dy == 0 {
			continue
		}
		d := ix.GCD(ix.Abs(dx), ix.Abs(dy))
		dx, dy = dx/d, dy/d
		tx, ty := fromX+dx, fromY+dy
		for (tx != to.X || ty != to.Y) && level.At(tx, ty) != '#' {
			tx, ty = tx+dx, ty+dy
		}
		if tx == to.X && ty == to.Y {
			visible = append(visible, util.P{to.X, to.Y})
		}
	}
	return visible
}

func angle(from, to util.P) float64 {
	dx, dy := float64(to.X-from.X), float64(to.Y-from.Y)
	return math.Mod(math.Atan2(dx, -dy)+2*math.Pi, 2*math.Pi)
}
