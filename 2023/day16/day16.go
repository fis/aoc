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

// Package day16 solves AoC 2023 day 16.
package day16

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 16, glue.FixedLevelSolver(solve))
}

func solve(l *util.FixedLevel) ([]string, error) {
	p1 := countEnergized(l, 0, 0, 1, 0, 0)
	p2 := maxEnergized(l, p1)
	return glue.Ints(p1, p2), nil
}

func maxEnergized(l *util.FixedLevel, best int) int {
	w, h := l.W, l.H
	for y := 1; y < h; y++ {
		if e := countEnergized(l, 0, y, 1, 0, 0); e > best {
			best = e
		}
	}
	for x := 0; x < w; x++ {
		if e := countEnergized(l, x, 0, 0, 1, 1); e > best {
			best = e
		}
	}
	for y := 0; y < h; y++ {
		if e := countEnergized(l, w-1, y, -1, 0, 2); e > best {
			best = e
		}
	}
	for x := 0; x < w; x++ {
		if e := countEnergized(l, x, h-1, 0, -1, 3); e > best {
			best = e
		}
	}
	return best
}

func countEnergized(l *util.FixedLevel, x, y, dx, dy, di int) (count int) {
	w, h := l.W, l.H
	energized := make([]byte, w*h)

	type beam struct {
		x, y       int
		dx, dy, di int
	}
	beams := []beam{{x, y, dx, dy, di}}

	for len(beams) > 0 {
		b := beams[len(beams)-1]
		beams = beams[:len(beams)-1]
		x, y, dx, dy, di = b.x, b.y, b.dx, b.dy, b.di

		for x >= 0 && x < w && y >= 0 && y < h && energized[y*w+x]&byte(1<<di) == 0 {
			energized[y*w+x] |= byte(1 << di)

			switch l.At(x, y) {
			case '\\':
				dx, dy, di = dy, dx, di^1
			case '/':
				dx, dy, di = -dy, -dx, di^3
			case '|':
				if dx != 0 {
					beams = append(beams, beam{x: x, y: y - 1, dx: 0, dy: -1, di: 3})
					dx, dy, di = 0, 1, 1
				}
			case '-':
				if dy != 0 {
					beams = append(beams, beam{x: x - 1, y: y, dx: -1, dy: 0, di: 2})
					dx, dy, di = 1, 0, 0
				}
			}

			x, y = x+dx, y+dy
		}
	}

	for _, e := range energized {
		if e != 0 {
			count++
		}
	}

	return count
}
