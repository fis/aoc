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

package day23

import "testing"

func TestInRange(t *testing.T) {
	bots := []nanobot{
		{p3{0, 0, 0}, 4},
		{p3{1, 0, 0}, 1},
		{p3{4, 0, 0}, 3},
		{p3{0, 2, 0}, 1},
		{p3{0, 5, 0}, 3},
		{p3{0, 0, 3}, 1},
		{p3{1, 1, 1}, 1},
		{p3{1, 1, 2}, 1},
		{p3{1, 3, 1}, 1},
	}
	want := 7
	got := inRange(bots)
	if got != want {
		t.Errorf("inRange(%v) = %d, want %d", bots, got, want)
	}
}

func TestBestPos(t *testing.T) {
	bots := []nanobot{
		{p3{10, 12, 12}, 2},
		{p3{12, 14, 12}, 2},
		{p3{16, 12, 12}, 4},
		{p3{14, 14, 14}, 6},
		{p3{50, 50, 50}, 200},
		{p3{10, 10, 10}, 5},
	}
	want := 36
	got := bestPos(bots)
	if got != want {
		t.Errorf("bestPos(%v) = %d, want %d", bots, got, want)
	}
}
