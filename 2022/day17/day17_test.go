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

package day17

import (
	"testing"
)

var ex = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func TestDropRocks(t *testing.T) {
	want := 3068
	got := dropRocks(ex, 2022)
	if got != want {
		t.Errorf("dropRocks(ex, 2022) = %d, want %d", got, want)
	}
}

func TestAnalyzeRocks(t *testing.T) {
	want := 1514285714288
	got := analyzeRocks(ex, 1000000000000)
	if got != want {
		t.Errorf("analyzeRocks(ex, 1000000000000) = %d, want %d", got, want)
	}
}
