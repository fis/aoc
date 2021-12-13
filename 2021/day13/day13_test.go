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

package day13

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

var ex = struct {
	points []util.P
	folds  []foldSpec
}{
	points: []util.P{
		{6, 10}, {0, 14}, {9, 10}, {0, 3}, {10, 4}, {4, 11}, {6, 0}, {6, 12}, {4, 1},
		{0, 13}, {10, 12}, {3, 4}, {3, 0}, {8, 4}, {1, 10}, {2, 14}, {8, 10}, {9, 0},
	},
	folds: []foldSpec{{vertical: false, at: 7}, {vertical: true, at: 5}},
}

func TestFold(t *testing.T) {
	points := append([]util.P(nil), ex.points...)

	want1 := 17
	ex.folds[0].apply(points)
	if got1 := countVisible(points); got1 != want1 {
		t.Errorf("visible after first fold = %d, want %d", got1, want1)
	}

	want2 := []string{`#####`, `#   #`, `#   #`, `#   #`, `#####`}
	ex.folds[1].apply(points)
	if got2 := printPoints(points); !cmp.Equal(got2, want2) {
		t.Errorf("points after folding = %v, want %v", got2, want2)
	}
}
