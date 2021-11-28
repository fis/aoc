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

// Package day15 solves AoC 2017 day 15.
package day15

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 15, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^Generator ([AB]) starts with (\d+)$`,
	})
}

func solve(input [][]string) ([]string, error) {
	if len(input) != 2 || input[0][0] != "A" || input[1][0] != "B" {
		return nil, fmt.Errorf("invalid input: expected initial values for generators A and B")
	}
	xA, _ := strconv.Atoi(input[0][1])
	xB, _ := strconv.Atoi(input[1][1])
	p1 := judge(xA, xB, 40000000)
	p2 := judge2(xA, xB, 5000000)
	return glue.Ints(p1, p2), nil
}

const (
	mA  = 16807
	mB  = 48271
	div = 2147483647
)

func judge(xA, xB, N int) (matches int) {
	for i := 0; i < N; i++ {
		xA = (xA * mA) % div
		xB = (xB * mB) % div
		if (xA & 0xffff) == (xB & 0xffff) {
			matches++
		}
	}
	return matches
}

func judge2(xA, xB, N int) (matches int) {
	for i := 0; i < N; i++ {
		for {
			xA = (xA * mA) % div
			if xA%4 == 0 {
				break
			}
		}
		for {
			xB = (xB * mB) % div
			if xB%8 == 0 {
				break
			}
		}
		if (xA & 0xffff) == (xB & 0xffff) {
			matches++
		}
	}
	return matches
}
