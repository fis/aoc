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

package day06

import (
	"testing"
)

var ex = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestCountWins(t *testing.T) {
	tests := []struct {
		time, dist int
		want       int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
		{71530, 940200, 71503},
	}
	impls := []struct {
		name string
		f    func(t, d int) int
	}{
		{"countWins", countWins},
		{"countWinsBinarySearch", countWinsBinarySearch},
	}
	for _, test := range tests {
		for _, impl := range impls {
			if got := impl.f(test.time, test.dist); got != test.want {
				t.Errorf("%s(%d, %d) = %d, want %d", impl.name, test.time, test.dist, got, test.want)
			}
		}
	}
}

func TestJoinDigits(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{7, 15, 30}, want: 71530},
		{input: []int{9, 40, 200}, want: 940200},
	}
	for _, test := range tests {
		if got := joinDigits(test.input); got != test.want {
			t.Errorf("joinDigits(%v) = %d, want %d", test.input, got, test.want)
		}
	}
}
