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

// Package day11 solves AoC 2021 day 11.
package day11

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 11, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	g, err := newGrid(lines)
	if err != nil {
		return nil, err
	}
	p1 := g.simulate(100)
	p2 := 100 + g.simulateToSync()
	return glue.Ints(p1, p2), nil
}

const size = 10

type grid [size][size]byte

func newGrid(lines []string) (*grid, error) {
	if len(lines) != size {
		return nil, fmt.Errorf("expected %d lines, got %d", size, len(lines))
	}
	g := &grid{}
	for y, line := range lines {
		if len(line) != size {
			return nil, fmt.Errorf("line %d: expected %d columns, got %d", y, size, len(line))
		}
		for x := 0; x < size; x++ {
			c := line[x]
			if c < '0' || c > '9' {
				return nil, fmt.Errorf("line %d, column %d: expected '0'..'9', got '%c'", y, x, c)
			}
			g[y][x] = c - '0'
		}
	}
	return g, nil
}

func (g *grid) simulate(steps int) (flashes int) {
	for step := 0; step < steps; step++ {
		flashes += g.step()
	}
	return flashes
}

func (g *grid) simulateToSync() (steps int) {
	steps = 1
	for g.step() < size*size {
		steps++
	}
	return steps
}

func (g *grid) step() (flashes int) {
	var newFlashes bitset
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			c := g[y][x] + 1
			if c == 10 {
				flashes++
				newFlashes.set(x, y)
				c = 0
			}
			g[y][x] = c
		}
	}
	for {
		x, y, ok := newFlashes.pop()
		if !ok {
			break
		}
		for _, n := range (util.P{x, y}).Neigh8() {
			if n.X < 0 || n.X >= size || n.Y < 0 || n.Y >= size {
				continue
			}
			c := g[n.Y][n.X] + 1
			if c == 1 { // already flashed this step
				continue
			} else if c == 10 {
				flashes++
				newFlashes.set(n.X, n.Y)
				c = 0
			}
			g[n.Y][n.X] = c
		}
	}
	return flashes
}

type bitset [(size*size + 63) / 64]uint64

func (bs *bitset) set(x, y int) {
	i := y*size + x
	bs[i>>6] |= uint64(1) << (i & 0x3f)
}

func (bs *bitset) pop() (x, y int, ok bool) {
	for b, n := range bs {
		if n == 0 {
			continue
		}
		i := bits.TrailingZeros64(n)
		if i == 64 {
			continue
		}
		bs[b] &= ^(uint64(1) << i)
		i += b * 64
		return i % size, i / size, true
	}
	return 0, 0, false
}
