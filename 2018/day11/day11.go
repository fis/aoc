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

// Package day11 solves AoC 2018 day 11.
package day11

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2018, 11, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	if len(input) != 1 {
		return nil, fmt.Errorf("expected a single number, got %d", len(input))
	}

	g := makeGrid(input[0])
	x, y, _ := sweetSpot(g)
	x2, y2, s2, _ := sweetSquare(g)

	return []string{
		fmt.Sprintf("%d,%d", x, y),
		fmt.Sprintf("%d,%d,%d", x2, y2, s2),
	}, nil
}

type grid [300][300]int8

func makeGrid(s int) *grid {
	var g grid
	for y := 0; y < 300; y++ {
		for x := 0; x < 300; x++ {
			g[y][x] = int8(powerLevel(x+1, y+1, s))
		}
	}
	return &g
}

func sweetSpot(g *grid) (bestX, bestY, totalPower int) {
	totalPower = -3*3*5 - 1
	for y := 0; y+2 < 300; y++ {
		pow := int(g[y][0] + g[y][1] + g[y][2] + g[y+1][0] + g[y+1][1] + g[y+1][2] + g[y+2][0] + g[y+2][1] + g[y+2][2])
		if pow > totalPower {
			bestX, bestY, totalPower = 0, y, pow
		}
		for x := 0; x+3 < 300; x++ {
			pow -= int(g[y][x] + g[y+1][x] + g[y+2][x])
			pow += int(g[y][x+3] + g[y+1][x+3] + g[y+2][x+3])
			if pow > totalPower {
				bestX, bestY, totalPower = x+1, y, pow
			}
		}
	}
	return bestX + 1, bestY + 1, totalPower
}

func sweetSquare(g *grid) (bestX, bestY, bestSize, totalPower int) {
	var cum [300][300]int
	for x, c := 0, 0; x < 300; x++ {
		c += int(g[0][x])
		cum[0][x] = c
	}
	for y := 1; y < 300; y++ {
		for x, c := 0, 0; x < 300; x++ {
			c += int(g[y][x])
			cum[y][x] = cum[y-1][x] + c
		}
	}
	totalPower = -300*300*5 - 1
	for s := 2; s < 300; s++ {
		for y := 0; y+s < 300; y++ {
			for x := 0; x+s < 300; x++ {
				pow := cum[y+s-1][x+s-1]
				if x > 0 {
					pow -= cum[y+s-1][x-1]
				}
				if y > 0 {
					pow -= cum[y-1][x+s-1]
				}
				if x > 0 && y > 0 {
					pow += cum[y-1][x-1]
				}
				if pow > totalPower {
					bestX, bestY, bestSize, totalPower = x, y, s, pow
				}
			}
		}
	}
	return bestX + 1, bestY + 1, bestSize, totalPower
}

func powerLevel(x, y, s int) int {
	rack := x + 10
	pow := rack * y
	pow += s
	pow *= rack
	pow = (pow / 100) % 10
	return pow - 5
}
