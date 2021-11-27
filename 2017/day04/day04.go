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

// Package day04 solves AoC 2017 day 4.
package day04

import (
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 4, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	p1, p2 := 0, 0
	for _, line := range lines {
		words := util.Words(line)
		if part1(words) {
			p1++
		}
		if part2(words) {
			p2++
		}
	}
	return []int{p1, p2}, nil
}

func part1(passphrase []string) bool {
	seen := make(map[string]struct{})
	for _, word := range passphrase {
		if _, found := seen[word]; found {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}

func part2(passphrase []string) bool {
	seen := make(map[string]struct{})
	for _, word := range passphrase {
		chars := []byte(word)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
		word = string(chars)
		if _, found := seen[word]; found {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}
