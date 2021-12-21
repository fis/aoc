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

// Package day21 solves AoC 2021 day 21.
package day21

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 21, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^Player ([12]) starting position: (\d+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	if len(lines) != 2 || lines[0][0] != "1" || lines[1][0] != "2" {
		return nil, fmt.Errorf("expected two players, 1 and 2")
	}
	start1, _ := strconv.Atoi(lines[0][1])
	start2, _ := strconv.Atoi(lines[1][1])
	p1 := part1(start1, start2)
	p2 := part2(start1, start2)
	return glue.Ints(p1, p2), nil
}

func part1(start1, start2 int) int {
	var (
		pos        = [2]uint32{uint32(start1) - 1, uint32(start2) - 1}
		score      [2]uint32
		die        uint32
		turn       uint32 = 1
		totalRolls uint32
	)
	for score[turn] < 1000 {
		totalRolls += 3
		turn = 1 - turn
		rolls := die + (die+1)%100 + (die+2)%100 + 3
		die = (die + 3) % 100
		pos[turn] = (pos[turn] + rolls) % 10
		score[turn] += pos[turn] + 1
	}
	return int(score[1-turn] * totalRolls)
}

func part2(start1, start2 int) int {
	dieRolls := []struct {
		sum   uint32
		count int
	}{
		{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1},
	}

	var countWins func(needA, needB, posA, posB uint32) (winA, winB int)
	var knownWins [22][22][10][10][2]int
	countWins = func(needA, needB, posA, posB uint32) (winA, winB int) {
		if kw := knownWins[needA][needB][posA][posB]; kw[0] > 0 {
			return kw[0], kw[1]
		}
		for _, roll := range dieRolls {
			rposA := (posA + roll.sum) % 10
			if 1+rposA >= needA {
				winA += roll.count
			} else {
				rneedA := needA - (1 + rposA)
				rwinB, rwinA := countWins(needB, rneedA, posB, rposA)
				winA += roll.count * rwinA
				winB += roll.count * rwinB
			}
		}
		knownWins[needA][needB][posA][posB] = [2]int{winA, winB}
		return winA, winB
	}

	wins1, wins2 := countWins(21, 21, uint32(start1-1), uint32(start2-1))
	if wins2 > wins1 {
		return wins2
	} else {
		return wins1
	}
}
