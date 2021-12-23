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

package day23

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var ex = []string{
	`#############`,
	`#...........#`,
	`###B#C#B#D###`,
	`  #A#D#C#A#`,
	`  #########`,
}

func TestShortestPath(t *testing.T) {
	want := 12521
	start := decodeState(ex)
	if got := shortestPath(start, sortedState); got != want {
		t.Errorf("shortestPath(%v, %v) = %d, want %d", start, sortedState, got, want)
	}
}

func TestShortestDeepPath(t *testing.T) {
	want := 44169
	start := convertState(decodeState(ex))
	if got := shortestDeepPath(start, sortedDeepState); got != want {
		t.Errorf("shortestDeepPath(%v, %v) = %d, want %d", start, sortedDeepState, got, want)
	}
}

func TestMoves(t *testing.T) {
	tests := map[string]struct {
		state amphiState
		want  []amphiPath
	}{
		"ex": {
			state: decodeState(ex),
			want: []amphiPath{
				// B out
				{state: 0x5047657640, energy: 20}, {state: 0x547657640, energy: 30},
				{state: 0x50047657640, energy: 20}, {state: 0x500047657640, energy: 40}, {state: 0x5000047657640, energy: 60}, {state: 0x50000047657640, energy: 80}, {state: 0x500000047657640, energy: 90},
				// C out
				{state: 0x60047657045, energy: 200}, {state: 0x6047657045, energy: 400}, {state: 0x647657045, energy: 500},
				{state: 0x600047657045, energy: 200}, {state: 0x6000047657045, energy: 400}, {state: 0x60000047657045, energy: 600}, {state: 0x600000047657045, energy: 700},
				// other B out
				{state: 0x500047607645, energy: 20}, {state: 0x50047607645, energy: 40}, {state: 0x5047607645, energy: 60}, {state: 0x547607645, energy: 70},
				{state: 0x5000047607645, energy: 20}, {state: 0x50000047607645, energy: 40}, {state: 0x500000047607645, energy: 50},
				// D out
				{state: 0x7000040657645, energy: 2000}, {state: 0x700040657645, energy: 4000}, {state: 0x70040657645, energy: 6000}, {state: 0x7040657645, energy: 8000}, {state: 0x740657645, energy: 9000},
				{state: 0x70000040657645, energy: 2000}, {state: 0x700000040657645, energy: 3000},
			},
		},
		"sorted": {
			state: sortedState,
			want:  nil,
		},
		"one missing, far left": {
			state: decodeState([]string{
				`#############`,
				`#C..........#`,
				`###A#B#.#D###`,
				`  #A#B#C#D#`,
				`  #########`,
			}),
			want: []amphiPath{{state: sortedState, energy: 700}},
		},
		"one missing, near left": {
			state: decodeState([]string{
				`#############`,
				`#.....C.....#`,
				`###A#B#.#D###`,
				`  #A#B#C#D#`,
				`  #########`,
			}),
			want: []amphiPath{{state: sortedState, energy: 200}},
		},
		"one missing, near right": {
			state: decodeState([]string{
				`#############`,
				`#...A.......#`,
				`###.#B#C#D###`,
				`  #A#B#C#D#`,
				`  #########`,
			}),
			want: []amphiPath{{state: sortedState, energy: 2}},
		},
		"one missing, far right": {
			state: decodeState([]string{
				`#############`,
				`#..........A#`,
				`###.#B#C#D###`,
				`  #A#B#C#D#`,
				`  #########`,
			}),
			want: []amphiPath{{state: sortedState, energy: 9}},
		},
		"deadlocked": {
			state: decodeState([]string{
				`#############`,
				`#...C.A.....#`,
				`###.#B#.#D###`,
				`  #A#B#C#D#`,
				`  #########`,
			}),
			want: nil,
		}}
	for name, test := range tests {
		if got := test.state.moves(nil); !cmp.Equal(got, test.want, cmp.AllowUnexported(amphiPath{})) {
			t.Errorf("[%s] (%v).moves:\n  %v;\nwant:\n  %v", name, test.state, got, test.want)
		}
	}
}
