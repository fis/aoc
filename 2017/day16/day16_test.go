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

package day16

import (
	"testing"
)

func TestPerform(t *testing.T) {
	dance := []string{"s1", "x3/4", "pe/b"}
	want := "baedc"
	if moves, err := parseMoves(dance); err != nil {
		t.Errorf("parseMoves(%v): %v", dance, err)
	} else if got := perform(5, moves); got != want {
		t.Errorf("perform = %s, want %s", got, want)
	}
}

func TestEncore(t *testing.T) {
	dance := []string{"s1", "x3/4", "pe/b"}
	moves, err := parseMoves(dance)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		times int
		want  string
	}{
		{times: 1, want: "baedc"},
		{times: 2, want: "ceadb"},
		{times: 1000000000, want: "abcde"},
	}
	for _, test := range tests {
		if got := encore(5, moves, test.times); got != test.want {
			t.Errorf("encore(..., %d) = %s, want %s", test.times, got, test.want)
		}
	}
}
