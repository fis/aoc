// Copyright 2024 Google LLC
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

package day03

import (
	"testing"
)

var (
	ex1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
	ex2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
)

func TestFindInstructions(t *testing.T) {
	want1, want2 := 161, 48
	got1, _ := findInstructions(ex1)
	got2, _ := findInstructions(ex2)
	if got1 != want1 {
		t.Errorf("findInstructions(ex1) = %d, _, want %d", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("findInstructions(ex2) = _, %d, want %d", got2, want2)
	}
}
