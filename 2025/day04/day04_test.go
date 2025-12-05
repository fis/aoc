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

package day04

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`, "\n")

func TestCountAccessible(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want := 13
	if got := countAccessible(level); got != want {
		t.Errorf("countAccessible(ex) = %d, want %d", got, want)
	}
}

func TestRemovePaper(t *testing.T) {
	level := util.ParseFixedLevel([]byte(ex))
	want1, want2 := 13, 43
	if got1, got2 := removePaper(level); got1 != want1 && got2 != want2 {
		t.Errorf("removePaper(ex) = (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
