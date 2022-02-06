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

// Package day06 solves AoC 2016 day 6.
package day06

import (
	"math"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2016, 6, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	p1, p2 := decode(lines)
	return []string{p1, p2}, nil
}

func decode(lines []string) (string, string) {
	N := len(lines[0])
	m1, m2 := make([]byte, N), make([]byte, N)
	for i := 0; i < N; i++ {
		var freqs [26]int
		for _, line := range lines {
			if b := line[i]; b >= 'a' && b <= 'z' {
				freqs[b-'a']++
			}
		}
		maxB, maxF, minB, minF := byte(0), math.MinInt, byte(0), math.MaxInt
		for b := byte(0); b < 26; b++ {
			f := freqs[b]
			if f == 0 {
				continue
			}
			if f > maxF {
				maxB, maxF = b, f
			}
			if f < minF {
				minB, minF = b, f
			}
		}
		m1[i] = 'a' + maxB
		m2[i] = 'a' + minB
	}
	return string(m1), string(m2)
}
