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

// Package day17 solves AoC 2021 day 17.
package day17

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 17, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one input line, got %d", len(lines))
	}
	var Tmin, Tmax util.P
	Tmin.X, _ = strconv.Atoi(lines[0][0])
	Tmax.X, _ = strconv.Atoi(lines[0][1])
	Tmin.Y, _ = strconv.Atoi(lines[0][2])
	Tmax.Y, _ = strconv.Atoi(lines[0][3])
	p1, p2 := findShots(Tmin, Tmax)
	return glue.Ints(p1, p2), nil
}

func findShots(Pmin, Pmax util.P) (height, count int) {
	vymin, vymax := Pmin.Y, -Pmin.Y-1
	vxmin, vxmax := quadGe0(1, 1, -2*Pmin.X, 0, Pmin.X), Pmax.X
	for vy0true := vymin; vy0true <= vymax; vy0true++ {
		vy0, ty0 := vy0true, 0
		if vy0 >= 0 {
			vy0, ty0 = -(vy0 + 1), 2*vy0+1
		}
		tymin := quadGe0(1, -(2*vy0+1), 2*Pmax.Y, 0, Pmax.Y/vy0+1) + ty0
		tymax := quadLe0(1, -(2*vy0+1), 2*Pmin.Y, 0, Pmin.Y/vy0+1) + ty0
		for vx0 := vxmin; vx0 <= vxmax; vx0++ {
			txmin := quadGe0(-1, 2*vx0+1, -2*Pmin.X, 0, vx0)
			txmax := math.MaxInt
			if vx0*(vx0+1)/2 > Pmax.X {
				txmax = quadLe0(-1, 2*vx0+1, -2*Pmax.X, 0, vx0)
			}
			tmin, tmax := max(tymin, txmin, 0), min(tymax, txmax)
			if tmin <= tmax {
				if vy0true > 0 {
					if h := vy0true * (vy0true + 1) / 2; h > height {
						height = h
					}
				}
				count++
			}
		}
	}
	return height, count
}

func quadGe0(a, b, c, x0, x1 int) (x int) {
	y0 := a*x0*x0 + b*x0 + c
	y1 := a*x1*x1 + b*x1 + c
	if y0 >= 0 || y1 < 0 {
		err := fmt.Sprintf("quadGe0: bad initial conditions: (%d)x^2 + (%d)x + (%c): %d -> %d, %d -> %d", a, b, c, x0, y0, x1, y1)
		panic(err)
	}
	for x1-x0 > 1 {
		x = x0 + (x1-x0)/2
		y := a*x*x + b*x + c
		if y < 0 {
			x0 = x
		} else {
			x1 = x
		}
	}
	return x1
}

func quadLe0(a, b, c, x0, x1 int) (x int) {
	y0 := a*x0*x0 + b*x0 + c
	y1 := a*x1*x1 + b*x1 + c
	if y0 > 0 || y1 <= 0 {
		err := fmt.Sprintf("quadLe0: bad initial conditions: (%d)x^2 + (%d)x + (%c): %d -> %d, %d -> %d", a, b, c, x0, y0, x1, y1)
		panic(err)
	}
	for x1-x0 > 1 {
		x = x0 + (x1-x0)/2
		y := a*x*x + b*x + c
		if y <= 0 {
			x0 = x
		} else {
			x1 = x
		}
	}
	return x0
}
