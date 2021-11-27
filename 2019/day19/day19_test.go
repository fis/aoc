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

package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const ex1 = `
#.........
.#........
..##......
...###....
....###...
.....####.
......####
......####
.......###
........##
`

const ex2 = `
#.......................................
.#......................................
..##....................................
...###..................................
....###.................................
.....####...............................
......#####.............................
......######............................
.......#######..........................
........########........................
.........#########......................
..........#########.....................
...........##########...................
...........############.................
............############................
.............#############..............
..............##############............
...............###############..........
................###############.........
................#################.......
.................########OOOOOOOOOO.....
..................#######OOOOOOOOOO#....
...................######OOOOOOOOOO###..
....................#####OOOOOOOOOO#####
.....................####OOOOOOOOOO#####
.....................####OOOOOOOOOO#####
......................###OOOOOOOOOO#####
.......................##OOOOOOOOOO#####
........................#OOOOOOOOOO#####
.........................OOOOOOOOOO#####
..........................##############
..........................##############
...........................#############
............................############
.............................###########
`

func TestEx1(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex1, "\n"), '.')
	got := part1(10, levelProber(level))
	if got != 27 {
		t.Errorf("part1 = %d, wanted 27", got)
	}
}

func TestEx2(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex2, "\n"), '.')
	got := part2(5, 10, levelProber(level))
	if got != 250020 {
		t.Errorf("part2 = %d, wanted 250020", got)
	}
}

func levelProber(level *util.Level) func(x, y int) bool {
	return func(x, y int) bool {
		return level.At(x, y) != '.'
	}
}
