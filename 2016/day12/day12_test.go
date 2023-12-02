// Copyright 2023 Google LLC
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

package day12

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

func TestRunProgram(t *testing.T) {
	ex := []string{
		"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}
	prog, err := fn.MapE(ex, parseInst)
	if err != nil {
		t.Fatal(err)
	}
	want := 42
	if got := runProgram(prog, 0); got != want {
		t.Errorf("runProgram(ex) = %d, want %d", got, want)
	}
}
