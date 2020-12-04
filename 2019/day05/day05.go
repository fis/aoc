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

// Package day05 solves AoC 2019 day 5.
package day05

import (
	"strconv"

	"github.com/fis/aoc-go/intcode"
	"github.com/fis/aoc-go/util"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}
	return []string{
		strconv.FormatInt(part1(prog), 10),
		strconv.FormatInt(part2(prog), 10),
	}, nil
}

func part1(prog []int64) int64 {
	out, _ := intcode.Run(prog, []int64{1})
	for _, i := range out {
		util.Diagf("out: %d\n", i)
	}
	return out[len(out)-1]
}

func part2(prog []int64) int64 {
	out, _ := intcode.Run(prog, []int64{5})
	return out[0]
}
