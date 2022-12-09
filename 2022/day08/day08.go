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

// Package day08 solves AoC 2022 day 8.
package day08

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 8, glue.LineSolver(solve))
}

func solve(forest []string) ([]string, error) {
	p1 := countVisible(forest)
	p2 := findBest(forest)
	return glue.Ints(p1, p2), nil
}

func countVisible(forest []string) (visible int) {
	w, h := len(forest[0]), len(forest)
	seen := fn.MapRange(1, h-1, func(int) []bool { return make([]bool, w-2) })
	visible = 2*w + 2*h - 4
	for y := 1; y < h-1; y++ {
		for x, z := 1, forest[y][0]; x < w-1 && z < '9'; x++ {
			if forest[y][x] > z {
				if !seen[y-1][x-1] {
					visible++
					seen[y-1][x-1] = true
				}
				z = forest[y][x]
			}
		}
		for x, z := w-2, forest[y][w-1]; x > 0 && z < '9'; x-- {
			if forest[y][x] > z {
				if !seen[y-1][x-1] {
					visible++
					seen[y-1][x-1] = true
				}
				z = forest[y][x]
			}
		}
	}
	for x := 1; x < w-1; x++ {
		for y, z := 1, forest[0][x]; y < h-1 && z < '9'; y++ {
			if forest[y][x] > z {
				if !seen[y-1][x-1] {
					visible++
					seen[y-1][x-1] = true
				}
				z = forest[y][x]
			}
		}
		for y, z := h-2, forest[h-1][x]; y > 0 && z < '9'; y-- {
			if forest[y][x] > z {
				if !seen[y-1][x-1] {
					visible++
					seen[y-1][x-1] = true
				}
				z = forest[y][x]
			}
		}
	}
	return visible
}

func findBest(forest []string) (bestScore int) {
	w, h := len(forest[0]), len(forest)
	scores := fn.MapRange(1, h-1, func(int) []int { return fn.MapRange(1, w-1, func(int) int { return 1 }) })
	for y := 1; y < h-1; y++ {
		var px [10]int
		for x := 1; x < w-1; x++ {
			scores[y-1][x-1] *= x - px[forest[y][x]-'0']
			for z := byte('0'); z <= forest[y][x]; z++ {
				px[z-'0'] = x
			}
		}
		px = [10]int{w - 1, w - 1, w - 1, w - 1, w - 1, w - 1, w - 1, w - 1, w - 1, w - 1}
		for x := w - 2; x > 0; x-- {
			scores[y-1][x-1] *= px[forest[y][x]-'0'] - x
			for z := byte('0'); z <= forest[y][x]; z++ {
				px[z-'0'] = x
			}
		}
	}
	for x := 1; x < w-1; x++ {
		var py [10]int
		for y := 1; y < h-1; y++ {
			scores[y-1][x-1] *= y - py[forest[y][x]-'0']
			for z := byte('0'); z <= forest[y][x]; z++ {
				py[z-'0'] = y
			}
		}
		py = [10]int{h - 1, h - 1, h - 1, h - 1, h - 1, h - 1, h - 1, h - 1, h - 1, h - 1}
		for y := h - 2; y > 0; y-- {
			scores[y-1][x-1] *= py[forest[y][x]-'0'] - y
			for z := byte('0'); z <= forest[y][x]; z++ {
				py[z-'0'] = y
			}
		}
	}
	return fn.Max(fn.Map(scores, func(s []int) int { return fn.Max(s) }))
}
