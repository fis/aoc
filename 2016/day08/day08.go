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

// Package day08 solves AoC 2016 day 8.
package day08

import (
	"math/bits"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 8, glue.RegexpSolver{Solver: solve, Regexp: inputRegexp})
}

const inputRegexp = `rect (\d+)x(\d+)|rotate \w+ ([xy])=(\d+) by (\d+)`

func solve(input [][]string) ([]string, error) {
	var s screen
	for _, op := range input {
		var a, b int
		if op[2] == "" {
			a, _ = strconv.Atoi(op[0])
			b, _ = strconv.Atoi(op[1])
		} else {
			a, _ = strconv.Atoi(op[3])
			b, _ = strconv.Atoi(op[4])
		}
		switch op[2] {
		case "":
			s.rect(a, b)
		case "y":
			s.rotRow(a, b)
		case "x":
			s.rotCol(a, b)
		}
	}
	return append(glue.Ints(s.countLit()), s.print()...), nil
}

const (
	W          = 50
	H          = 6
	screenMask = (1 << W) - 1
)

type screen [H]uint64

func (s *screen) rect(a, b int) {
	mask := (uint64(1) << a) - 1
	for y := 0; y < b; y++ {
		s[y] |= mask
	}
}

func (s *screen) rotRow(a, b int) {
	s[a] = ((s[a] << b) | (s[a] >> (W - b))) & screenMask
}

func (s *screen) rotCol(a, b int) {
	os := *s
	mask := uint64(1) << a
	for y := 0; y < H; y++ {
		s[y] &^= mask
		s[y] |= os[(y+H-b)%H] & mask
	}
}

func (s *screen) countLit() (sum int) {
	for _, row := range *s {
		sum += bits.OnesCount64(row)
	}
	return sum
}

func (s *screen) print() (rows []string) {
	var row [W]byte
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if (s[y] & (uint64(1) << x)) != 0 {
				row[x] = '#'
			} else {
				row[x] = ' '
			}
		}
		rows = append(rows, string(row[:]))
	}
	return rows
}
