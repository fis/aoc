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

// Package day10 solves AoC 2024 day 10.
package day10

import (
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 10, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	part1, part2 := totalScore(level)
	return glue.Ints(part1, part2), nil
}

func totalScore(level *util.FixedLevel) (sumScore, sumRating int) {
	s := newScorer(level)
	for y := range level.H {
		for x := range level.W {
			if level.At(x, y) == '0' {
				reachable, rating := s.analyze(x, y)
				sumScore += reachable.count()
				sumRating += rating
			}
		}
	}
	return sumScore, sumRating
}

type scorer struct {
	W, H      int
	level     *util.FixedLevel
	peaks     []int
	reachable []peakSet
	rating    []int
}

func newScorer(level *util.FixedLevel) *scorer {
	W, H := level.W, level.H
	peaks := make([]int, W*H)
	peak := 0
	for y := range H {
		for x := range W {
			if level.At(x, y) == '9' {
				peaks[y*W+x] = peak
				peak++
			} else {
				peaks[y*W+x] = -1
			}
		}
	}
	return &scorer{
		W: W, H: H, level: level, peaks: peaks,
		reachable: make([]peakSet, W*H),
		rating:    make([]int, W*H),
	}
}

func (s *scorer) analyze(x, y int) (reachable peakSet, rating int) {
	idx := y*s.W + x
	if r := s.rating[idx]-1; r >= 0 {
		return s.reachable[idx], r
	}
	h := s.level.At(x, y)
	if h == '9' {
		reachable.set(s.peaks[idx])
		rating = 1
	} else {
		for _, n := range (util.P{x, y}).Neigh() {
			if n.X < 0 || n.X >= s.W || n.Y < 0 || n.Y >= s.H {
				continue
			}
			if s.level.At(n.X, n.Y) == h+1 {
				re, ra := s.analyze(n.X, n.Y)
				reachable.setAll(re)
				rating += ra
			}
		}
	}
	s.reachable[idx] = reachable
	s.rating[idx] = rating+1
	return reachable, rating
}

// peakSet is a special-purpose bitmap type that's just wide enough to have bits for every peak
type peakSet [3]uint64

func (ps *peakSet) set(i int) {
	ps[i>>6] |= 1 << (i&63)
}

func (ps *peakSet) setAll(other peakSet) {
	ps[0] |= other[0]
	ps[1] |= other[1]
	ps[2] |= other[2]
}

func (ps *peakSet) count() int {
	return bits.OnesCount64(ps[0]) + bits.OnesCount64(ps[1]) + bits.OnesCount64(ps[2])
}
