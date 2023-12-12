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

// Package day11 solves AoC 2023 day 11.
package day11

import (
	"cmp"
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 11, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	galaxies := findGalaxies(lines)
	g1 := expandSpace(galaxies, 2)
	p1 := totalDistance(g1)
	g2 := expandSpace(galaxies, 1000000)
	p2 := totalDistance(g2)
	return glue.Ints(p1, p2), nil
}

func findGalaxies(lines []string) (galaxies []util.P) {
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, util.P{x, y})
			}
		}
	}
	return galaxies
}

func expandSpace(galaxies []util.P, factor int) (expanded []util.P) {
	expanded = make([]util.P, len(galaxies))

	minY := galaxies[0].Y
	oldY, newY := minY, minY
	for i, g := range galaxies {
		if g.Y > oldY {
			newY += factor*(g.Y-oldY-1) + 1
			oldY = g.Y
		}
		expanded[i] = util.P{X: g.X, Y: newY}
	}

	slices.SortFunc(expanded, func(a, b util.P) int { return cmp.Compare(a.X, b.X) })
	minX := expanded[0].X
	oldX, newX := minX, minX
	for i, g := range expanded {
		if g.X > oldX {
			newX += factor*(g.X-oldX-1) + 1
			oldX = g.X
		}
		expanded[i].X = newX
	}

	return expanded
}

func totalDistance(galaxies []util.P) (sum int) {
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += util.DistM(galaxies[i], galaxies[j])
		}
	}
	return sum
}
