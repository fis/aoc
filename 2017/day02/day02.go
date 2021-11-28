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

// Package day02 solves AoC 2017 day 2.
package day02

import (
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	sheet := make([][]int, len(lines))
	for i, line := range lines {
		words := util.Words(line)
		row := make([]int, len(words))
		for j, word := range words {
			n, err := strconv.Atoi(word)
			if err != nil {
				return nil, fmt.Errorf("not a number: %q: %w", word, err)
			}
			row[j] = n
		}
		sheet[i] = row
	}
	p1 := checksumSheet(sheet)
	p2, err := divisibleSheet(sheet)
	if err != nil {
		return nil, err
	}
	return glue.Ints(p1, p2), nil
}

func checksumSheet(rows [][]int) (sum int) {
	for _, row := range rows {
		sum += checksumRow(row)
	}
	return sum
}

func checksumRow(row []int) int {
	min, max := row[0], row[0]
	for _, n := range row[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min
}

func divisibleSheet(rows [][]int) (sum int, err error) {
	for _, row := range rows {
		d, err := divisibleRow(row)
		if err != nil {
			return 0, err
		}
		sum += d
	}
	return sum, nil
}

func divisibleRow(row []int) (int, error) {
	for _, i := range row {
		for _, j := range row {
			if i > j && i%j == 0 {
				return i / j, nil
			}
		}
	}
	return 0, fmt.Errorf("invalid input: %v: no evenly divisible pair", row)
}
