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

// Package day10 solves AoC 2022 day 10.
package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 10, glue.ParsableLineSolver[instruction]{
		Solver: solve,
		Parser: parseInstruction,
	})
}

func solve(prog []instruction) ([]string, error) {
	p1 := sigStrength(prog)
	img := render(prog)
	return append(glue.Ints(p1), img...), nil
}

func sigStrength(prog []instruction) (strength int) {
	const (
		sampleCycle = 19
		sampleMod   = 40
	)
	cycle, x := 0, 1
	for _, inst := range prog {
		c := cycle % sampleMod
		if c <= sampleCycle && c+inst.cycles > sampleCycle {
			strength += (cycle - c + sampleCycle + 1) * x
		}
		cycle, x = cycle+inst.cycles, x+inst.addend
	}
	return strength
}

func render(prog []instruction) []string {
	const (
		W = 40
		H = 6
	)
	var screen [H][W]byte
	cycle, x := 0, 1
	for _, inst := range prog {
		for c := 0; c < inst.cycles; c++ {
			px, py := (cycle+c)%W, (cycle+c)/W
			screen[py][px] = fn.If[byte](px >= x-1 && px <= x+1, '#', ' ')
		}
		cycle, x = cycle+inst.cycles, x+inst.addend
	}
	return fn.Map(screen[:], func(row [W]byte) string { return string(row[:]) })
}

type instruction struct {
	cycles int
	addend int
}

func parseInstruction(line string) (instruction, error) {
	if line == "noop" {
		return instruction{cycles: 1, addend: 0}, nil
	}
	if strings.HasPrefix(line, "addx ") {
		addend, err := strconv.Atoi(line[5:])
		if err != nil {
			return instruction{}, fmt.Errorf("invalid addend: %q: %w", line[5:], err)
		}
		return instruction{cycles: 2, addend: addend}, nil
	}
	return instruction{}, fmt.Errorf("invalid instruction: %q", line)
}
