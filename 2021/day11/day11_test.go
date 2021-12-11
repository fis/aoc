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

package day11

import (
	"testing"
)

var ex = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func TestSimulate(t *testing.T) {
	tests := []struct {
		steps int
		want  int
	}{
		{steps: 1, want: 0},
		{steps: 2, want: 35},
		{steps: 10, want: 204},
		{steps: 100, want: 1656},
	}
	for _, test := range tests {
		if g, err := newGrid(ex); err != nil {
			t.Errorf("newGrid: %v", err)
		} else if got := g.simulate(test.steps); got != test.want {
			t.Errorf("simulate(%d) = %d, want %d", test.steps, got, test.want)
		}
	}
}

func TestSimulateToSync(t *testing.T) {
	want := 195
	if g, err := newGrid(ex); err != nil {
		t.Errorf("newGrid: %v", err)
	} else if got := g.simulateToSync(); got != want {
		t.Errorf("simulateToSync = %d, want %d", got, want)
	}
}
