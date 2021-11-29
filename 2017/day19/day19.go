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

// Package day19 solves AoC 2017 day 19.
package day19

import (
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 19, glue.LevelSolver{
		Solver: solve,
		Empty:  ' ',
	})
}

func solve(l *util.Level) ([]string, error) {
	p1, p2 := tracePath(l)
	return []string{p1, strconv.Itoa(p2)}, nil
}

func tracePath(l *util.Level) (string, int) {
	var letters []byte
	steps := 0

	at := findStart(l)
	d := util.P{0, 1} // going down
loop:
	for {
		steps++
		if c := l.At(at.X, at.Y); c >= 'A' && c <= 'Z' {
			letters = append(letters, c)
		}
		if n := at.Add(d); l.At(n.X, n.Y) != ' ' {
			at = n
			continue loop
		}
		for _, turn := range [2]util.P{{d.Y, -d.X}, {-d.Y, d.X}} {
			if n := at.Add(turn); l.At(n.X, n.Y) != ' ' {
				at, d = n, turn
				continue loop
			}
		}
		break
	}

	return string(letters), steps
}

func findStart(l *util.Level) util.P {
	minP, maxP := l.Bounds()
	for x := minP.X; x <= maxP.X; x++ {
		if l.At(x, minP.Y) != ' ' {
			return util.P{x, minP.Y}
		}
	}
	panic("invalid level: no start found")
}
