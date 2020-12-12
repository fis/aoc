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
)

var example = [][]string{
	{"abc"},
	{"a", "b", "c"},
	{"ab", "ac"},
	{"a", "a", "a", "a"},
	{"b"},
}

func TestCountMerged(t *testing.T) {
	tests := []struct {
		name   string
		merger func([]answerSet) answerSet
		want   []int
	}{
		{"any", mergeAny, []int{3, 3, 3, 1, 1}},
		{"all", mergeAll, []int{3, 0, 1, 1, 1}},
	}
	for _, test := range tests {
		for i, group := range example {
			answers := parseGroup(group)
			got := countMerged(answers, test.merger)
			if got != test.want[i] {
				t.Errorf("countMerged(%v, %s) = %d, want %d", group, test.name, got, test.want[i])
			}
		}
	}
}
