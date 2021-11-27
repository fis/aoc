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

package day03

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

func TestParseClaim(t *testing.T) {
	tests := []struct {
		line string
		want claim
	}{
		{line: "#1 @ 1,3: 4x4", want: claim{ID: 1, Pos: util.P{1, 3}, Size: util.P{4, 4}}},
		{line: "#2 @ 3,1: 4x4", want: claim{ID: 2, Pos: util.P{3, 1}, Size: util.P{4, 4}}},
		{line: "#3 @ 5,5: 2x2", want: claim{ID: 3, Pos: util.P{5, 5}, Size: util.P{2, 2}}},
	}
	for _, test := range tests {
		if got, err := parseClaim(test.line); err != nil {
			t.Errorf("parseClaim(%q): %v", test.line, err)
		} else if !cmp.Equal(got, test.want) {
			t.Errorf("parseClaim(%q) = %v, want %v", test.line, got, test.want)
		}
	}
}

func TestTotalOverlap(t *testing.T) {
	claims := []claim{
		{ID: 1, Pos: util.P{1, 3}, Size: util.P{4, 4}},
		{ID: 2, Pos: util.P{3, 1}, Size: util.P{4, 4}},
		{ID: 3, Pos: util.P{5, 5}, Size: util.P{2, 2}},
	}
	wantOverlap, wantValid := 4, 3
	overlap, valid := totalOverlap(claims)
	if overlap != wantOverlap || valid != wantValid {
		t.Errorf("totalOverlap = (%d, %d), want (%d, %d)", overlap, valid, wantOverlap, wantValid)
	}
}
