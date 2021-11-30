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

// Package day23 solves AoC 2017 day 23.
package day23

import (
	"errors"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2017, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	test, low, high, skip, err := extractParams(lines)
	if err != nil {
		return nil, err
	}
	p1 := part1(test)
	p2 := countComposite(low, high, skip)
	return glue.Ints(p1, p2), nil
}

func part1(n int) int {
	return (n - 2) * (n - 2)
}

func countComposite(low, high, skip int) (composites int) {
	for n := low; n <= high; n += skip {
		if !isPrime(n) {
			composites++
		}
	}
	return composites
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for f := 3; f*f <= n; f += 2 {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func extractParams(lines []string) (test, low, high, skip int, err error) {
	var mb, ib, ic int
	for _, ex := range []struct {
		line   int
		prefix string
		out    *int
	}{
		{line: 0, prefix: "set b ", out: &test},
		{line: 4, prefix: "mul b ", out: &mb},
		{line: 5, prefix: "sub b -", out: &ib},
		{line: 7, prefix: "sub c -", out: &ic},
		{line: 30, prefix: "sub b -", out: &skip},
	} {
		line := lines[ex.line]
		if !strings.HasPrefix(line, ex.prefix) {
			return 0, 0, 0, 0, errors.New("bad program")
		}
		*ex.out, _ = strconv.Atoi(strings.TrimPrefix(line, ex.prefix))
	}
	low = test*mb + ib
	high = low + ic
	return test, low, high, skip, nil
}
