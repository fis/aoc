// Copyright 2024 Google LLC
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

package day16

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var (
	ex1 = strings.TrimPrefix(`
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############
`, "\n")
	ex2 = strings.TrimPrefix(`
#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################
	`, "\n")
)

func TestBestPath(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		want1, want2 int
	}{
		{"ex1", ex1, 7036, 45},
		{"ex2", ex2, 11048, 64},
	}
	for _, test := range tests {
		level := util.ParseFixedLevel([]byte(test.input))
		if got1, got2 := bestPath(level); got1 != test.want1 || got2 != test.want2 {
			t.Errorf("bestPath(%s) = (%d,%d), want (%d,%d)", test.name, got1, got2, test.want1, test.want2)
		}
	}
}
