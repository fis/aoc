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

package day18

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

var (
	ex1 = []string{"1,1,1", "2,1,1"}
	ex2 = []string{
		"2,2,2", "1,2,2", "3,2,2", "2,1,2", "2,3,2", "2,2,1", "2,2,3",
		"2,2,4", "2,2,6", "1,2,5", "3,2,5", "2,1,5", "2,3,5",
	}
)

func TestSurfaceAreas(t *testing.T) {
	tests := []struct {
		name         string
		lines        []string
		want1, want2 int
	}{
		{name: "ex1", lines: ex1, want1: 10, want2: 10},
		{name: "ex2", lines: ex2, want1: 64, want2: 58},
	}
	for _, test := range tests {
		if cubes, err := fn.MapE(test.lines, parseCube); err != nil {
			t.Errorf("parseCubes(%s): %v", test.name, err)
		} else if got1, got2 := surfaceAreas(cubes); got1 != test.want1 || got2 != test.want2 {
			t.Errorf("surfaceAreas(%s) = (%d, %d), want (%d, %d)", test.name, got1, got2, test.want1, test.want2)
		}
	}
}
