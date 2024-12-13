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

// Package day12 solves AoC 2024 day 12.
package day12

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2024, 12, glue.FixedLevelSolver(solve))
}

func solve(level *util.FixedLevel) ([]string, error) {
	part1, part2 := fencingPrice(level)
	return glue.Ints(part1, part2), nil
}

func fencingPrice(level *util.FixedLevel) (regular, discounted int) {
	for y := range level.H {
		for x := range level.W {
			if !isUpper(level.At(x, y)) {
				continue
			}
			area, perimeter, sides := findRegion(level, x, y)
			regular += area * perimeter
			discounted += area * sides
		}
	}
	return regular, discounted
}

func findRegion(level *util.FixedLevel, x, y int) (area, perimeter, sides int) {
	area, perimeter, sides = 1, 0, 0
	plant := level.At(x, y)
	level.Set(x, y, lower(plant))
	for _, n := range (util.P{x, y}).Neigh() {
		if !level.InBounds(n.X, n.Y) || upper(level.At(n.X, n.Y)) != plant {
			perimeter++
			if firstEdge(level, x, y, n.X, n.Y, plant) {
				sides++
			}
		} else if level.At(n.X, n.Y) == plant {
			na, np, ns := findRegion(level, n.X, n.Y)
			area, perimeter, sides = area+na, perimeter+np, sides+ns
		}
	}
	return area, perimeter, sides
}

// firstEdge returns true if the edge (px,py) | (nx,ny) is the first edge of that particular side.
func firstEdge(level *util.FixedLevel, px, py, nx, ny int, plant byte) bool {
	dx, dy := nx-px, ny-py
	px2, py2, nx2, ny2 := px+dy, py-dx, nx+dy, ny-dx
	return !level.InBounds(px2, py2) || upper(level.At(px2, py2)) != plant || (level.InBounds(nx2, ny2) && upper(level.At(nx2, ny2)) == plant)
}

func lower(c byte) byte { return c | 0x20 }
func upper(c byte) byte { return c &^ 0x20 }
func isUpper(c byte) bool { return c & 0x20 == 0 }

