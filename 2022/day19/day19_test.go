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

package day19

import (
	"testing"
)

var (
	ex1 = blueprint{4, 2, 3, 14, 2, 7}
	ex2 = blueprint{2, 3, 3, 8, 3, 12}
)

func TestMaxGeodes(t *testing.T) {
	tests := []struct {
		bp      blueprint
		maxTime int
		want    int
	}{
		{ex1, 24, 9},
		{ex2, 24, 12},
		{ex1, 32, 56},
		{ex2, 32, 62},
	}
	for _, test := range tests {
		if got := maxGeodes(test.bp, test.maxTime); got != test.want {
			t.Errorf("maxGeodes(%v, %d) = %d, want %d", test.bp, test.maxTime, got, test.want)
		}
	}
}
