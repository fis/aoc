// Copyright 2021 Google LLC
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

package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		square int
		want   int
	}{
		{square: 1, want: 0},
		{square: 12, want: 3},
		{square: 23, want: 2},
		{square: 1024, want: 31},
	}
	for _, test := range tests {
		if got := part1(test.square); got != test.want {
			t.Errorf("part1(%d) = %d, want %d", test.square, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		limit int
		want  int
	}{
		{limit: 50, want: 54},
		{limit: 55, want: 57},
		{limit: 100, want: 122},
		{limit: 300, want: 304},
		{limit: 500, want: 747},
		{limit: 800, want: 806},
	}
	for _, test := range tests {
		if got := part2(test.limit); got != test.want {
			t.Errorf("part2(%d) = %d, want %d", test.limit, got, test.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := [][2]int{
		{0, 0},
		{1, 1}, {2, 1}, {3, 1},
		{4, 2}, {5, 2}, {8, 2},
		{9, 3}, {10, 3}, {15, 3},
		{16, 4}, {17, 4}, {24, 4},
		{25, 5},
		{9999, 99}, {10000, 100},
	}
	for _, test := range tests {
		if got := sqrt(test[0]); got != test[1] {
			t.Errorf("sqrt(%d) = %d, want %d", test[0], got, test[1])
		}
	}
}
