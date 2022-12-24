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

// Package day24 solves AoC 2022 day 24.
package day24

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2022, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	lvl := parseLevel(lines)
	start, end := util.P{0, -1}, util.P{lvl.w - 1, lvl.h}
	p1 := findPath(lvl, 0, start, end)
	p2a := findPath(lvl, p1, end, start)
	p2b := findPath(lvl, p1+p2a, start, end)
	return glue.Ints(p1, p1+p2a+p2b), nil
}

func findPath(lvl level, t0 int, from, to util.P) (steps int) {
	period := len(lvl.blizzards)
	seen := make([]util.FixedBitmap2D, period)
	for i := range lvl.blizzards {
		seen[i] = util.MakeFixedBitmap2D(lvl.w, lvl.h+2)
	}

	type state struct {
		at util.P
		t  int
	}
	q := []state{{at: from, t: t0}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, n := range p.at.Neigh() {
			if n == to {
				return p.t + 1 - t0
			}
			if n.X < 0 || n.X >= lvl.w || n.Y < 0 || n.Y >= lvl.h {
				continue
			}
			if seen[(p.t+1)%period].Get(n.X, n.Y+1) || lvl.blizzards[(p.t+1)%period].Get(n.X, n.Y+1) {
				continue
			}
			seen[(p.t+1)%period].Set(n.X, n.Y+1)
			q = append(q, state{at: n, t: p.t + 1})
		}
		if !seen[(p.t+1)%period].Get(p.at.X, p.at.Y+1) && !lvl.blizzards[(p.t+1)%period].Get(p.at.X, p.at.Y+1) {
			seen[(p.t+1)%period].Set(p.at.X, p.at.Y+1)
			q = append(q, state{at: p.at, t: p.t + 1})
		}
	}
	return -1 // can't get there from here
}

type level struct {
	w, h      int
	blizzards []util.FixedBitmap2D
}

func parseLevel(lines []string) (lvl level) {
	w, h := len(lines[0])-2, len(lines)-2

	bmaps := map[byte]util.FixedBitmap2D{
		'>': util.MakeFixedBitmap2D(w, h),
		'v': util.MakeFixedBitmap2D(w, h),
		'<': util.MakeFixedBitmap2D(w, h),
		'^': util.MakeFixedBitmap2D(w, h),
	}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if bmp, ok := bmaps[lines[y+1][x+1]]; ok {
				bmp.Set(x, y)
			}
		}
	}

	period := ix.LCM(w, h)
	blizzards := make([]util.FixedBitmap2D, period)
	yD, yU, bw := 0, 0, (w+63)>>6
	bR, bD, bL, bU := bmaps['>'], bmaps['v'], bmaps['<'], bmaps['^']
	for t := 0; t < period; t++ {
		blizzards[t] = util.MakeFixedBitmap2D(w, h+2)
		for y := 0; y < h; y++ {
			for x := 0; x < bw; x++ {
				blizzards[t][y+1][x] |= bR[y][x]
				blizzards[t][y+1][x] |= bD[(yD+y)%h][x]
				blizzards[t][y+1][x] |= bL[y][x]
				blizzards[t][y+1][x] |= bU[(yU+y)%h][x]
			}
		}
		bR.RotateR(w)
		bL.RotateL(w)
		yD, yU = (yD+h-1)%h, (yU+1)%h
	}

	return level{w: w, h: h, blizzards: blizzards}
}
