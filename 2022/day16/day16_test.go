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

package day16

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/fis/aoc/util/fn"
)

func TestFindOne(t *testing.T) {
	lines, err := util.ReadLines("../../testdata/2022/day16.txt")
	if err != nil {
		t.Fatal(err)
	}
	scan, _ := fn.MapE(lines, ParseValveScan)
	sum := Preprocess(scan)
	want := 1789 // 2496
	if got := findOne(sum, 30); got != want {
		t.Errorf("findOne(day16, 30) = %d, want %d", got, want)
	}
}

func TestFindTwo(t *testing.T) {
	lines, err := util.ReadLines("../../testdata/2022/day16.txt")
	if err != nil {
		t.Fatal(err)
	}
	scan, _ := fn.MapE(lines, ParseValveScan)
	sum := Preprocess(scan)
	want := 2496
	if got := findTwo(sum, 26); got != want {
		t.Errorf("findTwo(day16, 26) = %d, want %d", got, want)
	}
}
