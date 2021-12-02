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

// Package day02 solves AoC 2021 day 2.
package day02

import (
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 2, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(forward|down|up) (\d+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	moves := parseInput(lines)
	p1 := applyMoves(util.P{0, 0}, moves)
	p2 := applyMoves2(util.P{0, 0}, moves)
	return glue.Ints(p1.X*p1.Y, p2.X*p2.Y), nil
}

func applyMoves(from util.P, moves []util.P) util.P {
	for _, move := range moves {
		from = from.Add(move)
	}
	return from
}

func applyMoves2(from util.P, moves []util.P) util.P {
	aim := 0
	for _, move := range moves {
		if move.Y == 0 {
			from.X += move.X
			from.Y += aim * move.X
		} else {
			aim += move.Y
		}
	}
	return from
}

func parseInput(lines [][]string) (moves []util.P) {
	commands := map[string]util.P{
		"forward": {1, 0},
		"down":    {0, 1},
		"up":      {0, -1},
	}
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1])
		moves = append(moves, commands[line[0]].Scale(n))
	}
	return moves
}
