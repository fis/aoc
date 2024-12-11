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

// Package day11 solves AoC 2024 day 11.
package day11

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2024, 11, glue.IntSolver(solve))
}

func solve(stones []int) ([]string, error) {
	var c counter
	part1 := c.countAll(stones, 25)
	part2 := c.countAll(stones, 75)
	return glue.Ints(part1, part2), nil
}

const (
	maxSteps = 75
	cacheSize = 1024
)

type counter struct {
	memo [maxSteps][cacheSize]int
}

func (c *counter) countAll(stones []int, steps int) (sum int) {
	for _, s := range stones {
		sum += c.count(s, steps)
	}
	return sum
}

func (c *counter) count(stone, steps int) (count int) {
	if steps == 0 {
		return 1
	}
	if count, ok := c.check(stone, steps); ok {
		return count
	}
	if stone == 0 {
		count = c.count(1, steps-1)
	} else {
		left, right, ok := split(stone)
		if ok {
			count = c.count(left, steps-1) + c.count(right, steps-1)
		} else {
			count = c.count(2024*stone, steps-1)
		}
	}
	c.cache(stone, steps, count)
	return count
}

func (c *counter) check(stone, steps int) (count int, ok bool) {
	if stone < cacheSize {
		count = c.memo[steps-1][stone]
	}
	return count, count != 0
}

func (c *counter) cache(stone, steps, count int) {
	if stone < cacheSize {
		c.memo[steps-1][stone] = count
	}
}

func split(stone int) (left, right int, ok bool) {
	n, d := 10, 10
	for {
		if stone < n {
			return 0, 0, false
		}
		if stone < 10*n {
			return stone / d, stone % d, true
		}
		n, d = n*100, d*10
	}
}
