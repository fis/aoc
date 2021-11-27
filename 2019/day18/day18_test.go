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

package day18

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

const ex1 = `
#########
#b.A.@.a#
#########
`

const ex2 = `
########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`

const ex3 = `
########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`

const ex4 = `
#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`

const ex5 = `
########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`

const ex6 = `
#######
#a.#Cd#
##@#@##
#######
##@#@##
#cB#Ab#
#######
`

const ex7 = `
###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############
`

const ex8 = `
#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############
`

const ex9 = `
#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############
`

func TestSolver(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int
	}{
		{name: "ex1", data: ex1, want: 8},
		{name: "ex2", data: ex2, want: 86},
		{name: "ex3", data: ex3, want: 132},
		{name: "ex4", data: ex4, want: 136},
		{name: "ex5", data: ex5, want: 81},
		{name: "ex6", data: ex6, want: 8},
		{name: "ex7", data: ex7, want: 24},
		{name: "ex8", data: ex8, want: 32},
		{name: "ex9", data: ex9, want: 72},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(test.data, "\n"), '#')
		got := solveLevel(level)
		if got != test.want {
			t.Errorf("%s = %d, want %d", test.name, got, test.want)
		}
	}
}
