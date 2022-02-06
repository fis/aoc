// Copyright 2022 Google LLC
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

package day01

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		steps string
		want  int
	}{
		{steps: "R2, L3", want: 5},
		{steps: "R2, R2, R2", want: 2},
		{steps: "R5, L5, R5, R3", want: 12},
	}
	for _, test := range tests {
		steps := strings.Split(test.steps, ", ")
		got, _ := find(steps)
		if got != test.want {
			t.Errorf("find(%v) = (%d, ?), want (%d, ?)", steps, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	steps := []string{"R8", "R4", "R4", "R8"}
	want := 4
	_, got := find(steps)
	if got != want {
		t.Errorf("find(%v) = (?, %d), want (?, %d)", steps, got, want)
	}
}
