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

// Package day25 solves AoC 2022 day 25.
package day25

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2022, 25, glue.LineSolver(glue.WithParser(parseSNAFU, solve)))
}

func solve(nums []int) ([]string, error) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return []string{formatSNAFU(sum)}, nil
}

func parseSNAFU(s string) (i int, err error) {
	for _, c := range s {
		var d int
		switch c {
		case '=':
			d = -2
		case '-':
			d = -1
		default:
			d = int(c) - '0'
		}
		i = 5*i + d
	}
	return i, nil
}

func formatSNAFU(i int) string {
	const snafuDigits = "=-012"
	var buf [28]byte // sufficient for nonnegative 64-bit signed integers
	n := 28
	carry := 0
	for i > 0 || carry == 1 {
		d := i%5 + carry
		i /= 5
		if d > 2 {
			carry = 1
			d -= 5
		} else {
			carry = 0
		}
		n--
		buf[n] = snafuDigits[2+d]
	}
	return string(buf[n:])
}
