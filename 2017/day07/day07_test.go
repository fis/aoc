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

package day07

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindRoot(t *testing.T) {
	lines := []string{
		"pbga (66)",
		"xhth (57)",
		"ebii (61)",
		"havc (66)",
		"ktlj (57)",
		"fwft (72) -> ktlj, cntj, xhth",
		"qoyq (66)",
		"padx (45) -> pbga, havc, qoyq",
		"tknk (41) -> ugml, padx, fwft",
		"jptl (61)",
		"ugml (68) -> gyxo, ebii, jptl",
		"gyxo (61)",
		"cntj (57)",
	}
	wantRoot, wantWeight := "tknk", 60
	if progs, err := parseLines(lines); err != nil {
		t.Errorf("parseLines: %v", err)
	} else {
		g := buildGraph(progs)
		root := findRoot(g)
		if root != wantRoot {
			t.Errorf("findRoot = %s, want %s", root, wantRoot)
		} else if _, w := fixWeight(root, progs); w != wantWeight {
			t.Errorf("fixWeight = %d, want %d", w, wantWeight)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		line string
		want program
	}{
		{
			line: "pbga (66)",
			want: program{name: "pbga", weight: 66},
		},
		{
			line: "fwft (72) -> ktlj, cntj, xhth",
			want: program{name: "fwft", weight: 72, subNames: []string{"ktlj", "cntj", "xhth"}},
		},
	}
	for _, test := range tests {
		if got, err := parseLine(test.line); err != nil {
			t.Errorf("parseLine(%s): %v", test.line, err)
		} else if !cmp.Equal(got, test.want, cmp.AllowUnexported(program{})) {
			t.Errorf("parseLine(%s) = %#v, want %#v", test.line, got, test.want)
		}
	}
}
