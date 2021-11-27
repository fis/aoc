// Copyright 2021 Google LLC
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

// Package day06 solves AoC 2017 day 6.
package day06

import "github.com/fis/aoc/glue"

func init() {
	glue.RegisterSolver(2017, 6, glue.IntSolver(solve))
}

func solve(input []int) ([]int, error) {
	banks := make([]byte, len(input))
	for i, n := range input {
		banks[i] = byte(n)
	}
	p1 := loopLength(banks)
	p2 := loopLength(banks)
	return []int{p1, p2}, nil
}

func loopLength(banks []byte) int {
	seen := map[string]struct{}{string(banks): {}}
	steps := 0
	for {
		balance(banks)
		steps++
		state := string(banks)
		if _, found := seen[state]; found {
			return steps
		}
		seen[state] = struct{}{}
	}
}

func balance(banks []byte) {
	N := byte(len(banks))
	src, blocks := argmax(banks)
	banks[src] = 0
	if blocks >= N {
		all := blocks / N
		blocks = blocks % N
		for i := range banks {
			banks[i] += all
		}
	}
	for i := byte(1); i <= blocks; i++ {
		banks[(byte(src)+i)%N]++
	}
}

func argmax(xs []byte) (maxI int, maxX byte) {
	maxI, maxX = -1, 0
	for i, x := range xs {
		if maxI == -1 || x > maxX {
			maxI, maxX = i, x
		}
	}
	return maxI, maxX
}
