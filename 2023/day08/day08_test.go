// Copyright 2023 Google LLC
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

package day08

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestCountSteps(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"ex1", ex1, 2},
		{"ex2", ex2, 6},
	}
	for _, test := range tests {
		if g, dirs, err := parseMaps(util.Chunks(test.input)); err != nil {
			t.Errorf("parseMaps(%s): %v", test.name, err)
		} else if got := countSteps(g, g.v("AAA"), g.v("ZZZ"), dirs); got != test.want {
			t.Errorf("countSteps(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}

func TestCountGhostSteps(t *testing.T) {
	want := 6
	if g, dirs, err := parseMaps(util.Chunks(ex3)); err != nil {
		t.Errorf("parseMaps(ex3): %v", err)
	} else if got := countGhostSteps(g, dirs); got != want {
		t.Errorf("countSteps(ex3) = %d, want %d", got, want)
	}
}
