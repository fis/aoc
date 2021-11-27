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

// Package day25 solves AoC 2018 day 25.
package day25

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 25, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	cs := constellationSet{}
	cs.addAll(lines)
	return []int{cs.size()}, nil
}

type constellationSet struct {
	points [][]p4
}

func (cs constellationSet) size() int {
	return len(cs.points)
}

func (cs *constellationSet) add(p p4) {
	var joinable []int
	for ci, c := range cs.points {
		for _, cp := range c {
			if distM(p, cp) <= 3 {
				joinable = append(joinable, ci)
				break
			}
		}
	}

	if len(joinable) == 0 {
		cs.points = append(cs.points, []p4{p})
		return
	}

	di := joinable[0]
	cs.points[di] = append(cs.points[di], p)
	for i := len(joinable) - 1; i >= 1; i-- {
		si := joinable[i]
		cs.points[di] = append(cs.points[di], cs.points[si]...)
		if si != len(cs.points)-1 {
			cs.points[si] = cs.points[len(cs.points)-1]
		}
		cs.points = cs.points[:len(cs.points)-1]
	}
}

func (cs *constellationSet) addAll(lines []string) {
	for _, line := range lines {
		cs.add(readP4(line))
	}
}

type p4 struct {
	x, y, z, t int
}

func readP4(s string) (p p4) {
	fmt.Sscanf(s, "%d,%d,%d,%d", &p.x, &p.y, &p.z, &p.t)
	return p
}

func distM(a, b p4) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z) + abs(a.t-b.t)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
