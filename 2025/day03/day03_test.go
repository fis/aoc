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

package day03

import "testing"

var ex = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestMaxJoltage2(t *testing.T) {
	want := []int{98, 89, 78, 92}
	for i, bank := range ex {
		if got := maxJoltage2(bank); got != want[i] {
			t.Errorf("maxJoltage(%s) = %d, want %d", bank, got, want[i])
		}
	}
}

func TestMaxJoltageN(t *testing.T) {
	want := []int{987654321111, 811111111119, 434234234278, 888911112111}
	for i, bank := range ex {
		if got := maxJoltageN(bank, 12); got != want[i] {
			t.Errorf("maxJoltage(%s, 12) = %d, want %d", bank, got, want[i])
		}
	}
}
