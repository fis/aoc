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
	"strconv"

	"github.com/fis/aoc-go/intcode"
)

func Solve(path string) ([]string, error) {
	prog, err := intcode.Load(path)
	if err != nil {
		return nil, err
	}

	p1, _ := intcode.Run(prog, []int64{1})
	p2, _ := intcode.Run(prog, []int64{2})

	return []string{
		strconv.FormatInt(p1[len(p1)-1], 10),
		strconv.FormatInt(p2[len(p2)-1], 10),
	}, nil
}
