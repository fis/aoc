// Copyright 2023 Google LLC
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

// Package day03 solves AoC 2023 day 3.
package day03

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 3, glue.LevelSolver{Solver: solve, Empty: '.'})
}

func solve(lv *util.Level) ([]string, error) {
	p1 := sumNumbers(lv, 0, fn.Sum[[]int])
	p2 := sumNumbers(lv, '*', gearRatio)
	return glue.Ints(p1, p2), nil
}

func gearRatio(nums []int) int {
	if len(nums) == 2 {
		return nums[0] * nums[1]
	}
	return 0
}

func sumNumbers(lv *util.Level, symFilter byte, f func([]int) int) (sum int) {
	lvMin, lvMax := lv.Bounds()
	for y := lvMin.Y; y <= lvMax.Y; y++ {
		for x := lvMin.X; x <= lvMax.X; x++ {
			sym := lv.At(x, y)
			if symFilter != 0 {
				if sym != symFilter {
					continue
				}
			} else {
				if sym == '.' || sym >= '0' && sym <= '9' {
					continue
				}
			}
			var nums []int
			for dy := y - 1; dy <= y+1; dy++ {
				for dx := x - 1; dx <= x+1; dx++ {
					num, nx := findNumber(lv, dx, dy)
					if num >= 0 {
						nums = append(nums, num)
						dx = nx
					}
				}
			}
			if len(nums) > 0 {
				sum += f(nums)
			}
		}
	}
	return sum
}

func findNumber(lv *util.Level, x, y int) (num, newX int) {
	d := lv.At(x, y)
	if d < '0' || d > '9' {
		return -1, x
	}
	for {
		if c := lv.At(x-1, y); c < '0' || c > '9' {
			break
		}
		x--
	}
	num = int(lv.At(x, y) - '0')
	for {
		d := lv.At(x+1, y)
		if d < '0' || d > '9' {
			break
		}
		x++
		num = 10*num + int(d-'0')
	}
	return num, x
}
