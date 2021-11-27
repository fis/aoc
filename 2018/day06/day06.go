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

// Package day06 solves AoC 2018 day 6.
package day06

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	points := parsePoints(lines)
	max := maxArea(points)
	safe := safeArea(points, 10000)
	return []int{max, safe}, nil
}

func parsePoints(lines []string) (points []util.P) {
	for _, line := range lines {
		var p util.P
		if _, err := fmt.Sscanf(line, "%d, %d", &p.X, &p.Y); err == nil {
			points = append(points, p)
		}
	}
	return points
}

func maxArea(points []util.P) int {
	min, max := util.Bounds(points)

	infinite := make([]bool, len(points))
	for x := min.X - 1; x <= max.X+1; x++ {
		if i := closestIndex(util.P{x, min.Y - 1}, points); i >= 0 {
			infinite[i] = true
		}
		if i := closestIndex(util.P{x, max.Y + 1}, points); i >= 0 {
			infinite[i] = true
		}
	}
	for y := min.Y; y <= max.Y; y++ {
		if i := closestIndex(util.P{min.X - 1, y}, points); i >= 0 {
			infinite[i] = true
		}
		if i := closestIndex(util.P{max.X + 1, y}, points); i >= 0 {
			infinite[i] = true
		}
	}

	areas := make([]int, len(points))
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if i := closestIndex(util.P{x, y}, points); i >= 0 && !infinite[i] {
				areas[i]++
			}
		}
	}

	maxA := areas[0]
	for _, a := range areas[1:] {
		if a > maxA {
			maxA = a
		}
	}
	return maxA
}

func safeArea(points []util.P, threshold int) (safe int) {
	min, max := util.Bounds(points)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			total := 0
			for _, p := range points {
				total += util.DistM(util.P{x, y}, p)
			}
			if total < threshold {
				safe++
			}
		}
	}
	return safe
}

func closestIndex(p util.P, points []util.P) int {
	minD, minI, minC := util.DistM(p, points[0]), 0, 1
	for i, p2 := range points[1:] {
		d := util.DistM(p, p2)
		if d < minD {
			minD, minI, minC = d, i+1, 1
		} else if d == minD {
			minC++
		}
	}
	if minC == 1 {
		return minI
	}
	return -1
}
