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

// Package day02 solves AoC 2016 day 2.
package day02

import (
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1 := decode(lines)
	p2 := decodeCross(lines)
	return []string{p1, p2}, nil
}

func decode(sheet []string) string {
	keypad := [3][3]byte{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}}
	x, y := 1, 1
	var code strings.Builder
	for _, steps := range sheet {
		for _, step := range steps {
			switch step {
			case 'L':
				x = max(x-1, 0)
			case 'R':
				x = min(x+1, 2)
			case 'U':
				y = max(y-1, 0)
			case 'D':
				y = min(y+1, 2)
			}
		}
		code.WriteByte(keypad[y][x])
	}
	return code.String()
}

func decodeCross(sheet []string) string {
	keypad := [5][5]byte{{0, 0, '1', 0, 0}, {0, '2', '3', '4', 0}, {'5', '6', '7', '8', '9'}, {0, 'A', 'B', 'C', 0}, {0, 0, 'D', 0, 0}}
	x, y := 0, 2
	var code strings.Builder
	for _, steps := range sheet {
		for _, step := range steps {
			nx, ny := x, y
			switch step {
			case 'L':
				nx = max(x-1, 0)
			case 'R':
				nx = min(x+1, 4)
			case 'U':
				ny = max(y-1, 0)
			case 'D':
				ny = min(y+1, 4)
			}
			if keypad[ny][nx] != 0 {
				x, y = nx, ny
			}
		}
		code.WriteByte(keypad[y][x])
	}
	return code.String()
}
