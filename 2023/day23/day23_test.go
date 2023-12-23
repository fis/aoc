// Copyright 2023 Google LLC
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

	"github.com/fis/aoc/util"
)

func TestLongestPath(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	g, _, _ := deconstruct(l)
	want := 94
	if got := longestPath(g); got != want {
		t.Errorf("longestPath(ex) = %d, want %d", got, want)
	}
}

func TestEvenLongestPath(t *testing.T) {
	l := util.ParseFixedLevel([]byte(ex))
	g, startV, endV := deconstruct(l)
	g.MakeUndirected()
	want := 154
	if got := evenLongestPath(g, startV, endV); got != want {
		t.Errorf("evenLongestPath(ex) = %d, want %d", got, want)
	}
}
