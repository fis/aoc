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

package day10

import (
	"testing"

	"github.com/fis/aoc/util/fn"
	"github.com/google/go-cmp/cmp"
)

var ex = []string{
	"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8", "addx 13", "addx 4",
	"noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1", "addx 5",
	"addx -1", "addx -35", "addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop",
	"noop", "addx 21", "addx -15", "noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3",
	"addx 8", "addx 1", "addx 5", "noop", "noop", "noop", "noop", "noop", "addx -36", "noop",
	"addx 1", "addx 7", "noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop",
	"noop", "addx 1", "noop", "noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7",
	"noop", "addx 1", "addx -33", "noop", "noop", "noop", "addx 2", "noop", "noop", "noop", "addx 8",
	"noop", "addx -1", "addx 2", "addx 1", "noop", "addx 17", "addx -9", "addx 1", "addx 1",
	"addx -3", "addx 11", "noop", "noop", "addx 1", "noop", "addx 1", "noop", "noop", "addx -13",
	"addx -19", "addx 1", "addx 3", "addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1",
	"noop", "noop", "noop", "addx -9", "addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9",
	"noop", "noop", "noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15",
	"addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop", "addx -10",
	"noop", "noop", "addx 20", "addx 1", "addx 2", "addx 2", "addx -6", "addx -11", "noop", "noop",
	"noop",
}

func TestSigStrength(t *testing.T) {
	prog, err := fn.MapE(ex, parseInstruction)
	if err != nil {
		t.Fatal(err)
	}
	want := 13140
	if got := sigStrength(prog); got != want {
		t.Errorf("sigStrength(ex) = %d, want %d", got, want)
	}
}

func TestRender(t *testing.T) {
	prog, err := fn.MapE(ex, parseInstruction)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"##  ##  ##  ##  ##  ##  ##  ##  ##  ##  ",
		"###   ###   ###   ###   ###   ###   ### ",
		"####    ####    ####    ####    ####    ",
		"#####     #####     #####     #####     ",
		"######      ######      ######      ####",
		"#######       #######       #######     ",
	}
	got := render(prog)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("render(ex) mismatch (-want +got):\n%s", diff)
	}
}
