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

// Package day03 solves AoC 2017 day 3.
package day03

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 3, glue.IntSolver(solve))
}

func solve(input []int) ([]int, error) {
	if len(input) != 1 {
		return nil, fmt.Errorf("expecting only 1 integer of input, got %d", len(input))
	}
	p1 := part1(input[0])
	p2 := part2(input[0])
	return []int{p1, p2}, nil
}

func part1(square int) int {
	if square == 1 {
		return 0
	}
	square--
	d := sqrt(square)
	if d%2 == 0 {
		d--
	}
	square -= d * d
	r := (d + 1) / 2
	return r + abs(square%(2*r)-(r-1))
}

func part2(limit int) int {
	grid := map[util.P]int{{0, 0}: 1}
	r := 1
	for {
		p := util.P{X: r, Y: -r}
		for _, d := range []util.P{{0, 1}, {-1, 0}, {0, -1}, {1, 0}} {
			for i := 0; i < 2*r; i++ {
				p = p.Add(d)
				v := 0
				for _, n := range p.Neigh8() {
					v += grid[n]
				}
				if v > limit {
					return v
				}
				grid[p] = v
			}
		}
		r++
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sqrt(y int) int {
	if y < 0 {
		panic("sqrt(neg)")
	} else if y <= 1 {
		return y
	}
	x0 := y / 2
	x1 := (x0 + y/x0) / 2
	for x1 < x0 {
		x0 = x1
		x1 = (x0 + y/x0) / 2
	}
	return x0
}
