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

package day03

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const example = `
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestCountTrees(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(example, "\n"), '.')

	tests := []struct {
		slope util.P
		want  int
	}{
		{slope: util.P{1, 1}, want: 2},
		{slope: util.P{3, 1}, want: 7},
		{slope: util.P{5, 1}, want: 3},
		{slope: util.P{7, 1}, want: 4},
		{slope: util.P{1, 2}, want: 2},
	}
	for _, test := range tests {
		got := countTrees(level, test.slope)
		if got != test.want {
			t.Errorf("countTrees(..., %v) = %d, want %d", test.slope, got, test.want)
		}
	}
}
