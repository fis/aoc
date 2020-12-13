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

package day09

import (
	"testing"
)

func TestGame(t *testing.T) {
	tests := []struct {
		players, marbles int
		want             int
	}{
		{players: 9, marbles: 25, want: 32},
		{players: 10, marbles: 1618, want: 8317},
		{players: 13, marbles: 7999, want: 146373},
		{players: 17, marbles: 1104, want: 2764},
		{players: 21, marbles: 6111, want: 54718},
		{players: 30, marbles: 5807, want: 37305},
	}
	for _, test := range tests {
		got := game(test.players, test.marbles)
		if got != test.want {
			t.Errorf("game(%d, %d) = %d, want %d", test.players, test.marbles, got, test.want)
		}
	}
}
