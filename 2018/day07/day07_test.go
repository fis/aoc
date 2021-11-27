// Copyright 2020 Google LLC
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

package day07

import (
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

func TestToplexSort(t *testing.T) {
	g := parseRules(util.Lines(example))
	want := []string{"C", "A", "B", "D", "F", "E"}
	got := toplexSort(g)
	if !cmp.Equal(got, want) {
		t.Errorf("topLexSort = %v, want %v", got, want)
	}
}

func TestTimedSort(t *testing.T) {
	g := parseRules(util.Lines(example))
	wantOrder, wantTime := []string{"C", "A", "B", "F", "D", "E"}, 15
	gotOrder, gotTime := timedSort(g, 2, 0)
	if !cmp.Equal(gotOrder, wantOrder) || gotTime != wantTime {
		t.Errorf("timedSort = (%v, %d), want (%v, %d)", gotOrder, gotTime, wantOrder, wantTime)
	}
}
