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

// Package day17 solves AoC 2023 day 17.
package day17

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var (
	ex = strings.TrimPrefix(`
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
`, "\n")
	ex2 = strings.TrimPrefix(`
111111111111
999999999991
999999999991
999999999991
999999999991
`, "\n")
)

func TestHeatLoss(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	want := 102
	if got := heatLoss(l); got != want {
		t.Errorf("heatLoss(ex) = %d, want %d", got, want)
	}
}

func TestUltraHeatLoss(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ex", ex, 94},
		{"ex2", ex2, 71},
	}
	for _, test := range tests {
		l := util.ParseFixedLevel([]byte(test.input))
		if got := ultraHeatLoss(l); got != test.want {
			t.Errorf("ultraHeatLoss(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}
