// Copyright 2019 Google LLC
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

package day24

import (
	"strings"
	"testing"
)

var ex = strings.Split(strings.TrimSpace(`
....#
#..#.
#..##
..#..
#....
`), "\n")

func TestFindRepeating(t *testing.T) {
	initial := parseState(ex)
	want := state(2129920)
	got := findRepeating(initial)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCountBugs(t *testing.T) {
	initial := parseState(ex)
	got := countBugs(initial, 10)
	if got != 99 {
		t.Errorf("got %d, want 99", got)
	}
}
