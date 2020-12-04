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

package day17

import (
	"strings"
	"testing"

	"github.com/fis/aoc-go/util"
)

const ex = `
..#..........
..#..........
#######...###
#.#...#...#.#
#############
..#...#...#..
..#####...^..
`

func TestCrosses(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
	got := crosses(level)
	if got != 76 {
		t.Errorf("got %d, want 76", got)
	}
}
