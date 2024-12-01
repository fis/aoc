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

// Package day01 solves AoC 2024 day 1.
package day01

import (
	"slices"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2024, 1, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	listA, listB := parse(lines)
	part1 := totalDistance(listA, listB)
	part2 := similarity(listA, listB)
	return glue.Ints(part1, part2), nil
}

func parse(lines []string) (listA, listB []int) {
	for _, line := range lines {
		ints := util.Ints(line)
		listA = append(listA, ints[0])
		listB = append(listB, ints[1])
	}
	slices.Sort(listA)
	slices.Sort(listB)
	return listA, listB
}

func totalDistance(listA, listB []int) (dist int) {
	for i := range listA {
		dist += ix.Abs(listA[i] - listB[i])
	}
	return dist
}

func similarity(listA, listB []int) (sim int) {
	i, j := 0, 0
	for i < len(listA) && j < len(listB) {
		vA, vB := listA[i], listB[j]
		switch {
		case vA < vB:
			i++
			continue
		case vA > vB:
			j++
			continue
		}
		nA, nB := 0, 0
		for i < len(listA) && listA[i] == vA {
			nA++
			i++
		}
		for j < len(listB) && listB[j] == vB {
			nB++
			j++
		}
		sim += nA * nB * vA
	}
	return sim
}
