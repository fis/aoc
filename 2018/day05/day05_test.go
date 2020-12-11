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

package day05

import (
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "aA", want: ""},
		{input: "abBA", want: ""},
		{input: "abAB", want: "abAB"},
		{input: "aabAAB", want: "aabAAB"},
		{input: "dabAcCaCBAcCcaDA", want: "dabCBAcaDA"},
	}
	for _, test := range tests {
		buf := []byte(test.input)
		buf = reduce(buf)
		got := string(buf)
		if got != test.want {
			t.Errorf("reduce(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestImprove(t *testing.T) {
	input := "dabAcCaCBAcCcaDA"
	want := 4
	got := improve(input)
	if got != want {
		t.Errorf("improve(%q) = %d, want %d", input, got, want)
	}
}
