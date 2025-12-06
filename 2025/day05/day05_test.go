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

package day05

import "testing"

var (
	exFresh       = []interval{{3, 5}, {10, 14}, {16, 20}, {12, 18}}
	exIngredients = []int{1, 5, 8, 11, 17, 32}
)

func TestCountFresh(t *testing.T) {
	want1, want2 := 3, 14
	if got1, got2 := countFresh(exFresh, exIngredients); got1 != want1 || got2 != want2 {
		t.Errorf("countFresh(%v, %v) = (%d, %d), want (%d, %d)", exFresh, exIngredients, got1, got2, want1, want2)
	}
}
