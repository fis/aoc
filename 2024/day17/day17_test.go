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

package day17

import (
	"strings"
	"testing"
)

var ex = strings.TrimPrefix(`
Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`, "\n")

func TestInterp(t *testing.T) {
	regs, prog := parseInput(strings.NewReader(ex))
	want := "4,6,3,5,6,3,5,2,1,0"
	if got := interp(regs, prog); got != want {
		t.Errorf("interp(ex) = %q, want %q", got, want)
	}
}
