// Copyright 2020 Google LLC
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

package day20

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

var (
	ex1     = `^WNE$`
	layout1 = `
#####
#.|.#
#-###
#.|X#
#####
`
	ex2     = `^ENWWW(NEEE|SSE(EE|N))$`
	layout2 = `
#########
#.|.|.|.#
#-#######
#.|.|.|.#
#-#####-#
#.#.#X|.#
#-#-#####
#.|.|.|.#
#########
`
	ex3     = `^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$`
	layout3 = `
###########
#.|.#.|.#.#
#-###-#-#-#
#.|.|.#.#.#
#-#####-#-#
#.#.#X|.#.#
#-#-#####-#
#.#.|.|.|.#
#-###-###-#
#.|.|.#.|.#
###########
`
	ex4     = `^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$`
	layout4 = `
#############
#.|.|.|.|.|.#
#-#####-###-#
#.#.|.#.#.#.#
#-#-###-#-#-#
#.#.#.|.#.|.#
#-#-#-#####-#
#.#.#.#X|.#.#
#-#-#-###-#-#
#.|.#.|.#.#.#
###-#-###-#-#
#.|.#.|.|.#.#
#############
`
	ex5     = `^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$`
	layout5 = `
###############
#.|.|.|.#.|.|.#
#-###-###-#-#-#
#.|.#.|.|.#.#.#
#-#########-#-#
#.#.|.|.|.|.#.#
#-#-#########-#
#.#.#.|X#.|.#.#
###-#-###-#-#-#
#.|.#.#.|.#.|.#
#-###-#####-###
#.|.#.|.|.#.#.#
#-#-#####-#-#-#
#.#.|.|.|.#.|.#
###############
`
)

func TestTrace(t *testing.T) {
	tests := []struct {
		ex     string
		layout string
	}{
		{ex: ex1, layout: layout1},
		{ex: ex2, layout: layout2},
		{ex: ex3, layout: layout3},
		{ex: ex4, layout: layout4},
		{ex: ex5, layout: layout5},
	}
	for _, test := range tests {
		ex, err := parseDirex(test.ex)
		if err != nil {
			t.Errorf("parseDirex(%s): %v", test.ex, err)
			continue
		}
		l := make(layout)
		ex.trace([]util.P{{0, 0}}, l)
		want := strings.TrimPrefix(test.layout, "\n")
		got := l.String()
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("(%s).trace mismatch (-want +got):\n%s", test.ex, diff)
		}
	}
}

func TestRadius(t *testing.T) {
	tests := []struct {
		ex   string
		want int
	}{
		{ex: ex1, want: 3},
		{ex: ex2, want: 10},
		{ex: ex3, want: 18},
		{ex: ex4, want: 23},
		{ex: ex5, want: 31},
	}
	for _, test := range tests {
		ex, err := parseDirex(test.ex)
		if err != nil {
			t.Errorf("parseDirex(%s): %v", test.ex, err)
			continue
		}
		l := make(layout)
		ex.trace([]util.P{{0, 0}}, l)
		got, _ := l.radius()
		if got != test.want {
			t.Errorf("(%s).radius() = %d, want %d", test.ex, got, test.want)
		}
	}
}
