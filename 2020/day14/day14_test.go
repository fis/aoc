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

package day14

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	ex1 = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}
	ex2 = []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}
)

func TestEvaluate1(t *testing.T) {
	code, _ := parseCode(ex1)
	want := uint(165)
	got := sumValues(evaluate1(code))
	if got != want {
		t.Errorf("sum(evaluate1) = %d, want %d", got, want)
	}
}

func TestEvaluate2(t *testing.T) {
	code, _ := parseCode(ex2)
	want := uint(208)
	got := sumValues(evaluate2(code))
	if got != want {
		t.Errorf("sum(evaluate2) = %d, want %d", got, want)
	}
}

func TestMakeFloatBits(t *testing.T) {
	mask := uint(0b10011)
	want := []uint{0b00000, 0b00001, 0b00010, 0b00011, 0b10000, 0b10001, 0b10010, 0b10011}
	got := makeFloatBits(mask, []uint(nil))
	if !cmp.Equal(got, want) {
		t.Errorf("makeFloatBits = %v, want %v", got, want)
	}
}
