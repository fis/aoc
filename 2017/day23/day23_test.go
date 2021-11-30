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

package day23

import (
	"testing"
)

func TestPart1(t *testing.T) {
	n := 57
	want := 3025
	if got := part1(n); got != want {
		t.Errorf("part1(%d) = %d, want %d", n, got, want)
	}
}

func TestCountComposite(t *testing.T) {
	low, high, skip := 105700, 122700, 17
	want := 915
	if got := countComposite(low, high, skip); got != want {
		t.Errorf("countComposite(%d, %d, %d) = %d, want %d", low, high, skip, got, want)
	}
}
