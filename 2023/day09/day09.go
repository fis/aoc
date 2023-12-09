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

// Package day09 solves AoC 2023 day 9.
package day09

import (
	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2023, 9, glue.LineSolver(glue.WithParser(func(line string) ([]int, error) { return util.Ints(line), nil }, solve)))
}

func solve(seqs [][]int) ([]string, error) {
	p1 := fn.SumF(seqs, predict)
	p2 := fn.SumF(seqs, unpredict)
	return glue.Ints(p1, p2), nil
}

func predict(seq []int) int {
	buf := append([]int(nil), seq...)
	var n int
	for n = len(buf) - 1; n > 0; n-- {
		nz := false
		for i := 0; i < n; i++ {
			buf[i] = buf[i+1] - buf[i]
			nz = nz || buf[i] != 0
		}
		if !nz {
			break
		}
	}
	return fn.Sum(buf[n:])
}

func unpredict(seq []int) int {
	buf := append([]int(nil), seq...)
	var s int
	for s = 1; s < len(buf); s++ {
		nz := false
		for i := len(buf) - 1; i >= s; i-- {
			buf[i] = buf[i-1] - buf[i]
			nz = nz || buf[i] != 0
		}
		if !nz {
			break
		}
	}
	return fn.Sum(buf[:s])
}
