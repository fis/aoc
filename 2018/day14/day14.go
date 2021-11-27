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

// Package day14 solves AoC 2018 day 14.
package day14

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 14, glue.IntSolver(solve))
}

func solve(input []int) ([]int, error) {
	if len(input) != 1 {
		return nil, fmt.Errorf("expected one int, got %d", len(input))
	}
	part1 := asInt(evolve(input[0], 10))
	part2 := find(asBytes(input[0]))
	return []int{part1, part2}, nil
}

func evolve(skip, keep int) []byte {
	scoreboard := []byte{3, 7}
	ep, fp := 0, 1
	for len(scoreboard) < skip+keep {
		er, fr := scoreboard[ep], scoreboard[fp]
		nr := er + fr
		if nr >= 10 {
			scoreboard = append(scoreboard, nr/10, nr%10)
		} else {
			scoreboard = append(scoreboard, nr)
		}
		ep = (ep + 1 + int(er)) % len(scoreboard)
		fp = (fp + 1 + int(fr)) % len(scoreboard)
	}
	return scoreboard[skip : skip+keep]
}

func find(scores []byte) int {
	scoreboard := []byte{3, 7}
	ep, fp := 0, 1
	for {
		er, fr := scoreboard[ep], scoreboard[fp]
		nr := er + fr
		i := len(scoreboard)
		if nr >= 10 {
			scoreboard = append(scoreboard, nr/10, nr%10)
		} else {
			scoreboard = append(scoreboard, nr)
		}
		for i < len(scoreboard) {
			if j := i - len(scores); j >= 0 && bytesEqual(scoreboard[j:i], scores) {
				return j
			}
			i++
		}
		ep = (ep + 1 + int(er)) % len(scoreboard)
		fp = (fp + 1 + int(fr)) % len(scoreboard)
	}
}

func asInt(scores []byte) (out int) {
	for _, s := range scores {
		out = 10*out + int(s)
	}
	return out
}

func asBytes(in int) (out []byte) {
	var (
		buf [16]byte
		at  int
	)
	for at = 15; at >= 0 && in > 0; at-- {
		buf[at] = byte(in % 10)
		in /= 10
	}
	return buf[at+1 : 16]
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
