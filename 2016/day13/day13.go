// Copyright 2023 Google LLC
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

// Package day13 solves AoC 2016 day 13.
package day13

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2016, 13, glue.IntSolver(solve))
}

func solve(nums []int) ([]string, error) {
	if len(nums) != 1 {
		return nil, fmt.Errorf("expected 1 number, got %d", len(nums))
	}
	key := nums[0]
	p1 := shortestPath(key, util.P{1, 1}, util.P{31, 39})
	p2 := coverage(key, util.P{1, 1}, 50)
	return glue.Ints(p1, p2), nil
}

func shortestPath(key int, from, to util.P) int {
	W, H := 2*max(from.X, to.X), 2*max(from.Y, to.Y)
	bitmap := fn.MapRange(1, H, func(int) []bool { return make([]bool, W) })
	bitmap[from.Y][from.X] = true

	type path struct {
		p util.P
		d int
	}
	q := []path{{p: from, d: 0}}
	for len(q) > 0 {
		at := q[0]
		q = q[1:]
		for _, n := range at.p.Neigh() {
			if n == to {
				return at.d + 1
			}
			if n.X < 0 || n.X >= W || n.Y < 0 || n.Y >= H || bitmap[n.Y][n.X] {
				continue
			}
			bitmap[n.Y][n.X] = true
			open := bits.OnesCount(uint(n.X*n.X+3*n.X+2*n.X*n.Y+n.Y+n.Y*n.Y+key))&1 == 0
			if open {
				q = append(q, path{p: n, d: at.d + 1})
			}
		}
	}
	return -1
}

func coverage(key int, from util.P, maxD int) (visited int) {
	W, H := from.X+maxD+1, from.Y+maxD+1
	bitmap := fn.MapRange(1, H, func(int) []bool { return make([]bool, W) })
	bitmap[from.Y][from.X] = true
	visited = 1

	type path struct {
		p util.P
		d int
	}
	q := []path{{p: from, d: 0}}
	for len(q) > 0 {
		at := q[0]
		q = q[1:]
		if at.d >= maxD {
			break
		}
		for _, n := range at.p.Neigh() {
			if n.X < 0 || n.X >= W || n.Y < 0 || n.Y >= H || bitmap[n.Y][n.X] {
				continue
			}
			bitmap[n.Y][n.X] = true
			open := bits.OnesCount(uint(n.X*n.X+3*n.X+2*n.X*n.Y+n.Y+n.Y*n.Y+key))&1 == 0
			if open {
				visited++
				q = append(q, path{p: n, d: at.d + 1})
			}
		}
	}
	return visited
}
