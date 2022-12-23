// Copyright 2022 Google LLC
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

// Package day23 solves AoC 2022 day 23.
package day23

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := simulate(lines)
	return glue.Ints(p1, p2), nil
}

func simulate(lines []string) (space10, lastRound int) {
	type proposal struct {
		elf int
		to  util.P
	}
	var (
		elves       []util.P
		elfMap      util.Bitmap2D
		props       []proposal
		propMap     util.Bitmap2D
		propBlocked util.Bitmap2D
	)

	for y, line := range lines {
		for x, b := range []byte(line) {
			if b == '#' {
				elves = append(elves, util.P{x, y})
				elfMap.Set(x, y)
			}
		}
	}
	props = make([]proposal, 0, len(elves))

	directions := [4]struct {
		d    util.P
		test uint32
	}{
		{util.P{0, -1}, 0b111_000_000},
		{util.P{0, 1}, 0b000_000_111},
		{util.P{-1, 0}, 0b100_100_100},
		{util.P{1, 0}, 0b001_001_001},
	}

	round := 0
	for {
		for i, elf := range elves {
			neigh := uint32(elfMap.GetR(elf.X-1, elf.Y-1, elf.X+1, elf.Y+1))
			if neigh == 0b000_010_000 {
				continue
			}
			for j := 0; j < 4; j++ {
				dir := directions[(round+j)&3]
				if neigh&dir.test == 0 {
					to := elf.Add(dir.d)
					if propMap.GetSet(to.X, to.Y) {
						propBlocked.Set(to.X, to.Y)
					} else {
						props = append(props, proposal{elf: i, to: to})
					}
					break
				}
			}
		}

		moved := false
		for _, prop := range props {
			if !propBlocked.GetClear(prop.to.X, prop.to.Y) {
				moved = true
				elfMap.Clear(elves[prop.elf].X, elves[prop.elf].Y)
				elfMap.Set(prop.to.X, prop.to.Y)
				elves[prop.elf] = prop.to
			}
			propMap.Clear(prop.to.X, prop.to.Y)
		}
		if !moved {
			break
		}
		props = props[:0]

		round++
		if round == 10 {
			getX := func(p util.P) int { return p.X }
			getY := func(p util.P) int { return p.Y }
			minX, maxX := fn.MinF(elves, getX), fn.MaxF(elves, getX)
			minY, maxY := fn.MinF(elves, getY), fn.MaxF(elves, getY)
			space10 = (maxX-minX+1)*(maxY-minY+1) - len(elves)
		}
	}

	return space10, round + 1
}
