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

package day07

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestCountAncestors(t *testing.T) {
	bag := "shiny gold"
	want := 4
	if g, err := parseRules(util.Lines(ex1)); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if got := countAncestors(g, bag); got != want {
		t.Errorf("countAncestors(%s) = %d, want %d", bag, got, want)
	}
}

func TestCountDescendants(t *testing.T) {
	bag := "shiny gold"
	tests := []struct {
		name  string
		rules string
		want  int
	}{
		{"ex1", ex1, 32},
		{"ex2", ex2, 126},
	}
	for _, test := range tests {
		if g, err := parseRules(util.Lines(test.rules)); err != nil {
			t.Errorf("parseRules(%s): %v", test.name, err)
		} else if got := countDescendants(g, bag); got != test.want {
			t.Errorf("countDescendants(%s, %s) = %d, want %d", test.name, bag, got, test.want)
		}
	}
}
