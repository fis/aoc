// Copyright 2019 Google LLC
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

// Package day09 solves AoC 2019 day 9.
package day09

import (
	"github.com/fis/aoc-go/2019/intcode"
	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2019, 9, intcode.Solver(solve))
}

func solve(prog []int64) ([]int64, error) {
	p1, _ := intcode.Run(prog, []int64{1})
	p2, _ := intcode.Run(prog, []int64{2})
	return []int64{p1[len(p1)-1], p2[len(p2)-1]}, nil
}
