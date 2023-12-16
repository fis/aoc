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

package day16

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = strings.TrimPrefix(`
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`, "\n")

func TestCountEnergized(t *testing.T) {
	tests := []struct {
		x, y, dx, dy, di int
		want             int
	}{
		{0, 0, 1, 0, 0, 46},
		{3, 0, 0, 1, 1, 51},
	}
	l := util.ParseFixedLevel([]byte(ex))
	for _, test := range tests {
		if got := countEnergized(l, test.x, test.y, test.dx, test.dy, test.di); got != test.want {
			t.Errorf("countEnergized(ex, %d, %d, ...) = %d, want %d", test.x, test.y, got, test.want)
		}
	}
}
