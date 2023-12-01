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

// Package day01 solves AoC 2023 day 1.
package day01

import (
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 1, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1 := fn.SumF(lines, calibrationValue)
	p2 := fn.SumF(lines, calibrationValueEx)
	return glue.Ints(p1, p2), nil
}

const digits = "123456789"

var words = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func calibrationValue(line string) int {
	i1, i2 := strings.IndexAny(line, digits), strings.LastIndexAny(line, digits)
	d1, d2 := line[i1]-'0', line[i2]-'0'
	return 10*int(d1) + int(d2)
}

func calibrationValueEx(line string) int {
	i1, i2 := strings.IndexAny(line, digits), strings.LastIndexAny(line, digits)
	if i1 == -1 {
		i1 = len(line) - 1
	}
	if i2 == -1 {
		i2 = 0
	}
	d1, d2 := int(line[i1]-'0'), int(line[i2]-'0')
	for wi, word := range words {
		if i := strings.Index(line[:i1+1], word); i >= 0 {
			i1, d1 = i+len(word)-1, 1+wi
		}
		if i := strings.LastIndex(line[i2:], word); i >= 0 {
			i2, d2 = i2+i, 1+wi
		}
	}
	return 10*d1 + d2
}
