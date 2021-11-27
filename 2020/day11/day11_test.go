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

package day11

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

func TestFixedPoint(t *testing.T) {
	tests := []struct {
		name      string
		mapper    func(*util.Level) [][8]int
		tolerance int
		want      int
	}{
		{name: "near", mapper: nearMap, tolerance: 4, want: 37},
		{name: "far", mapper: farMap, tolerance: 5, want: 26},
	}
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), '.')
	for _, test := range tests {
		got := fixedPoint(level, test.mapper, test.tolerance)
		if got != test.want {
			t.Errorf("fixedPoint(%s, %d) = %d, want %d", test.name, test.tolerance, got, test.want)
		}
	}
}
