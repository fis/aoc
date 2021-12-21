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

package day21

import "testing"

func TestPart1(t *testing.T) {
	want := 739785
	if got := part1(4, 8); got != want {
		t.Errorf("part1(4, 8) = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 444356092776315
	if got := part2(4, 8); got != want {
		t.Errorf("part2(4, 8) = %d, want %d", got, want)
	}
}
