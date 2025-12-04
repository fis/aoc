// Copyright 2025 Google LLC
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

// Package day03 solves AoC 2025 day 3.
package day03

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2025, 3, glue.LineSolver(solve))
}

func solve(banks []string) ([]string, error) {
	p1 := fn.SumF(banks, maxJoltage2)
	p2 := fn.SumF(banks, func(bank string) int { return maxJoltageN(bank, 12) })
	return glue.Ints(p1, p2), nil
}

func maxJoltage2(bank string) int {
	n := len(bank)
	if n < 2 {
		return -1
	}
	i1, d1 := maxDigit(bank[:n-1])
	_, d2 := maxDigit(bank[i1+1:])
	return 10*d1 + d2
}

func maxJoltageN(bank string, c int) int {
	n := len(bank)
	if n < c {
		return -1
	}
	at, joltage := 0, 0
	for ; c > 0; c-- {
		i, d := maxDigit(bank[at : n-(c-1)])
		joltage = 10*joltage + d
		at += i + 1
	}
	return joltage
}

func maxDigit(s string) (maxI, maxD int) {
	for i, b := range []byte(s) {
		if d := int(b - '0'); d > maxD {
			maxI, maxD = i, d
		}
	}
	return
}
