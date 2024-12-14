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

// Package day14 solves AoC 2024 day 14.
package day14

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2024, 14, glue.IntSolver(solve))
}

func solve(nums []int) ([]string, error) {
	robots := parseRobots(nums)
	part1 := safetyFactor(robots, 101, 103, 100)
	part2 := findTree(robots, 101, 103)
	return glue.Ints(part1, part2), nil
}

func safetyFactor(robots []robot, w, h, t int) int {
	tl, tr, bl, br := 0, 0, 0, 0
	for _, r := range robots {
		p := r.predict(w, h, t)
		if p.Y < h/2 {
			if p.X < w/2 {
				tl++
			} else if p.X > w/2 {
				tr++
			}
		} else if p.Y > h/2 {
			if p.X < w/2 {
				bl++
			} else if p.X > w/2 {
				br++
			}
		}
	}
	return tl * tr * bl * br
}

func findTree(robots []robot, w, h int) int {
	marginH, marginV := make([]int, w), make([]int, h)
	varH, varV := 2*ix.CeilDiv(len(robots)*(w-1), w*w), 2*ix.CeilDiv(len(robots)*(h-1), h*h)
	tH, tV := 0, 0
	for t := 1; tH == 0 || tV == 0; t++ {
		for i := range robots {
			robots[i].step(w, h)
		}
		if tH == 0 {
			for x := range w {
				marginH[x] = 0
			}
			for _, r := range robots {
				marginH[r.p.X]++
			}
			if svar(marginH) > varH {
				tH = t
			}
		}
		if tV == 0 {
			for y := range h {
				marginV[y] = 0
			}
			for _, r := range robots {
				marginV[r.p.Y]++
			}
			if svar(marginV) > varV {
				tV = t
			}
		}
	}
	for t := tH; ; t += w {
		if t%h == tV {
			return t
		}
	}
}

func svar(data []int) int {
	mean := 0
	for _, x := range data {
		mean += x
	}
	mean /= len(data)
	svar := 0
	for _, x := range data {
		svar += (x - mean) * (x - mean)
	}
	return svar/len(data) - 1
}

type robot struct{ p, v util.P }

func (r *robot) step(w, h int) {
	r.p = r.p.Add(r.v)
	if r.p.X < 0 {
		r.p.X += w
	} else if r.p.X >= w {
		r.p.X -= w
	}
	if r.p.Y < 0 {
		r.p.Y += h
	} else if r.p.Y >= h {
		r.p.Y -= h
	}
}

func (r *robot) predict(w, h, t int) util.P {
	p := r.p.Add(r.v.Scale(t))
	p.X, p.Y = (p.X%w+w)%w, (p.Y%h+h)%h
	return p
}

func parseRobots(nums []int) []robot {
	robots := make([]robot, len(nums)>>2)
	for i := range robots {
		p, v := util.P{nums[i<<2], nums[i<<2|1]}, util.P{nums[i<<2|2], nums[i<<2|3]}
		robots[i] = robot{p, v}
	}
	return robots
}
