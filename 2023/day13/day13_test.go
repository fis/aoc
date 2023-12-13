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

package day13

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
`, "\n")

func TestSumReflections(t *testing.T) {
	ps := fn.Map(util.Chunks(ex), parsePattern)
	tests := []struct {
		name string
		find func(p pattern) (rx, ry int)
		want int
	}{
		{"findReflection", findReflection, 405},
		{"findSmudged", findSmudged, 400},
	}
	for _, test := range tests {
		if got := sumReflections(ps, test.find); got != test.want {
			t.Errorf("sumReflections(ex, %s) = %d, want %d", test.name, got, test.want)
		}

	}
}
