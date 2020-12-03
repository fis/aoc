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

// Package day03 solves AoC 2020 day 3.
package day03

import (
	"strconv"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	level, err := util.ReadLevel(path, '.')
	if err != nil {
		return nil, err
	}

	trees := countTrees(level, util.P{3, 1})

	allTrees := trees
	for _, slope := range []util.P{{1, 1}, {5, 1}, {7, 1}, {1, 2}} {
		allTrees *= countTrees(level, slope)
	}

	return []string{strconv.Itoa(trees), strconv.Itoa(allTrees)}, nil
}

func countTrees(level *util.Level, slope util.P) int {
	trees := 0
	_, max := level.Bounds()
	for x, y, w := 0, 0, max.X+1; y <= max.Y; x, y = (x+slope.X)%w, y+slope.Y {
		if level.At(x, y) == '#' {
			trees++
		}
	}
	return trees
}
