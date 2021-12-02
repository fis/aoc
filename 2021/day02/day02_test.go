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

package day02

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = [][]string{
	{"forward", "5"},
	{"down", "5"},
	{"forward", "8"},
	{"up", "3"},
	{"down", "8"},
	{"forward", "2"},
}

func TestApplyMoves(t *testing.T) {
	want := util.P{15, 10}
	if got := applyMoves(util.P{0, 0}, parseInput(ex)); got != want {
		t.Errorf("applyMoves(0, %v) = %d, want %d", ex, got, want)
	}
}

func TestApplyMoves2(t *testing.T) {
	want := util.P{15, 60}
	if got := applyMoves2(util.P{0, 0}, parseInput(ex)); got != want {
		t.Errorf("applyMoves2(0, %v) = %d, want %d", ex, got, want)
	}
}
