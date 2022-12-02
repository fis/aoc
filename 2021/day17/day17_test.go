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

package day17

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestFindShots(t *testing.T) {
	Tmin, Tmax := util.P{20, -10}, util.P{30, -5}
	want1, want2 := 45, 112
	if got1, got2 := findShots(Tmin, Tmax); got1 != want1 || got2 != want2 {
		t.Errorf("bestShot(%v, %v) = (%d, %d), want (%d, %d)", Tmin, Tmax, got1, got2, want1, want2)
	}
}
