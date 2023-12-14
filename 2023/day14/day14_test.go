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

package day14

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var ex = strings.TrimPrefix(`
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`, "\n")

func TestTotalLoad(t *testing.T) {
	l := parseLevel([]byte(ex))
	slideNorth(l)
	want := 136
	if got := totalLoad(l); got != want {
		t.Errorf("totalLoad(slideNorth(ex)) = %d, want %d", got, want)
	}
}

func TestSlideNorth(t *testing.T) {
	want := parseLevel([]byte(strings.TrimPrefix(`
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....
`, "\n")))
	got := parseLevel([]byte(ex))
	slideNorth(got)
	if diff := cmp.Diff(printLevel(want), printLevel(got)); diff != "" {
		t.Errorf("slideNorth(ex) mismatch (-want +got):\n%s", diff)
	}
}

func printLevel(l *level) string {
	var buf strings.Builder
	for y := 0; y < l.h; y++ {
		fmt.Fprintf(&buf, "%s\n", l.row(y))
	}
	return buf.String()
}
