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

package day09

import (
	"testing"
)

var ex1 = []move{
	{dirRight, 4},
	{dirUp, 4},
	{dirLeft, 3},
	{dirDown, 1},
	{dirRight, 4},
	{dirDown, 1},
	{dirLeft, 5},
	{dirRight, 2},
}

var ex2 = []move{
	{dirRight, 5},
	{dirUp, 8},
	{dirLeft, 8},
	{dirDown, 3},
	{dirRight, 17},
	{dirDown, 10},
	{dirLeft, 25},
	{dirUp, 20},
}

func TestMeasureTail(t *testing.T) {
	want := 13
	if got := measureTail(ex1); got != want {
		t.Errorf("measureTail(ex1) = %d, want %d", got, want)
	}
}

func TestMeasureLongTail(t *testing.T) {
	tests := []struct {
		name  string
		moves []move
		want  int
	}{
		{"ex1", ex1, 1},
		{"ex2", ex2, 36},
	}
	for _, test := range tests {
		if got := measureLongTail(test.moves); got != test.want {
			t.Errorf("measureLongTail(%s) = %d, want %d", test.name, got, test.want)
		}
	}
}
