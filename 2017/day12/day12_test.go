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

package day12

import (
	"testing"
)

func TestPartition(t *testing.T) {
	lines := []string{
		`0 <-> 2`,
		`1 <-> 1`,
		`2 <-> 0, 3, 4`,
		`3 <-> 2, 4`,
		`4 <-> 2, 3, 6`,
		`5 <-> 6`,
		`6 <-> 4, 5`,
	}
	want := 6
	if g, err := parseLines(lines); err != nil {
		t.Errorf("parseLines: %v", err)
	} else {
		vertGroup, groupVerts := partition(g)
		if got := len(groupVerts[vertGroup[g.V("0")]]); got != want {
			t.Errorf("part1 = %d, want %d", got, want)
		}
	}
}
