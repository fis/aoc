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

package day18

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var example = `
.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.
`

func TestEvolve(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(example, "\n"), ' ')
	want := 1147
	got := value(evolve(level, 10))
	if got != want {
		t.Errorf("evolve(10) -> value %d, want %d", got, want)
	}
}
