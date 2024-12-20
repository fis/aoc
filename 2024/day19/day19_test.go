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

package day19

import (
	"testing"
)

var (
	towels  = []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	designs = []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}
)

func TestCountPossible(t *testing.T) {
	want1, want2 := 6, 16
	if got1, got2 := countPossible(designs, towels); got1 != want1 || got2 != want2 {
		t.Errorf("countPossible(ex) = (%d,%d), want (%d,%d)", got1, got2, want1, want2)
	}
}
