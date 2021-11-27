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

package day08

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestLoopCheck(t *testing.T) {
	want := 5
	if prog, err := parseCode(util.Lines(example)); err != nil {
		t.Errorf("parseCode: %v", err)
	} else if loop, acc := loopCheck(prog); !loop || acc != want {
		t.Errorf("loopCheck = %v, %d, want true, %d", loop, acc, want)
	}
}

func TestRepair(t *testing.T) {
	want := 8
	if prog, err := parseCode(util.Lines(example)); err != nil {
		t.Errorf("parseCode: %v", err)
	} else if got := repair(prog); got != want {
		t.Errorf("repair = %d, want %d", got, want)
	}
}
