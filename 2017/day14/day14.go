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

// Package day14 solves AoC 2017 day 14.
package day14

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/2017/knot"
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	hashes, p1 := hash(lines[0])
	lvl := buildLevel(hashes)
	p2 := countRegions(lvl)
	return glue.Ints(p1, p2), nil
}

func hash(key string) (hashes [][]byte, used int) {
	for i := 0; i < 128; i++ {
		rowKey := fmt.Sprintf("%s-%d", key, i)
		hash := knot.Hash(knot.N, knot.Rounds, rowKey)
		hashes = append(hashes, hash)
		for _, b := range hash {
			used += bits.OnesCount8(b)
		}
	}
	return hashes, used
}

func buildLevel(hashes [][]byte) *util.Level {
	lvl := util.EmptyLevel(util.P{0, 0}, util.P{knot.N - 1, len(hashes) - 1}, '.')
	for y, row := range hashes {
		for x1, b := range row {
			for x2 := 0; x2 < 8; x2++ {
				if ((b << x2) & 0x80) != 0 {
					lvl.Set(8*x1+x2, y, '#')
				}
			}
		}
	}
	return lvl
}

func countRegions(lvl *util.Level) (regions int) {
	seen := map[util.P]struct{}{}
	lvl.Range(func(x, y int, b byte) {
		if _, found := seen[util.P{x, y}]; found {
			return
		}
		regions++
		edge := []util.P{{x, y}}
		for len(edge) > 0 {
			at := edge[len(edge)-1]
			edge = edge[:len(edge)-1]
			if _, found := seen[at]; found {
				continue
			}
			seen[at] = struct{}{}
			for _, n := range at.Neigh() {
				if _, found := seen[n]; !found && lvl.At(n.X, n.Y) == '#' {
					edge = append(edge, n)
				}
			}
		}
	})
	return regions
}
