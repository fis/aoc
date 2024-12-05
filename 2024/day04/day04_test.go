// Copyright 2024 Google LLC
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

package day04

import "testing"

var ex = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestCountMatches(t *testing.T) {
	want := 18
	if got := countMatches(ex, "XMAS"); got != want {
		t.Errorf("countMatches(ex, XMAS) = %d, want %d", got, want)
	}
}

func TestCountXmas(t *testing.T) {
	want := 9
	if got := countXmas(ex); got != want {
		t.Errorf("countXmas(ex) = %d, want %d", got, want)
	}
}
