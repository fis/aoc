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

// Package day18 solves AoC 2018 day 18.
package day18

import (
	"hash/fnv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 18, glue.LevelSolver{Solver: solve, Empty: ' '})
}

func solve(level *util.Level) ([]string, error) {
	part1 := value(evolve(level.Copy(), 10))

	target := 1000000000
	level, at, period := findCycle(level)
	at += (target - at) / period * period
	if at < target {
		level = evolve(level, target-at)
	}
	part2 := value(level)

	return glue.Ints(part1, part2), nil
}

func evolve(level *util.Level, generations int) *util.Level {
	next := level.Copy()
	for g := 0; g < generations; g++ {
		step(level, next)
		level, next = next, level
	}
	return level
}

func findCycle(level *util.Level) (out *util.Level, at, period int) {
	seen := map[uint64]int{}
	next := level.Copy()
	for at = 0; ; at++ {
		h := hashLevel(level)
		if prev, ok := seen[h]; ok {
			return level, at, at - prev
		}
		seen[h] = at
		step(level, next)
		level, next = next, level
	}
}

func step(in, out *util.Level) {
	min, max := in.Bounds()
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			nt, nl := 0, 0
			for _, n := range (util.P{x, y}).Neigh8() {
				switch in.At(n.X, n.Y) {
				case '|':
					nt++
				case '#':
					nl++
				}
			}
			v := in.At(x, y)
			switch {
			case v == '.' && nt >= 3:
				v = '|'
			case v == '|' && nl >= 3:
				v = '#'
			case v == '#' && (nt == 0 || nl == 0):
				v = '.'
			}
			out.Set(x, y, v)
		}
	}
}

func value(level *util.Level) int {
	nt, nl := 0, 0
	level.Range(func(x, y int, v byte) {
		switch v {
		case '|':
			nt++
		case '#':
			nl++
		}
	})
	return nt * nl
}

func hashLevel(level *util.Level) uint64 {
	h := fnv.New64a()
	level.Write(h)
	return h.Sum64()
}
