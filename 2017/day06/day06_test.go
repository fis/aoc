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

package day06

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLoopLength(t *testing.T) {
	banks := []byte{0, 2, 7, 0}
	want1, want2 := 5, 4
	if got := loopLength(banks); got != want1 {
		t.Errorf("loopLength (1st run) = %d, want %d", got, want1)
	}
	if got := loopLength(banks); got != want2 {
		t.Errorf("loopLength (2nd run) = %d, want %d", got, want2)
	}
}

func TestBalance(t *testing.T) {
	steps := [][]byte{
		{0, 2, 7, 0},
		{2, 4, 1, 2},
		{3, 1, 2, 3},
		{0, 2, 3, 4},
		{1, 3, 4, 1},
		{2, 4, 1, 2},
	}
	for i := 0; i < len(steps)-1; i++ {
		banks, want := steps[i], steps[i+1]
		got := append(banks[:0:0], banks...)
		balance(got)
		if !cmp.Equal(got, want) {
			t.Errorf("balance(%v) -> %v, want %v", banks, got, want)
		}
	}
}
