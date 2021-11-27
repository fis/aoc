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

package day17

import (
	"testing"

	"github.com/fis/aoc/util"
)

var example = `x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
`

func TestFill(t *testing.T) {
	lines := util.Lines(example)
	level, minY, maxY := readScan(lines)
	fill(level, util.P{500, 0}, make(map[util.P]struct{}))
	want1, want2 := 57, 29
	got1, got2 := measureWater(level, minY, maxY)
	if got1 != want1 || got2 != want2 {
		t.Errorf("measureWater = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
