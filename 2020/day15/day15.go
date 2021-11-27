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

// Package day15 solves AoC 2020 day 15.
package day15

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 15, glue.IntSolver(solve))
}

func solve(numbers []int) ([]int, error) {
	p1 := simulate(numbers, 2020)
	p2 := simulate(numbers, 30000000)
	return []int{p1, p2}, nil
}

func simulate(initial []int, upTo int) (num int) {
	lastTurn := []int(nil)
	for turn, num := range initial[:len(initial)-1] {
		lastTurn = ensureSize(lastTurn, num)
		lastTurn[num] = turn + 1
	}
	num = initial[len(initial)-1]
	for turn := len(initial); turn < upTo; turn++ {
		var next int
		lastTurn = ensureSize(lastTurn, num)
		if last := lastTurn[num]; last == 0 {
			next = 0
		} else {
			next = turn - last
		}
		lastTurn[num], num = turn, next
	}
	return num
}

func ensureSize(arr []int, num int) []int {
	if num < len(arr) {
		return arr
	}
	return append(arr, make([]int, num-len(arr)+1)...)
}
