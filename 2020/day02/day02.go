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

// Package day02 solves AoC 2020 day 2.
package day02

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2020, 2, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	p1, err := countValid(lines, policy.validateSled)
	if err != nil {
		return nil, err
	}
	p2, err := countValid(lines, policy.validateToboggan)
	if err != nil {
		return nil, err
	}
	return []int{p1, p2}, nil
}

func countValid(lines []string, validator func(policy, string) bool) (int, error) {
	valid := 0
	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) < 2 {
			return 0, fmt.Errorf("invalid password record: %q", line)
		}
		pol, err := parsePolicy(parts[0])
		if err != nil {
			return 0, err
		}
		if validator(pol, parts[1]) {
			valid++
		}
	}
	return valid, nil
}

type policy struct {
	Min, Max int
	C        rune
}

func parsePolicy(p string) (policy, error) {
	var (
		min, max int
		c        rune
	)
	if n, err := fmt.Sscanf(p, "%d-%d %c", &min, &max, &c); err != nil || n != 3 {
		return policy{}, fmt.Errorf("invalid policy string: %q", p)
	}
	return policy{Min: int(min), Max: int(max), C: c}, nil
}

func (p policy) validateSled(pass string) bool {
	freq := 0
	for _, c := range pass {
		if c == p.C {
			freq++
		}
	}
	return freq >= p.Min && freq <= p.Max
}

func (p policy) validateToboggan(pass string) bool {
	c1, _ := utf8.DecodeRuneInString(pass[p.Min-1:])
	c2, _ := utf8.DecodeRuneInString(pass[p.Max-1:])
	return (c1 == p.C) != (c2 == p.C)
}

func (p policy) String() string {
	return fmt.Sprintf("(%d-%d %c)", p.Min, p.Max, p.C)
}
