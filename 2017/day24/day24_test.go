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

package day24

import (
	"testing"
)

func TestBestBridge(t *testing.T) {
	components := []component{
		{0, 2}, {2, 2}, {2, 3}, {3, 4},
		{3, 5}, {0, 1}, {10, 1}, {9, 10},
	}
	want1, want2, want3 := 31, 19, 4

	inv := inventory{}
	for i := range components {
		inv.give(&components[i])
	}

	if got1, got2, got3 := bestBridge(0, &inv); got1 != want1 || got2 != want2 || got3 != want3 {
		t.Errorf("bestBridge = (%d, %d, %d), want (%d, %d, %d)", got1, got2, got3, want1, want2, want3)
	}
}
