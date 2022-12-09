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

package day08

import (
	"testing"
)

var ex = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestCountVisible(t *testing.T) {
	want := 21
	if got := countVisible(ex); got != want {
		t.Errorf("countVisible(ex) = %d, want %d", got, want)
	}
}

func TestFindBest(t *testing.T) {
	want := 8
	if got := findBest(ex); got != want {
		t.Errorf("findBest(ex) = %d, want %d", got, want)
	}
}
