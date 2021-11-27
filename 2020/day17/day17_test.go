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
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var example = `
.#.
..#
###
`

func TestCycle3(t *testing.T) {
	state, min, max := loadLevel3(util.ParseLevelString(strings.TrimSpace(example), '.'))
	for i := 0; i < 6; i++ {
		state, min, max = cycle3(state, min, max)
	}
	want := 112
	got := len(state)
	if got != want {
		t.Errorf("cycle3*6 -> %d, want %d", got, want)
	}
}

func TestCycle4(t *testing.T) {
	state, min, max := loadLevel4(util.ParseLevelString(strings.TrimSpace(example), '.'))
	for i := 0; i < 6; i++ {
		state, min, max = cycle4(state, min, max)
	}
	want := 848
	got := len(state)
	if got != want {
		t.Errorf("cycle4*6 -> %d, want %d", got, want)
	}
}
