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

package day12

import "testing"

var tests = []struct {
	row          string
	groups       []int
	ways         int
	unfoldedWays int
}{
	{"???.###", []int{1, 1, 3}, 1, 1},
	{".??..??...?##.", []int{1, 1, 3}, 4, 16384},
	{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1, 1},
	{"????.#...#...", []int{4, 1, 1}, 1, 16},
	{"????.######..#####.", []int{1, 6, 5}, 4, 2500},
	{"?###????????", []int{3, 2, 1}, 10, 506250},
}

func TestCountWays(t *testing.T) {
	for _, test := range tests {
		if got := countWays([]byte(test.row), test.groups); got != test.ways {
			t.Errorf("countWays(%q, %v) = %d, want %d", test.row, test.groups, got, test.ways)
		}
	}
}

func TestCountWaysUnfolded(t *testing.T) {
	for _, test := range tests {
		r := unfold([]record{{[]byte(test.row), test.groups}}, 5)[0]
		if got := countWays(r.row, r.groups); got != test.unfoldedWays {
			t.Errorf("countWays(%q*5, %v*5) = %d, want %d", test.row, test.groups, got, test.unfoldedWays)
		}
	}
}
