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

package day22

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
..#
#..
...
`

func TestSimulateSimple(t *testing.T) {
	tests := []struct {
		rounds int
		want   int
	}{
		{rounds: 7, want: 5},
		{rounds: 70, want: 41},
		{rounds: 10000, want: 5587},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
		if got := simulateSimple(level, test.rounds); got != test.want {
			t.Errorf("simulateSimple(..., %d) = %d, want %d", test.rounds, got, test.want)
		}
	}
}

func TestSimulateEvolved(t *testing.T) {
	tests := []struct {
		rounds int
		want   int
	}{
		{rounds: 7, want: 1},
		{rounds: 100, want: 26},
		{rounds: 10000000, want: 2511944},
	}
	for _, test := range tests {
		level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
		if got := simulateEvolved(level, test.rounds); got != test.want {
			t.Errorf("simulateEvolved(..., %d) = %d, want %d", test.rounds, got, test.want)
		}
	}
}
