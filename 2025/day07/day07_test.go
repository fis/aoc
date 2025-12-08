// Copyright 2025 Google LLC
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

package day07

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`, "\n")

func TestCountSplits(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 21
	if got := countSplits(level); got != want {
		t.Errorf("countSplits(ex) = %d, want %d", got, want)
	}
}

func TestCountPaths(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want1, want2 := 21, 40
	if got1, got2 := countPaths(level); got1 != want1 || got2 != want2 {
		t.Errorf("countPaths(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
