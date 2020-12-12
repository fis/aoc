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

package day12

import (
	"testing"
)

var ex = []string{"F10", "N3", "F7", "R90", "F11"}

func TestTurtle(t *testing.T) {
	x := newTurtle()
	x.move(parseInput(ex))
	want := 25
	got := x.distance()
	if got != want {
		t.Errorf("distance = %d, want %d", got, want)
	}
}

func TestShip(t *testing.T) {
	x := newShip()
	x.move(parseInput(ex))
	want := 286
	got := x.distance()
	if got != want {
		t.Errorf("distance = %d, want %d", got, want)
	}
}
