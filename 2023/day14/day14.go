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

// Package day14 solves AoC 2023 day 14.
package day14

import (
	"bytes"
	"hash/fnv"
	"io"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2023, 14, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	l := util.ParseFixedLevel(data)
	slideNorth(l)
	p1 := totalLoad(l)
	slideWest(l)
	slideSouth(l)
	slideEast(l)
	runCycles(l, 999999999)
	p2 := totalLoad(l)
	return glue.Ints(p1, p2), nil
}

var singleRoundRock = []byte{'O'}

func totalLoad(l *util.FixedLevel) (sum int) {
	h := l.H
	for y := 0; y < h; y++ {
		load := h - y
		sum += bytes.Count(l.Row(y), singleRoundRock) * load
	}
	return sum
}

func runCycles(l *util.FixedLevel, cycles int) {
	seen := make(map[uint64]int)
	seen[hashLevel(l)] = 0
	for cycle := 1; cycle <= cycles; cycle++ {
		slideCycle(l)
		hash := hashLevel(l)
		if prev, ok := seen[hash]; ok {
			n := cycle - prev
			left := (cycles - cycle) % n
			for i := 0; i < left; i++ {
				slideCycle(l)
			}
			return
		}
		seen[hash] = cycle
	}
}

func hashLevel(l *util.FixedLevel) uint64 {
	h := fnv.New64a()
	h.Write(l.Data)
	return h.Sum64()
}

func slideCycle(l *util.FixedLevel) {
	slideNorth(l)
	slideWest(l)
	slideSouth(l)
	slideEast(l)
}

func slideNorth(l *util.FixedLevel) {
	w, h := l.W, l.H
	for x := 0; x < w; x++ {
		for y, startY, rocks := h-1, h-1, 0; y >= -1; y-- {
			if y == -1 || l.At(x, y) == '#' {
				if startY > y+1 && rocks > 0 {
					y2 := y + 1
					for ; y2 <= y+rocks; y2++ {
						l.Set(x, y2, 'O')
					}
					for ; y2 <= startY; y2++ {
						l.Set(x, y2, '.')
					}
				}
				startY, rocks = y-1, 0
				continue
			}
			if l.At(x, y) == 'O' {
				rocks++
			}
		}
	}
}

func slideWest(l *util.FixedLevel) {
	w, h := l.W, l.H
	for y := 0; y < h; y++ {
		for x, startX, rocks := w-1, w-1, 0; x >= -1; x-- {
			if x == -1 || l.At(x, y) == '#' {
				if startX > x+1 && rocks > 0 {
					x2 := x + 1
					for ; x2 <= x+rocks; x2++ {
						l.Set(x2, y, 'O')
					}
					for ; x2 <= startX; x2++ {
						l.Set(x2, y, '.')
					}
				}
				startX, rocks = x-1, 0
				continue
			}
			if l.At(x, y) == 'O' {
				rocks++
			}
		}
	}
}

func slideSouth(l *util.FixedLevel) {
	w, h := l.W, l.H
	for x := 0; x < w; x++ {
		for y, startY, rocks := 0, 0, 0; y <= h; y++ {
			if y == h || l.At(x, y) == '#' {
				if startY < y-1 && rocks > 0 {
					y2 := y - 1
					for ; y2 >= y-rocks; y2-- {
						l.Set(x, y2, 'O')
					}
					for ; y2 >= startY; y2-- {
						l.Set(x, y2, '.')
					}
				}
				startY, rocks = y+1, 0
				continue
			}
			if l.At(x, y) == 'O' {
				rocks++
			}
		}
	}
}

func slideEast(l *util.FixedLevel) {
	w, h := l.W, l.H
	for y := 0; y < h; y++ {
		for x, startX, rocks := 0, 0, 0; x <= w; x++ {
			if x == w || l.At(x, y) == '#' {
				if startX < x-1 && rocks > 0 {
					x2 := x - 1
					for ; x2 >= x-rocks; x2-- {
						l.Set(x2, y, 'O')
					}
					for ; x2 >= startX; x2-- {
						l.Set(x2, y, '.')
					}
				}
				startX, rocks = x+1, 0
				continue
			}
			if l.At(x, y) == 'O' {
				rocks++
			}
		}
	}
}
