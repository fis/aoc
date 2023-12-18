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

// Package day18 solves AoC 2023 day 18.
package day18

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var ex = [][]string{
	{"R", "6", "70c710"},
	{"D", "5", "0dc571"},
	{"L", "2", "5713f0"},
	{"D", "2", "d2c081"},
	{"R", "2", "59c680"},
	{"D", "2", "411b91"},
	{"L", "5", "8ceee2"},
	{"U", "2", "caa173"},
	{"L", "1", "1b58a2"},
	{"U", "2", "caa171"},
	{"R", "2", "7807d2"},
	{"U", "3", "a77fa3"},
	{"L", "2", "015232"},
	{"U", "2", "7a21e3"},
}

func TestPolygonArea(t *testing.T) {
	tests := []struct {
		name string
		f    func(instruction) (direction, int)
		want int
	}{
		{"small", instruction.small, 62},
		{"big", instruction.big, 952408144115},
	}
	plan, err := fn.MapE(ex, parseDigPlan)
	if err != nil {
		t.Fatalf("parseDigPlan(ex): %v", err)
	}
	for _, test := range tests {
		if got := polygonArea(traceLoop(plan, test.f)); got != test.want {
			t.Errorf("polygonArea(ex, %s) = %d, want %d", test.name, got, test.want)
		}
	}
}
