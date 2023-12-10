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

package day10

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var (
	ex1 = strings.TrimPrefix(`
-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`, "\n")
	ex2 = strings.TrimPrefix(`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`, "\n")
	ex3 = strings.TrimPrefix(`
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`, "\n")
	ex4 = strings.TrimPrefix(`
.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...
`, "\n")
	ex5 = strings.TrimPrefix(`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`, "\n")
)

func TestTraceLoop(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		maxDist int
	}{
		{"ex1", ex1, 4},
		{"ex2", ex2, 8},
	}
	for _, test := range tests {
		l := util.ParseLevelString(test.input, '.')
		lp := traceLoop(l)
		if got := len(lp.tiles) / 2; got != test.maxDist {
			t.Errorf("traceLoop(%s) = %d, want %d", test.name, got, test.maxDist)
		}
	}
}

func TestEnclosedArea(t *testing.T) {
	tests := []struct {
		name  string
		input string
		area  int
	}{
		{"ex3", ex3, 4},
		{"ex4", ex4, 8},
		{"ex5", ex5, 10},
	}
	for _, test := range tests {
		l := util.ParseLevelString(test.input, '.')
		lp := traceLoop(l)
		if got := enclosedArea(l, lp); got != test.area {
			t.Errorf("enclosedArea(%s) = %d, want %d", test.name, got, test.area)
		}
	}
}
