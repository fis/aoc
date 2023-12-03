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
	hwalls, vwalls := readWalls(lines)
	minP, maxP := util.P{500, math.MaxInt}, util.P{500, math.MinInt}
	for _, hw := range hwalls {
		minP.X = min(minP.X, hw.x1)
		minP.Y = min(minP.Y, hw.y)
		maxP.X = max(maxP.X, hw.x2)
		maxP.Y = max(maxP.Y, hw.y)
	}
	for _, vw := range vwalls {
		minP.X = min(minP.X, vw.x)
		minP.Y = min(minP.Y, vw.y1)
		maxP.X = max(maxP.X, vw.x)
		maxP.Y = max(maxP.Y, vw.y2)
	}
	level = util.EmptyLevel(util.P{minP.X, min(minP.Y, 0)}, util.P{maxP.X, max(maxP.Y, 0)}, '.')
	level.Set(500, 0, '+')
	for _, hw := range hwalls {
		for x := hw.x1; x <= hw.x2; x++ {
			level.Set(x, hw.y, '#')
		}
	}
	for _, vw := range vwalls {
		for y := vw.y1; y <= vw.y2; y++ {
			level.Set(vw.x, y, '#')
		}
	}
	return level, minP.Y, maxP.Y
}

func readWalls(lines []string) (hwalls []hwall, vwalls []vwall) {
	for _, line := range lines {
		var a, bl, bh int
		if _, err := fmt.Sscanf(line, "y=%d, x=%d..%d", &a, &bl, &bh); err == nil {
			hwalls = append(hwalls, hwall{y: a, x1: bl, x2: bh})
		} else if _, err := fmt.Sscanf(line, "x=%d, y=%d..%d", &a, &bl, &bh); err == nil {
			vwalls = append(vwalls, vwall{x: a, y1: bl, y2: bh})
		}
	}
	return hwalls, vwalls
}

type hwall struct{ y, x1, x2 int }
type vwall struct{ x, y1, y2 int }

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
