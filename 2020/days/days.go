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

// Package days contains the glue and tests for all AoC 2020 days.
package days

import (
	"fmt"

	"github.com/fis/aoc-go/2020/day01"
)

var solvers = map[int]func(string) ([]string, error){
	1: day01.Solve,
}

func Solve(day int, path string) ([]string, error) {
	solver, ok := solvers[day]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d", day)
	}
	return solver(path)
}
