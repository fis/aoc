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

package day08

import (
	"testing"
)

func TestMaxRegs(t *testing.T) {
	code := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	wantFinal, wantEver := 1, 10
	if gotFinal, gotEver, err := maxRegs(code); err != nil {
		t.Errorf("part1: %v", err)
	} else if gotFinal != wantFinal || gotEver != wantEver {
		t.Errorf("part1 = (%d, %d), want (%d, %d)", gotFinal, gotEver, wantFinal, wantEver)
	}
}
