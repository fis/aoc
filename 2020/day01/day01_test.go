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

package day01

import (
	"testing"
)

var input = []int{1721, 979, 366, 299, 675, 1456}

func TestPart1(t *testing.T) {
	want := 514579
	got := part1(input)
	if got != want {
		t.Errorf("part1 = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 241861950
	got := part2(input)
	if got != want {
		t.Errorf("part2 = %d, want %d", got, want)
	}
}
