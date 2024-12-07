// Copyright 2024 Google LLC
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

	"github.com/fis/aoc/util/fn"
)

var ex = []string{
	"190: 10 19",
	"3267: 81 40 27",
	"83: 17 5",
	"156: 15 6",
	"7290: 6 8 6 15",
	"161011: 16 10 13",
	"192: 17 8 14",
	"21037: 9 7 18 13",
	"292: 11 6 16 20",
}

func TestSumPossible(t *testing.T) {
	eqs, _ := fn.MapE(ex, parseEquation)
	want := 3749
	if got := sumPossible(eqs); got != want {
		t.Errorf("sumPossible(ex) = %d, want %d", got, want)
	}
}

func TestSumPossible3(t *testing.T) {
	eqs, _ := fn.MapE(ex, parseEquation)
	want := 11387
	if got := sumPossible3(eqs); got != want {
		t.Errorf("sumPossible3(ex) = %d, want %d", got, want)
	}
}
