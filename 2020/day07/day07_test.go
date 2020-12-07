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
)

var ex1 = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

var ex2 = []string{
	"shiny gold bags contain 2 dark red bags.",
	"dark red bags contain 2 dark orange bags.",
	"dark orange bags contain 2 dark yellow bags.",
	"dark yellow bags contain 2 dark green bags.",
	"dark green bags contain 2 dark blue bags.",
	"dark blue bags contain 2 dark violet bags.",
	"dark violet bags contain no other bags.",
}

func TestCountAncestors(t *testing.T) {
	bag := "shiny gold"
	want := 4
	if g, _, err := parseRules(ex1); err != nil {
		t.Errorf("parseRules: %v", err)
	} else if got := countAncestors(g, bag); got != want {
		t.Errorf("countAncestors(%s) = %d, want %d", bag, got, want)
	}
}

func TestCountDescendants(t *testing.T) {
	bag := "shiny gold"
	tests := []struct {
		name  string
		rules []string
		want  int
	}{
		{"ex1", ex1, 32},
		{"ex2", ex2, 126},
	}
	for _, test := range tests {
		if g, w, err := parseRules(test.rules); err != nil {
			t.Errorf("parseRules(%s): %v", test.name, err)
		} else if got := countDescendants(g, w, bag); got != test.want {
			t.Errorf("countDescendants(%s, %s) = %d, want %d", test.name, bag, got, test.want)
		}
	}
}
