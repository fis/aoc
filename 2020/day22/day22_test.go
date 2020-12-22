// Copyright 2020 Google LLC
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

import "testing"

var (
	deck1 = []string{"Player 1:", "9", "2", "6", "3", "1"}
	deck2 = []string{"Player 2:", "5", "8", "4", "7", "10"}
)

func TestGame(t *testing.T) {
	p1, _ := parseDeck(1, deck1)
	p2, _ := parseDeck(2, deck2)
	want := 306
	got := game(p1, p2)
	if got != want {
		t.Errorf("game = %d, want %d", got, want)
	}
}

func TestRecursive(t *testing.T) {
	p1, _ := parseDeck(1, deck1)
	p2, _ := parseDeck(2, deck2)
	want := -291
	got := recursive(p1, p2, map[uint64]int{})
	if got != want {
		t.Errorf("recursive = %d, want %d", got, want)
	}
}
