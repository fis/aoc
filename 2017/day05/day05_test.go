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

package day05

import (
	"testing"
)

func TestPart1(t *testing.T) {
	offsets := []int{0, 3, 0, 1, -3}
	want := 5
	if got := part1(offsets); got != want {
		t.Errorf("part1(%v) = %d, want %d", offsets, got, want)
	}
}

func TestPart2(t *testing.T) {
	offsets := []int{0, 3, 0, 1, -3}
	want := 10
	if got := part2(offsets); got != want {
		t.Errorf("part2(%v) = %d, want %d", offsets, got, want)
	}
}
