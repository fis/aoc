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

package day06

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`, "\n")

func TestWalkOff(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 41
	if got := walkOff(level); got != want {
		t.Errorf("walkOff(ex) = %d, want %d", got, want)
	}
}

func TestCountLoops(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 6
	if got := countLoops(level); got != want {
		t.Errorf("countLoops(ex) = %d, want %d", got, want)
	}

}
