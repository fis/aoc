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

package day21

import "testing"

var example = []string{
	"mxmxvkd kfcds sqjhc nhms (contains dairy, fish)",
	"trh fvjkl sbzzf mxmxvkd (contains dairy)",
	"sqjhc fvjkl (contains soy)",
	"sqjhc mxmxvkd sbzzf (contains fish)",
}

func TestAnalyze(t *testing.T) {
	want1, want2 := 5, "mxmxvkd,sqjhc,fvjkl"
	if labels, err := parseInput(example); err != nil {
		t.Errorf("parseInput: %v", err)
	} else if got1, got2 := analyze(labels); got1 != want1 || got2 != want2 {
		t.Errorf("part1 = (%d,%q), want (%d,%q)", got1, got2, want1, want2)
	}
}
