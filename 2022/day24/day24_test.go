// Copyright 2022 Google LLC
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

package day24

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []string{
	"#.######",
	"#>>.<^<#",
	"#.<..<<#",
	"#>v.><>#",
	"#<^v^^>#",
	"######.#",
}

func TestFindPath(t *testing.T) {
	lvl := parseLevel(ex)
	tests := []struct {
		t0         int
		start, end util.P
		want       int
	}{
		{0, util.P{0, -1}, util.P{5, 4}, 18},
		{18, util.P{5, 4}, util.P{0, -1}, 23},
		{18 + 23, util.P{0, -1}, util.P{5, 4}, 13},
	}
	for _, test := range tests {
		if got := findPath(lvl, test.t0, test.start, test.end); got != test.want {
			t.Errorf("findPath(ex, %d, %v, %v) = %d, want %d", test.t0, test.start, test.end, got, test.want)
		}
	}
}
