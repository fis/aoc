// Copyright 2025 Google LLC
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
	"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}",
	"[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}",
	"[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
}

func TestConfigLights(t *testing.T) {
	machines, err := fn.MapE(ex, parseMachine)
	if err != nil {
		t.Fatalf("parseMachine(ex): %v", err)
	}
	want := []int{2, 3, 2}
	if got := fn.Map(machines, machine.configLights); !cmp.Equal(got, want) {
		t.Errorf("(ex).configLights = %v, want %v", got, want)
	}
}

func TestConfigJolts(t *testing.T) {
	machines, err := fn.MapE(ex, parseMachine)
	if err != nil {
		t.Fatalf("parseMachine(ex): %v", err)
	}
	want := []int{10, 12, 11}
	if got := fn.Map(machines, machine.configJolts); !cmp.Equal(got, want) {
		t.Errorf("(ex).configJolts = %v, want %v", got, want)
	}
}
