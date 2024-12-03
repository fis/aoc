// Copyright 2024 Google LLC
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

// Package day03 solves AoC 2024 day 3.
package day03

import (
	"io"
	"regexp"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2024, 3, glue.GenericSolver(solve))
}

func solve(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	mem := string(data)
	part1, part2 := findInstructions(mem)
	return glue.Ints(part1, part2), nil
}

var instRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)|(do(?:n't)?)\(\)`)

func findInstructions(mem string) (uncondSum, condSum int) {
	enabled := true
	for _, match := range instRegexp.FindAllStringSubmatch(mem, -1) {
		switch match[3] {
		case "do":
			enabled = true
			continue
		case "don't":
			enabled = false
			continue
		}
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		val := a * b
		uncondSum += val
		if enabled {
			condSum += val
		}
	}
	return uncondSum, condSum
}
