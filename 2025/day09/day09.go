// Copyright 2025 Google LLC
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

// Package day09 solves AoC 2025 day 9.
package day09

import (
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fi"
)

func init() {
	glue.RegisterSolver(2025, 9, glue.LineSolver(glue.WithParser(util.ParseP, solve)))
}

func solve(redTiles []util.P) ([]string, error) {
	p1 := largestRect(redTiles)
	p2 := largestValid(redTiles)
	return glue.Ints(p1, p2), nil
}

func largestRect(redTiles []util.P) (maxArea int) {
	for p1, p2 := range fi.Pairwise(redTiles) {
		x1, y1 := min(p1.X, p2.X), min(p1.Y, p2.Y)
		x2, y2 := max(p1.X, p2.X), max(p1.Y, p2.Y)
		area := (x2 - x1 + 1) * (y2 - y1 + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func largestValid(redTiles []util.P) (maxArea int) {
	N := len(redTiles)

	signedArea := 0
	for i, p1 := range redTiles {
		p2 := redTiles[(i+1)%N]
		signedArea += p1.X*p2.Y - p2.X*p1.Y
	}
	// corner adjustments expect CW winding, which corresponds to >0 signed area in this coordinate system
	if signedArea < 0 {
		slices.Reverse(redTiles)
	}

	corners := make([]util.P, N)
	for i, p := range redTiles {
		prev, next := redTiles[(i+N-1)%N], redTiles[(i+1)%N]
		/*
		   Adjust the points to be the actual corners of the enclosed region:

		   case: ↳  ↲  ↱  ↰  ⬎  ⬏  ⬐  ⬑
		   dx12:  0  0  0  0 +1 +1 -1 -1
		   dy12: +1 +1 -1 -1  0  0  0  0
		   dx23: +1 -1 +1 -1  0  0  0  0
		   dy23:  0  0  0  0 +1 -1 +1 -1
		     dx:  1  1  0  0  1  0  1  0
		     dy:  0  1  0  1  0  0  1  1
		*/
		c := p
		if prev.Y < p.Y || next.Y > p.Y {
			c.X++
		}
		if prev.X > p.X || next.X < p.X {
			c.Y++
		}
		corners[i] = c
	}

	type lineSeg struct{ u, v1, v2 int }
	vlines, hlines := make([]lineSeg, 0, N/2), make([]lineSeg, 0, N/2)
	for i, c1 := range corners {
		c2 := corners[(i+1)%N]
		if c1.X == c2.X {
			vlines = append(vlines, lineSeg{c1.X, min(c1.Y, c2.Y), max(c1.Y, c2.Y)})
		} else {
			hlines = append(hlines, lineSeg{c1.Y, min(c1.X, c2.X), max(c1.X, c2.X)})
		}
	}

nextPair:
	for p1, p2 := range fi.Pairwise(redTiles) {
		x1, y1 := min(p1.X, p2.X), min(p1.Y, p2.Y)
		x2, y2 := max(p1.X, p2.X), max(p1.Y, p2.Y)
		area := (x2 - x1 + 1) * (y2 - y1 + 1)
		if area > maxArea {
			for _, l := range vlines {
				if l.u > x1 && l.u <= x2 && l.v2 > y1 && l.v1 <= y2 {
					continue nextPair
				}
			}
			for _, l := range hlines {
				if l.v2 > x1 && l.v1 <= x2 && l.u > y1 && l.u <= y2 {
					continue nextPair
				}
			}
			maxArea = area
		}
	}
	return maxArea
}
