// Copyright 2023 Google LLC
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
	"strings"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

var ex = strings.TrimPrefix(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`, "\n")

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		symFilter byte
		f         func([]int) int
		ft        string
		want      int
	}{
		{
			symFilter: 0,
			f:         fn.Sum[[]int], ft: "sum",
			want: 4361,
		},
		{
			symFilter: '*',
			f:         gearRatio, ft: "gearRatio",
			want: 467835,
		},
	}
	lv := util.ParseLevelString(ex, '.')
	for _, test := range tests {
		if got := sumNumbers(lv, test.symFilter, test.f); got != test.want {
			t.Errorf("sumNumbers(ex, %d, %s) = %d, want %d", test.symFilter, test.ft, got, test.want)
		}
	}
}
