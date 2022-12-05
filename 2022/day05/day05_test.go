// Copyright 2022 Google LLC
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

package day05

import (
	"testing"
)

var ex = []string{
	"    [D]    ",
	"[N] [C]    ",
	"[Z] [M] [P]",
	" 1   2   3 ",
	"",
	"move 1 from 2 to 1",
	"move 3 from 1 to 3",
	"move 2 from 2 to 1",
	"move 1 from 1 to 2",
}

func TestApplyMoves9000(t *testing.T) {
	stacks, moves := parseInput(ex)
	want := "CMZ"
	stacks = applyMoves(stacks, moves, move.apply9000)
	if got := tops(stacks); got != want {
		t.Errorf("applyMoves(ex) -> %s, want %s", got, want)
	}
}

func TestApplyMoves9001(t *testing.T) {
	stacks, moves := parseInput(ex)
	want := "MCD"
	stacks = applyMoves(stacks, moves, move.apply9001)
	if got := tops(stacks); got != want {
		t.Errorf("applyMoves(ex) -> %s, want %s", got, want)
	}
}
