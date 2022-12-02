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

// Package day02 solves AoC 2022 day 2.
package day02

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	assumedGuide := parseGuideAssumed(lines)
	realGuide := parseGuideReal(lines)
	p1 := totalScoreShapes(assumedGuide)
	p2 := totalScoreRounds(realGuide)
	return glue.Ints(p1, p2), nil
}

func totalScoreShapes(guide [][2]shape) int {
	return fn.Sum(fn.Map(guide, scoreShapes))
}

func totalScoreRounds(guide []guideLine) int {
	return fn.Sum(fn.Map(guide, scoreRound))
}

func scoreShapes(shapes [2]shape) int {
	return shapes[1].score() + match(shapes[1], shapes[0]).score()
}

func scoreRound(round guideLine) int {
	return findShape(round.theirs, round.outcome).score() + round.outcome.score()
}

type shape int

const (
	rock shape = iota
	paper
	scissors
)

func (s shape) score() int {
	return int(s) + 1
}

type outcome int

const (
	loss outcome = iota
	draw
	win
)

func (o outcome) score() int {
	return 3 * int(o)
}

var outcomes = [3][3]outcome{
	rock:     {rock: draw, paper: loss, scissors: win},
	paper:    {rock: win, paper: draw, scissors: loss},
	scissors: {rock: loss, paper: win, scissors: draw},
}

func match(yours, theirs shape) outcome {
	return outcomes[yours][theirs]
}

var shapeMap = [3][3]shape{
	rock:     {loss: scissors, draw: rock, win: paper},
	paper:    {loss: rock, draw: paper, win: scissors},
	scissors: {loss: paper, draw: scissors, win: rock},
}

func findShape(theirs shape, want outcome) shape {
	return shapeMap[theirs][want]
}

func parseGuideAssumed(lines []string) [][2]shape {
	guide := make([][2]shape, len(lines))
	for i, line := range lines {
		guide[i] = [2]shape{shape(line[0] - 'A'), shape(line[2] - 'X')}
	}
	return guide
}

type guideLine struct {
	theirs  shape
	outcome outcome
}

func parseGuideReal(lines []string) []guideLine {
	guide := make([]guideLine, len(lines))
	for i, line := range lines {
		guide[i] = guideLine{theirs: shape(line[0] - 'A'), outcome: outcome(line[2] - 'X')}
	}
	return guide
}
