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

// Package day17 solves AoC 2018 day 17.
package day17

import (
	"fmt"
	"math"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2018, 17, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	level, minY, maxY := readScan(lines)
	fill(level, util.P{500, 0}, make(map[util.P]struct{}))
	total, pooled := measureWater(level, minY, maxY)
	return glue.Ints(total, pooled), nil
}

func readScan(lines []string) (level *util.Level, minY, maxY int) {
	level = util.ParseLevelAt([]byte{'+'}, '.', util.P{500, 0})
	minY, maxY = math.MaxInt32, math.MinInt32
	for _, line := range lines {
		var a, bl, bh int
		if _, err := fmt.Sscanf(line, "x=%d, y=%d..%d", &a, &bl, &bh); err == nil {
			if bl < minY {
				minY = bl
			}
			if bh > maxY {
				maxY = bh
			}
			for b := bl; b <= bh; b++ {
				level.Set(a, b, '#')
			}
		} else if _, err := fmt.Sscanf(line, "y=%d, x=%d..%d", &a, &bl, &bh); err == nil {
			if a < minY {
				minY = a
			}
			if a > maxY {
				maxY = a
			}
			for b := bl; b <= bh; b++ {
				level.Set(b, a, '#')
			}
		}
	}
	return level, minY, maxY
}

func fill(level *util.Level, source util.P, filled map[util.P]struct{}) {
	p := source
	for level.InBounds(p.X, p.Y+1) && sandy(level.At(p.X, p.Y+1)) {
		p.Y++
		level.Set(p.X, p.Y, '|')
	}
	if !level.InBounds(p.X, p.Y+1) {
		return
	}
	for p.Y > source.Y {
		l, ls, r, rs := p.X, false, p.X, false
		for sandy(level.At(l-1, p.Y)) {
			l--
			level.Set(l, p.Y, '|')
			if sandy(level.At(l, p.Y+1)) {
				if _, ok := filled[util.P{l, p.Y}]; !ok {
					fill(level, util.P{l, p.Y}, filled)
					filled[util.P{l, p.Y}] = struct{}{}
				}
				if sandy(level.At(l, p.Y+1)) {
					ls = true
					break
				}
			}
		}
		for sandy(level.At(r+1, p.Y)) {
			r++
			level.Set(r, p.Y, '|')
			if sandy(level.At(r, p.Y+1)) {
				if _, ok := filled[util.P{r, p.Y}]; !ok {
					fill(level, util.P{r, p.Y}, filled)
					filled[util.P{r, p.Y}] = struct{}{}
				}
				if sandy(level.At(r, p.Y+1)) {
					rs = true
					break
				}
			}
		}
		if !ls && !rs {
			for x := l; x <= r; x++ {
				level.Set(x, p.Y, '~')
			}
			p.Y--
		} else {
			break
		}
	}
}

func measureWater(level *util.Level, minY, maxY int) (total int, pooled int) {
	level.Range(func(_, y int, b byte) {
		if y >= minY && y <= maxY && (b == '~' || b == '|') {
			total++
			if b == '~' {
				pooled++
			}
		}
	})
	return total, pooled
}

func sandy(p byte) bool {
	return p == '.' || p == '|'
}
