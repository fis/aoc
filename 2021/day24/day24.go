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

// Package day24 solves AoC 2021 day 24.
package day24

import (
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2021, 24, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	coeffs := parseCoeffs(lines)
	p1 := coeffs.findValid(true)
	p2 := coeffs.findValid(false)
	return []string{p1, p2}, nil
}

var (
	bigFirst   = [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	smallFirst = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func (coeffs *monadCoeffs) findValid(wantBig bool) (num string) {
	var digits [ndigits]byte
	pairs := coeffs.pairs()
	for _, pair := range pairs {
		d := coeffs[pair[0]].b + coeffs[pair[1]].a
		var d1, d2 int
		if wantBig {
			d1, d2 = ix.Min(9, 9-d), ix.Min(9, 9+d)
		} else {
			d1, d2 = ix.Max(1, 1-d), ix.Max(1, 1+d)
		}
		digits[pair[0]] = byte('0' + d1)
		digits[pair[1]] = byte('0' + d2)
	}
	return string(digits[:])
}

func (coeffs *monadCoeffs) pairs() (pairs [][2]int) {
	var stack []int
	for i, c := range coeffs {
		if !c.pop {
			stack = append(stack, i)
		} else if len(stack) > 0 {
			pairs = append(pairs, [2]int{stack[len(stack)-1], i})
			stack = stack[:len(stack)-1]
		} else {
			panic("coefficients do not pair up: stack empty")
		}
	}
	if len(stack) > 0 {
		panic("coefficients do not pair up: leftovers")
	}
	return pairs
}

const ndigits = 14

type monadCoeff struct {
	pop  bool
	a, b int
}

type monadCoeffs [ndigits]monadCoeff

func parseCoeffs(lines []string) (coeffs monadCoeffs) {
	for i := 0; i < ndigits; i++ {
		base := i * 18
		if words := strings.Split(lines[base+4], " "); words[2] == "26" {
			coeffs[i].pop = true
		}
		words := strings.Split(lines[base+5], " ")
		coeffs[i].a, _ = strconv.Atoi(words[2])
		words = strings.Split(lines[base+15], " ")
		coeffs[i].b, _ = strconv.Atoi(words[2])
	}
	return coeffs
}
