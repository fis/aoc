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

// Package day05 solves AoC 2018 day 5.
package day05

import (
	"fmt"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2018, 5, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expecting only 1 line of input, got %d", len(lines))
	}
	p1 := reducedLen(lines[0])
	p2 := improve(lines[0])
	return []int{p1, p2}, nil
}

func reducedLen(polymer string) int {
	buf := reduce([]byte(polymer))
	return len(buf)
}

func reduce(polymer []byte) []byte {
	for {
		var i, o int
		for i < len(polymer) {
			if i+1 < len(polymer) && polymer[i] == unitPairs[polymer[i+1]] {
				i += 2
			} else {
				if o < i {
					polymer[o] = polymer[i]
				}
				i++
				o++
			}
		}
		if o < i {
			polymer = polymer[:o]
		} else {
			return polymer
		}
	}
}

func improve(polymer string) int {
	min := len(polymer)
	for i := byte(0); i < 26; i++ {
		buf := reduce(prune([]byte(polymer), 'a'+i, 'A'+i))
		if len(buf) < min {
			min = len(buf)
		}
	}
	return min
}

func prune(polymer []byte, u1, u2 byte) []byte {
	var i, o int
	for i < len(polymer) {
		if polymer[i] != u1 && polymer[i] != u2 {
			if o < i {
				polymer[o] = polymer[i]
			}
			o++
		}
		i++
	}
	return polymer[:o]
}

var unitPairs [256]byte

func init() {
	for i := byte(0); i < 26; i++ {
		unitPairs['a'+i] = 'A' + i
		unitPairs['A'+i] = 'a' + i
	}
}
