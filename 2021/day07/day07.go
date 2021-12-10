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

// Package day07 solves AoC 2021 day 7.
package day07

import (
	"math"
	"sort"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 7, glue.IntSolver(solve))
}

func solve(input []int) ([]string, error) {
	_, p1 := align1MedianQS(input)
	_, p2 := align2Mean(input)
	return glue.Ints(p1, p2), nil
}

func align(input []int, f func(n, x int) int) (x, cost int) {
	min, max := bounds(input)
	costs := make([]int, max-min+1)
	for _, n := range input {
		for i := range costs {
			costs[i] += f(n, min+i)
		}
	}
	x, cost = argmin(costs)
	return min + x, cost
}

func cost1(n, x int) int {
	return abs(n - x)
}

func cost2(n, x int) int {
	d := abs(n - x)
	return d * (d + 1) / 2
}

func align1Points(input []int) (x, cost int) {
	min, max := bounds(input)
	costs := make([]int, max-min+1)
	for i := range costs {
		costs[i] = math.MaxInt
	}
	for _, x := range input {
		if costs[x-min] == math.MaxInt {
			costs[x-min] = 0
			for _, n := range input {
				costs[x-min] += abs(n - x)
			}
		}
	}
	x, cost = argmin(costs)
	return min + x, cost
}

func align1MedianSort(input []int) (x, cost int) {
	sorted := append([]int(nil), input...)
	sort.Ints(sorted)
	x = sorted[len(sorted)/2]
	for _, n := range input {
		cost += abs(n - x)
	}
	return x, cost
}

func align1MedianQS(input []int) (x, cost int) {
	x = quickSelect(input, len(input)/2)
	for _, n := range input {
		cost += abs(n - x)
	}
	return x, cost
}

func align2Mean(input []int) (x, cost int) {
	mean := 0
	for _, n := range input {
		mean += n
	}
	mean = (mean + len(input)/2) / len(input)
	cost = math.MaxInt
	for estX := mean - 2; estX <= mean+2; estX++ {
		estCost := 0
		for _, n := range input {
			d := abs(n - estX)
			estCost += d * (d + 1) / 2
		}
		if estCost < cost {
			cost = estCost
			x = estX
		}
	}
	return x, cost
}

func quickSelect(input []int, k int) int {
	const cutoff = 12

	origInput := true
	next := []int(nil)

	for len(input) > cutoff {
		if origInput {
			origInput = false
			input = append([]int(nil), input...)
			next = make([]int, len(input))
		}

		_, pivot, _ := util.Sort3(input[0], input[len(input)/2], input[len(input)-1])
		lt, gt := 0, 0
		for _, n := range input {
			switch {
			case n < pivot:
				next[lt] = n
				lt++
			case n > pivot:
				next[len(next)-1-gt] = n
				gt++
			}
		}
		switch {
		case k < lt:
			input, next = next[:lt], input[:lt]
		case k >= len(input)-gt:
			input, next, k = next[len(input)-gt:], input[len(input)-gt:], k-(len(input)-gt)
		default:
			return pivot
		}
	}

	if len(input) == 1 {
		return input[0]
	} else if len(input) == 2 {
		if k == 0 {
			if input[0] <= input[1] {
				return input[0]
			} else {
				return input[1]
			}
		} else {
			if input[0] <= input[1] {
				return input[0]
			} else {
				return input[1]
			}
		}
	} else if len(input) == 3 {
		a, b, c := util.Sort3(input[0], input[1], input[2])
		switch k {
		case 0:
			return a
		case 1:
			return b
		default:
			return c
		}
	}

	if origInput {
		var tmp [cutoff]int
		copy(tmp[:len(input)], input)
		input = tmp[:len(input)]
	}
	sort.Ints(input)
	return input[k]
}

func bounds(input []int) (min, max int) {
	min = input[0]
	max = input[0]
	for _, n := range input[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func argmin(input []int) (minI, minN int) {
	minI, minN = 0, input[0]
	for i, n := range input[1:] {
		if n < minN {
			minI = i + 1
			minN = n
		}
	}
	return minI, minN
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
