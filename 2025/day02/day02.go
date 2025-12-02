// Copyright 2025 Google LLC
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

// Package day02 solves AoC 2025 day 2.
package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/ix"
)

func init() {
	glue.RegisterSolver(2025, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line, got %d", len(lines))
	}
	ranges, err := parseRanges(lines[0])
	if err != nil {
		return nil, err
	}
	p1, p2 := countInvalidFast(ranges)
	return glue.Ints(p1, p2), nil
}

func countInvalid(ranges [][2]int) (p1, p2 int) {
	for _, r := range ranges {
		start, end := r[0], r[1]
	nextNum:
		for i := start; i <= end; i++ {
			l := ix.Len(i)
			// test for 123:123, invalid by both definitions, possible only for even lengths
			if l&1 == 0 {
				p := ix.Pow(10, l/2)
				if i/p == i%p {
					p1 += i
					p2 += i
					continue nextNum
				}
			}
			// test for 3, 5, ... repeats; 4, 6, ... already covered (ab:ab:ab:ab is necessarily also abab:abab)
		nextLen:
			for n := 3; n <= l; n += 2 {
				if l%n != 0 {
					continue nextLen // doesn't divide evenly
				}
				bl := l / n
				p := ix.Pow(10, bl)
				b := i % p
				for j := i / p; j > 0; j /= p {
					if j%p != b {
						continue nextLen // doesn't repeat
					}
				}
				p2 += i
				continue nextNum
			}
		}
	}
	return p1, p2
}

func countInvalidFast(ranges [][2]int) (p1, p2 int) {
	reps := make(map[int]bool)
	for _, r := range ranges {
		clear(reps)
		startL, endL := ix.Len(r[0]), ix.Len(r[1])
		for l := startL; l <= endL; l++ {
			start, end := r[0], r[1]
			if l > startL {
				start = ix.Pow(10, l-1)
			}
			if l < endL {
				end = ix.Pow(10, l) - 1
			}
			for bl := 1; bl <= l/2; bl++ {
				if l%bl != 0 {
					continue // doesn't divide evenly, can't be repeating
				}
				r := l / bl
				p, pp := ix.Pow(10, bl), ix.Pow(10, l-bl)
				startB, endB := start/pp, end/pp
				for b := startB; b <= endB; b++ {
					c := 0
					for bb, j := b, 0; j < r; bb, j = bb*p, j+1 {
						c += bb
					}
					if c >= start && c <= end {
						reps[c] = reps[c] || 2*bl == l
					}
				}
			}
		}
		for c, isHalf := range reps {
			if isHalf {
				p1 += c
			}
			p2 += c
		}
	}
	return p1, p2
}

func parseRanges(line string) (ranges [][2]int, err error) {
	for spec := range strings.SplitSeq(line, ",") {
		start, end, ok := strings.Cut(spec, "-")
		if !ok {
			return nil, fmt.Errorf("invalid range: no separator: %q", spec)
		}
		var r [2]int
		for i, n := range [2]string{start, end} {
			r[i], err = strconv.Atoi(n)
			if err != nil {
				return nil, fmt.Errorf("invalid endpoint: not a number: %q: %w", n, err)
			}
		}
		ranges = append(ranges, r)
	}
	return ranges, nil
}
