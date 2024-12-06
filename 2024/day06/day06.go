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

// Package day06 solves AoC 2024 day 6.
package day06

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 6, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	part1 := walkOff(level.Copy())
	part2 := countLoops(level)
	return glue.Ints(part1, part2), nil
}

func walkOff(level *util.FixedLevel) (positions int) {
	W, H := level.W, level.H
	x, y, _ := level.Find('^')
	dx, dy := 0, -1
	for {
		if level.At(x, y) != 'X' {
			positions++
		}
		level.Set(x, y, 'X')
		for {
			nx, ny := x+dx, y+dy
			if nx < 0 || nx >= W || ny < 0 || ny >= H {
				return positions
			}
			if level.At(nx, ny) != '#' {
				x, y = nx, ny
				break
			}
			dx, dy = -dy, dx
		}
	}
}

func countLoops(level *util.FixedLevel) (loops int) {
	W, H := level.W, level.H
	x, y, _ := level.Find('^')
	d, dx, dy := 0, 0, -1
	seen := util.MakeFixedBitmap1D((W * H) << 2)
	for {
		level.Set(x, y, 'X')
		for {
			seen.Set((y*W+x)<<2 | d)
			nx, ny := x+dx, y+dy
			if nx < 0 || nx >= W || ny < 0 || ny >= H {
				return loops
			}
			tx, ty := -dy, dx
			if nc := level.At(nx, ny); nc != '#' {
				if nc == '.' && wouldLoop(level, x, y, (d+1)&3, tx, ty, seen) {
					loops++ // would loop if (nx, ny) was blocked
				}
				x, y = nx, ny
				break
			}
			d, dx, dy = (d+1)&3, tx, ty
		}
	}
}

func wouldLoop(level *util.FixedLevel, x, y, d, dx, dy int, seen util.FixedBitmap1D) (loops bool) {
	W, H := level.W, level.H
	var backtrack []int
	ox, oy := x+dy, y-dx
	level.Set(ox, oy, '#')
	for {
		pos := (y*W+x)<<2 | d
		if seen.Get(pos) {
			loops = true
			break
		}
		seen.Set(pos)
		backtrack = append(backtrack, pos)
		nx, ny := x+dx, y+dy
		if nx < 0 || nx >= W || ny < 0 || ny >= H {
			break
		}
		if level.At(nx, ny) == '#' {
			d, dx, dy = (d+1)&3, -dy, dx
			continue
		}
		x, y = nx, ny
	}
	level.Set(ox, oy, '.')
	for _, pos := range backtrack {
		seen.Unset(pos)
	}
	return loops
}
