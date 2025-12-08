// Copyright 2025 Google LLC
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

// Package day07 solves AoC 2025 day 7.
package day07

import (
	"maps"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fi"
)

func init() {
	glue.RegisterSolver(2025, 7, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	p1, p2 := countPaths(level)
	return glue.Ints(p1, p2), nil
}

func countSplits(level *util.FixedLevel) (splits int) {
	startX, startY, ok := level.Find('S')
	if !ok || startY != 0 {
		return -1
	}
	beams, newBeams := make(map[int]struct{}), make(map[int]struct{})
	beams[startX] = struct{}{}
	// N.B. assumes odd rows are empty
	for y := 2; y < level.H; y += 2 {
		clear(newBeams)
		for x := range beams {
			switch level.At(x, y) {
			case '^':
				splits++
				newBeams[x-1] = struct{}{}
				newBeams[x+1] = struct{}{}
			default:
				newBeams[x] = struct{}{}
			}
		}
		beams, newBeams = newBeams, beams
	}
	return splits
}

func countPaths(level *util.FixedLevel) (splits, paths int) {
	startX, startY, ok := level.Find('S')
	if !ok || startY != 0 {
		return -1, -1
	}
	beams, newBeams := make(map[int]int), make(map[int]int)
	beams[startX] = 1
	// N.B. assumes odd rows are empty
	for y := 2; y < level.H; y += 2 {
		clear(newBeams)
		for x, p := range beams {
			switch level.At(x, y) {
			case '^':
				splits++
				newBeams[x-1] += p
				newBeams[x+1] += p
			default:
				newBeams[x] += p
			}
		}
		beams, newBeams = newBeams, beams
	}
	return splits, fi.Sum(maps.Values(beams))
}
