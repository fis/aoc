// Copyright 2019 Google LLC
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

// Package day19 solves AoC 2019 day 19.
package day19

import (
	"github.com/fis/aoc-go/intcode"
	"github.com/fis/aoc-go/util"
)

func init() {
	util.RegisterSolver(19, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	probe := prober(prog)
	return []int64{
		part1(50, probe),
		part2(50, 100, probe),
	}, nil
}

func part1(size int, probe func(x, y int) bool) int64 {
	var count int64
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			c := " "
			if probe(x, y) {
				count++
				c = "#"
			}
			util.Diag(c)
		}
		util.Diag("\n")
	}
	return count
}

func part2(start, size int, probe func(x, y int) bool) int64 {
	left := 0
	for !probe(left, start) {
		left++
	}
	right := left
	for probe(right+1, start) {
		right++
	}

	history := make([]beam, size)
	history[start%size] = beam{left, right}

	for y := start + 1; ; /* ever */ y++ {
		for !probe(left, y) {
			left++
		}
		for probe(right+1, y) {
			right++
		}
		history[y%size] = beam{left, right}
		if y < start+size {
			continue
		}
		prev := history[(y-size+1)%size]
		if left >= prev.left && left+size-1 <= prev.right {
			bx, by := left, y-size+1
			return int64(10000*bx + by)
		}
	}
}

func prober(prog []int64) func(x, y int) bool {
	return func(x, y int) bool {
		out, _ := intcode.Run(prog, []int64{int64(x), int64(y)})
		return out[0] != 0
	}
}

type beam struct {
	left  int
	right int
}
