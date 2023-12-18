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

// Package day18 solves AoC 2023 day 18.
package day18

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 18, glue.RegexpSolver{
		Solver: glue.WithParser(parseDigPlan, solve),
		Regexp: `([UDLR]) (\d+) \(#([0-9a-f]{6})\)`,
	})
}

func solve(plan []instruction) ([]string, error) {
	p1 := polygonArea(traceLoop(plan, instruction.small))
	p2 := polygonArea(traceLoop(plan, instruction.big))
	return glue.Ints(p1, p2), nil
}

func traceLoop(plan []instruction, f func(instruction) (direction, int)) (verts cycle) {
	startI, startX, startY := 0, 0, 0
	x, y := 0, 0
	cubes := make(cycle, len(plan))
	for i, inst := range plan {
		cubes[i] = util.P{x, y}
		if y < startY || (y == startY && x < startX) {
			startI, startX, startY = i, x, y
		}
		dir, dist := f(inst)
		dx, dy := dir.delta()
		x, y = x+dist*dx, y+dist*dy
	}
	if x != 0 || y != 0 {
		panic("not a closed loop")
	}
	verts = make(cycle, len(cubes))
	di := 0
	if a, b, c := cubes.at(startI-1), cubes.at(startI), cubes.at(startI+1); a.X == b.X && a.Y > b.Y && c.X > b.X && c.Y == b.Y {
		di = 1
	} else if a.X > b.X && a.Y == b.Y && c.X == b.X && c.Y > b.Y {
		di = -1
	} else {
		panic("top-left corner is not right")
	}
	for i := 0; i < len(cubes); i++ {
		bi := startI + i*di
		a, b, c := cubes.at(bi-di), cubes.at(bi), cubes.at(bi+di)
		dx, dy := 0, 0
		if c.Y > b.Y || a.Y < b.Y {
			dx = 1
		}
		if c.X < b.X || a.X > b.X {
			dy = 1
		}
		verts[i] = util.P{b.X + dx, b.Y + dy}
	}
	return verts
}

func polygonArea(verts cycle) (area int) {
	for i := 0; i < len(verts); i += 2 {
		a, b := verts.at(i), verts.at(i+1)
		area += a.Y * (a.X - b.X)
	}
	return area
}

type cycle []util.P

func (c cycle) at(i int) util.P { return c[(i+len(c))%len(c)] }

type direction byte

const (
	dirR direction = 0
	dirD direction = 1
	dirL direction = 2
	dirU direction = 3
)

var dirDelta = [4]util.P{
	dirR: {1, 0},
	dirD: {0, 1},
	dirL: {-1, 0},
	dirU: {0, -1},
}

func (d direction) delta() (dx, dy int) { return dirDelta[d].X, dirDelta[d].Y }

var dirLetters = map[byte]direction{
	'R': dirR,
	'D': dirD,
	'L': dirL,
	'U': dirU,
}

type instruction struct {
	smallDir  direction
	smallDist int
	bigDir    direction
	bigDist   int
}

func (i instruction) small() (direction, int) { return i.smallDir, i.smallDist }
func (i instruction) big() (direction, int)   { return i.bigDir, i.bigDist }

func parseDigPlan(parts []string) (instruction, error) {
	smallDir, ok := dirLetters[parts[0][0]]
	if !ok {
		return instruction{}, fmt.Errorf("bad direction: %q", parts[0])
	}
	smallDist, err := strconv.Atoi(parts[1])
	if err != nil {
		return instruction{}, err
	}
	hex, err := strconv.ParseInt(parts[2], 16, 64)
	if err != nil {
		return instruction{}, err
	}
	hexDir := hex & 0xf
	if hexDir >= 4 {
		return instruction{}, fmt.Errorf("bad hex direction: %x", hexDir)
	}
	return instruction{
		smallDir: smallDir, smallDist: smallDist,
		bigDir: direction(hexDir), bigDist: int(hex >> 4),
	}, nil
}
