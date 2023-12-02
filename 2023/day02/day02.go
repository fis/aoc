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

// Package day02 solves AoC 2023 day 2.
package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 2, glue.LineSolver(glue.WithParser(parseGame, solve)))
}

func solve(games []game) ([]string, error) {
	p1 := findPossible(games, cubeCount{colorRed: 12, colorGreen: 13, colorBlue: 14})
	p2 := findPower(games)
	return glue.Ints(p1, p2), nil
}

func findPossible(games []game, cubes cubeCount) (idSum int) {
nextGame:
	for _, g := range games {
		for _, s := range g.samples {
			for c := 0; c < colorCount; c++ {
				if s[c] > cubes[c] {
					continue nextGame
				}
			}
		}
		idSum += g.id
	}
	return idSum
}

func findPower(games []game) int {
	return fn.SumF(games, func(g game) int {
		var minCubes cubeCount
		for c := 0; c < colorCount; c++ {
			minCubes[c] = fn.MaxF(g.samples, func(s cubeCount) int { return s[c] })
		}
		return minCubes[colorRed] * minCubes[colorGreen] * minCubes[colorBlue]
	})
}

type game struct {
	id      int
	samples []cubeCount
}

type cubeCount [3]int

const (
	colorRed   = 0
	colorGreen = 1
	colorBlue  = 2
	colorCount = 3
)

var colorNames = map[string]int{
	"red": colorRed, "green": colorGreen, "blue": colorBlue,
}

func parseGame(line string) (g game, err error) {
	prefix, body, ok := strings.Cut(line, ": ")
	if !ok {
		return game{}, fmt.Errorf("missing prefix/body separator: %s", line)
	}
	prefix, ok = strings.CutPrefix(prefix, "Game ")
	if !ok {
		return game{}, fmt.Errorf("missing game header: %s", line)
	}
	g.id, err = strconv.Atoi(prefix)
	if err != nil {
		return game{}, fmt.Errorf("bad game ID: %s: %w", line, err)
	}
	for _, sampleSpec := range strings.Split(body, "; ") {
		var s cubeCount
		for _, colorSpec := range strings.Split(sampleSpec, ", ") {
			parts := strings.Split(colorSpec, " ")
			if len(parts) != 2 {
				return game{}, fmt.Errorf("wrong number of parts in color specification %q: %d", colorSpec, len(parts))
			}
			colorId, ok := colorNames[parts[1]]
			if !ok {
				return game{}, fmt.Errorf("bad color name: %q", parts[1])
			}
			s[colorId], err = strconv.Atoi(parts[0])
			if err != nil {
				return game{}, fmt.Errorf("bad count in color specification %q: %w", colorSpec, err)
			}
		}
		g.samples = append(g.samples, s)
	}
	return g, err
}
