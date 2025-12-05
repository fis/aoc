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

// Package day04 solves AoC 2025 day 4.
package day04

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2025, 4, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	p1, p2 := removePaper(level)
	return glue.Ints(p1, p2), nil
}

func removePaper(level *util.FixedLevel) (firstRound, totalRemoved int) {
	firstRound = -1
	remove := make([]util.P, 0, 2048)
	for {
		remove = remove[:0]
		for p := range level.NonEmpty('.') {
			rolls := 0
			for _, n := range p.Neigh8() {
				if level.InBounds(n.X, n.Y) && level.AtP(n) == '@' {
					rolls++
				}
			}
			if rolls < 4 {
				remove = append(remove, p)
			}
		}
		if len(remove) == 0 {
			break
		}
		for _, p := range remove {
			level.Set(p.X, p.Y, '.')
		}
		if firstRound == -1 {
			firstRound = len(remove)
		}
		totalRemoved += len(remove)
	}
	return firstRound, totalRemoved
}
