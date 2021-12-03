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

// Package day03 solves AoC 2021 day 3.
package day03

import (
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 3, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	gamma, epsilon := gammaEpsilon(lines)
	oxygen := filterBits(lines, 0)
	co2 := filterBits(lines, 1)
	return glue.Ints(gamma*epsilon, oxygen*co2), nil
}

func gammaEpsilon(lines []string) (gamma, epsilon int) {
	N, K := len(lines), len(lines[0])
	ones := make([]int, K)
	for _, line := range lines {
		for i := 0; i < K; i++ {
			if line[i]&1 != 0 {
				ones[i]++
			}
		}
	}
	for i := 0; i < K; i++ {
		gamma <<= 1
		if ones[i] > N/2 {
			gamma |= 1
		}
	}
	epsilon = gamma ^ ((1 << K) - 1)
	return gamma, epsilon
}

func filterBits(lines []string, keepLCB int) int {
	var surviving []string
	K := len(lines[0])
	for i := 0; i < K; i++ {
		ones := 0
		for _, line := range lines {
			if line[i]&1 != 0 {
				ones++
			}
		}
		keep := (2 * ones / len(lines)) ^ keepLCB
		for _, line := range lines {
			if int(line[i]&1) == keep {
				surviving = append(surviving, line)
			}
		}
		lines = surviving
		if len(lines) == 1 {
			n, _ := strconv.ParseInt(lines[0], 2, 64)
			return int(n)
		}
		surviving = lines[:0]
	}
	panic("invalid input: no unique survivor")
}
