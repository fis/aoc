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

package day01

import (
	"testing"
)

func TestMaxCalories(t *testing.T) {
	data := [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
	want1, want2 := 24000, 45000
	if got1, got2 := maxCalories(data); got1 != want1 || got2 != want2 {
		t.Errorf("maxCalories(%v) = (%d, %d), want (%d, %d)", data, got1, got2, want1, want2)
	}
}
