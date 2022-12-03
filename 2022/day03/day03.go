// Copyright 2022 Google LLC
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

// Package day03 solves AoC 2022 day 3.
package day03

import (
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 3, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	sacks := pack(lines)
	p1 := part1(sacks)
	p2 := part2(sacks)
	return glue.Ints(p1, p2), nil
}

func part1(sacks []rucksack) int {
	return fn.Sum(fn.Map(sacks, rucksack.overlap))
}

func part2(sacks []rucksack) (sum int) {
	for i := 0; i < len(sacks); i += 3 {
		sum += overlap3(sacks[i], sacks[i+1], sacks[i+2])
	}
	return sum
}

func pack(lines []string) (sacks []rucksack) {
	sacks = make([]rucksack, len(lines))
	for i, line := range lines {
		sacks[i].pack(line)
	}
	return sacks
}

type rucksack [2]uint64

func (rs *rucksack) pack(items string) {
	rs[0], rs[1] = 0, 0
	n := len(items) / 2
	for i := 0; i < n; i++ {
		rs[0] |= 1 << itemPriority(items[i])
		rs[1] |= 1 << itemPriority(items[n+i])
	}
}

func (rs rucksack) overlap() int {
	return bits.TrailingZeros64(rs[0] & rs[1])
}

func overlap3(a, b, c rucksack) int {
	is := (a[0] | a[1]) & (b[0] | b[1]) & (c[0] | c[1])
	return bits.TrailingZeros64(is)
}

func itemPriority(i byte) int {
	if i >= 'a' {
		return 1 + int(i-'a')
	}
	return 27 + int(i-'A')
}
