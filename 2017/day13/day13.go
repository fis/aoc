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

// Package day13 solves AoC 2017 day 13.
package day13

import (
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 13, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	layers, err := parseInput(lines)
	if err != nil {
		return nil, err
	}
	_, p1 := severity(layers, 0)
	p2 := crack(layers)
	return glue.Ints(p1, p2), nil
}

type layer struct {
	depth int
	size  int
}

func severity(layers []layer, delay int) (caught bool, sev int) {
	for _, l := range layers {
		pos := (delay + l.depth) % (2 * (l.size - 1))
		if pos == 0 {
			caught = true
			sev += l.depth * l.size
		}
	}
	return caught, sev
}

func crack(layers []layer) (delay int) {
	delay = 0
loop:
	for {
		for _, l := range layers {
			pos := (delay + l.depth) % (2 * (l.size - 1))
			if pos == 0 {
				delay++
				continue loop
			}
		}
		return delay
	}
}

func parseInput(lines []string) (layers []layer, err error) {
	layers = make([]layer, len(lines))
	for i, line := range lines {
		if _, err := fmt.Sscanf(line, "%d: %d", &layers[i].depth, &layers[i].size); err != nil {
			return nil, fmt.Errorf("invalid line: %q: %w", line, err)
		}
	}
	return layers, nil
}
