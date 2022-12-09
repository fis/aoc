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

// Package day12 solves AoC 2020 day 12.
package day12

import (
	"fmt"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2020, 12, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	actions := parseInput(lines)

	t := newTurtle()
	t.move(actions)
	p1 := t.distance()

	s := newShip()
	s.move(actions)
	p2 := s.distance()

	return glue.Ints(p1, p2), nil
}

type action struct {
	command byte
	arg     int
}

func parseInput(lines []string) (out []action) {
	for _, line := range lines {
		c := action{}
		if _, err := fmt.Sscanf(line, "%c%d", &c.command, &c.arg); err == nil {
			out = append(out, c)
		}
	}
	return out
}

type turtle struct {
	pos util.P
	dir util.P
}

func newTurtle() turtle {
	return turtle{pos: util.P{0, 0}, dir: util.P{1, 0}}
}

func (t *turtle) move(actions []action) {
	for _, a := range actions {
		switch a.command {
		case 'N':
			t.pos.Y -= a.arg
		case 'S':
			t.pos.Y += a.arg
		case 'E':
			t.pos.X += a.arg
		case 'W':
			t.pos.X -= a.arg
		case 'L':
			t.dir = rotate(t.dir, a.arg)
		case 'R':
			t.dir = rotate(t.dir, 360-a.arg)
		case 'F':
			t.pos.X += a.arg * t.dir.X
			t.pos.Y += a.arg * t.dir.Y
		}
	}
}

func (t turtle) distance() int {
	return ix.Abs(t.pos.X) + ix.Abs(t.pos.Y)
}

type ship struct {
	pos util.P
	wp  util.P
}

func newShip() ship {
	return ship{pos: util.P{0, 0}, wp: util.P{10, -1}}
}

func (s *ship) move(actions []action) {
	for _, a := range actions {
		switch a.command {
		case 'N':
			s.wp.Y -= a.arg
		case 'S':
			s.wp.Y += a.arg
		case 'E':
			s.wp.X += a.arg
		case 'W':
			s.wp.X -= a.arg
		case 'L':
			s.wp = rotate(s.wp, a.arg)
		case 'R':
			s.wp = rotate(s.wp, 360-a.arg)
		case 'F':
			s.pos.X += a.arg * s.wp.X
			s.pos.Y += a.arg * s.wp.Y
		}
	}
}

func (s ship) distance() int {
	return ix.Abs(s.pos.X) + ix.Abs(s.pos.Y)
}

func rotate(p util.P, deg int) util.P {
	switch deg {
	case 90:
		return util.P{p.Y, -p.X}
	case 180:
		return util.P{-p.X, -p.Y}
	case 270:
		return util.P{-p.Y, p.X}
	}
	return p
}
