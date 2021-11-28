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

package day09

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `{}`, want: 1},
		{input: `{{{}}}`, want: 6},
		{input: `{{},{}}`, want: 5},
		{input: `{{{},{},{{}}}}`, want: 16},
		{input: `{<a>,<a>,<a>,<a>}`, want: 1},
		{input: `{{<ab>},{<ab>},{<ab>},{<ab>}}`, want: 9},
		{input: `{{<!!>},{<!!>},{<!!>},{<!!>}}`, want: 9},
		{input: `{{<a!>},{<a!>},{<a!>},{<ab>}}`, want: 3},
	}
	for _, test := range tests {
		if groups, err := parseStream([]byte(test.input)); err != nil {
			t.Errorf("parseStream(%q): %v", test.input, err)
		} else if got, _ := score(groups, 1); got != test.want {
			t.Errorf("score(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestParseGarbage(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `<>`, want: 0},
		{input: `<random characters>`, want: 17},
		{input: `<<<<>`, want: 3},
		{input: `<{!>}>`, want: 2},
		{input: `<!!>`, want: 0},
		{input: `<!!!>>`, want: 0},
		{input: `<{o"i!a,<{i<a>`, want: 10},
	}
	for _, test := range tests {
		if tail, got, err := parseGarbage([]byte(test.input)); err != nil {
			t.Errorf("parseGarbage(%q): %v", test.input, err)
		} else if len(tail) > 0 {
			t.Errorf("parseGarbage(%q): leftover tail: %q", test.input, tail)
		} else if got != test.want {
			t.Errorf("parseGarbage(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}
