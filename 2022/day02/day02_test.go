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

package day02

import (
	"testing"
)

func TestTotalScoreShapes(t *testing.T) {
	guide := [][2]shape{
		{rock, paper},
		{paper, rock},
		{scissors, scissors},
	}
	want := 15
	if got := totalScoreShapes(guide); got != want {
		t.Errorf("totalScoreShapes(%v) = %d, want %d", guide, got, want)
	}
}

func TestTotalScoreRounds(t *testing.T) {
	guide := []guideLine{
		{rock, draw},
		{paper, loss},
		{scissors, win},
	}
	want := 12
	if got := totalScoreRounds(guide); got != want {
		t.Errorf("totalScoreRounds(%v) = %d, want %d", guide, got, want)
	}
}
