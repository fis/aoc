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

package day11

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`, "\n")

func TestTotalDistance(t *testing.T) {
	tests := []struct {
		factor int
		want   int
	}{
		{2, 374},
		{10, 1030},
		{100, 8410},
	}
	galaxies := findGalaxies(util.Lines(ex))
	for _, test := range tests {
		expanded := expandSpace(galaxies, test.factor)
		if got := totalDistance(expanded); got != test.want {
			t.Errorf("totalDistance(ex * %d) = %d, want %d", test.factor, got, test.want)
		}

	}
}
