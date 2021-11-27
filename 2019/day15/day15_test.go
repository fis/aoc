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

package day15

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

const ex = `
 ##
#..##
#.#..#
#.O.#
 ###
`

func TestExplore(t *testing.T) {
	real := util.ParseLevelStringAt(strings.TrimPrefix(ex, "\n"), ' ', util.P{-3, -2})
	dr := fakeDroid{level: real}
	explored := explore(&dr)

	realMap := strings.Join(real.Lines(real.Bounds()), "\n")
	exploredMap := strings.Join(explored.Lines(explored.Bounds()), "\n")

	if !cmp.Equal(exploredMap, realMap) {
		t.Errorf("map mismatch\nexplored:\n%s\nreal:\n%s", exploredMap, realMap)
	}
}

func TestDistance(t *testing.T) {
	level := util.ParseLevelStringAt(strings.TrimPrefix(ex, "\n"), ' ', util.P{-3, -2})
	move, target := distance(level, util.P{0, 0}, 'O')
	fill, _ := distance(level, util.P{-1, 1}, ' ')
	if move != 2 || fill != 4 || target != (util.P{-1, 1}) {
		t.Errorf("got %d, %d, %v; want 2, 4, {-1 1}", move, fill, target)
	}
}

type fakeDroid struct {
	level *util.Level
	pos   util.P
}

func (dr *fakeDroid) tryMove(to util.P) byte {
	tile := dr.level.At(to.X, to.Y)
	if tile != '#' {
		dr.pos = to
	}
	return tile
}

func (dr *fakeDroid) mustMove(to util.P) {
	dr.pos = to
}
