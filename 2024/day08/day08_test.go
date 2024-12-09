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

package day08

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`, "\n")

func TestCountAntinodes(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	antennae := findAntennae(level)
	want := 14
	if got := countAntinodes(level, antennae); got != want {
		t.Errorf("countAntinodes(ex) = %d, want %d", got, want)
	}
}

func TestCountAllAntinodes(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	antennae := findAntennae(level)
	want := 34
	if got := countAllAntinodes(level, antennae); got != want {
		t.Errorf("countAllAntinodes(ex) = %d, want %d", got, want)
	}
}
