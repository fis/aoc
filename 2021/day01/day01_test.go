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

package day01

import (
	"testing"
)

func TestIncreases(t *testing.T) {
	depths := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want1, want2 := 7, 5
	if got1, got2 := increases(depths); got1 != want1 || got2 != want2 {
		t.Errorf("increases(%v) = (%d, %d), want (%d, %d)", depths, got1, got2, want1, want2)
	}
}
