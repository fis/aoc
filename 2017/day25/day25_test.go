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

package day25

import (
	"strings"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = `
Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
`

func TestRun(t *testing.T) {
	blocks, err := util.ScanAll(strings.NewReader(strings.TrimPrefix(ex, "\n")), util.ScanChunks)
	if err != nil {
		t.Fatal(err)
	}
	tm, steps, err := parseTM(blocks)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	tm.run(steps)
	if got := tm.tape.popCount(); got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}
