// Copyright 2020 Google LLC
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

// Package day01 solves AoC 2020 day 1.
package day01

import (
	"strconv"

	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	expenses, err := util.ReadIntRows(path)
	if err != nil {
		return nil, err
	}
	return []string{
		strconv.Itoa(part1(expenses)),
		strconv.Itoa(part2(expenses)),
	}, nil
}

func part1(expenses []int) int {
	for i, l := range expenses[0 : len(expenses)-1] {
		for _, r := range expenses[i+1:] {
			if l+r == 2020 {
				return l * r
			}
		}
	}
	return -1
}

func part2(expenses []int) int {
	for i, a := range expenses[0 : len(expenses)-2] {
		for j, b := range expenses[i+1 : len(expenses)-1] {
			for _, c := range expenses[i+j+2:] {
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return -1
}
