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

// Package day13 solves AoC 2023 day 13.
package day13

import (
	"math/bits"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 13, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	ps := fn.Map(chunks, parsePattern)
	p1 := sumReflections(ps, findReflection)
	p2 := sumReflections(ps, findSmudged)
	return glue.Ints(p1, p2), nil
}

func sumReflections(ps []pattern, find func(p pattern) (rx, ry int)) (sum int) {
	for _, p := range ps {
		rx, ry := find(p)
		sum += rx + 100*ry
	}
	return sum
}

func findReflection(p pattern) (rx, ry int) {
findX:
	for x := 1; x < len(p.cols); x++ {
		s := min(x, len(p.cols)-x)
		for i := 0; i < s; i++ {
			if p.cols[x-1-i] != p.cols[x+i] {
				continue findX
			}
		}
		rx = x
		break findX
	}
	if rx > 0 {
		return rx, 0
	}
findY:
	for y := 1; y < len(p.lines); y++ {
		s := min(y, len(p.lines)-y)
		for i := 0; i < s; i++ {
			if p.lines[y-1-i] != p.lines[y+i] {
				continue findY
			}
		}
		ry = y
		break findY
	}
	return 0, ry
}

func findSmudged(p pattern) (rx, ry int) {
	for x := 1; x < len(p.cols); x++ {
		s := min(x, len(p.cols)-x)
		errors := 0
		for i := 0; i < s && errors <= 1; i++ {
			errors += bits.OnesCount32(p.cols[x-1-i] ^ p.cols[x+i])
		}
		if errors == 1 {
			return x, 0
		}
	}
	for y := 1; y < len(p.lines); y++ {
		s := min(y, len(p.lines)-y)
		errors := 0
		for i := 0; i < s && errors <= 1; i++ {
			errors += bits.OnesCount32(p.lines[y-1-i] ^ p.lines[y+i])
		}
		if errors == 1 {
			return 0, y
		}
	}
	return 0, 0
}

type pattern struct {
	lines []uint32
	cols  []uint32
}

func parsePattern(text string) pattern {
	if text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	w, h := strings.IndexByte(text, '\n'), 1+strings.Count(text, "\n")
	lines := make([]uint32, h)
	cols := make([]uint32, w)
	x, y := 0, 0
	for _, c := range text {
		if c == '\n' {
			x, y = 0, y+1
			continue
		}
		lines[y] <<= 1
		cols[x] <<= 1
		if c == '#' {
			lines[y] |= 1
			cols[x] |= 1
		}
		x++
	}
	return pattern{lines, cols}
}
