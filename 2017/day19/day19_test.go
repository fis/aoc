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

package day19

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
    |
    |  +--+
    A  |  C
F---|----E|--+
    |  |  |  D
    +B-+  +--+
`

func TestTracePath(t *testing.T) {
	level := util.ParseLevelString(strings.TrimPrefix(ex, "\n"), ' ')
	want1, want2 := "ABCDEF", 38
	if got1, got2 := tracePath(level); got1 != want1 || got2 != want2 {
		t.Errorf("tracePath = (%s, %d), want (%s, %d)", got1, got2, want1, want2)
	}
}
