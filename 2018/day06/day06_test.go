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

package day06

import (
	"testing"

	"github.com/fis/aoc/util"
)

var example = []util.P{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}}

func TestMaxArea(t *testing.T) {
	want := 17
	got := maxArea(example)
	if got != want {
		t.Errorf("maxArea = %d, want %d", got, want)
	}
}

func TestSafeArea(t *testing.T) {
	want := 16
	got := safeArea(example, 32)
	if got != want {
		t.Errorf("safeArea = %d, want %d", got, want)
	}
}
