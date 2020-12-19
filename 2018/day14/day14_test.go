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

package day14

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEvolve(t *testing.T) {
	tests := []struct {
		skip int
		want []byte
	}{
		{skip: 5, want: []byte{0, 1, 2, 4, 5, 1, 5, 8, 9, 1}},
		{skip: 9, want: []byte{5, 1, 5, 8, 9, 1, 6, 7, 7, 9}},
		{skip: 18, want: []byte{9, 2, 5, 1, 0, 7, 1, 0, 8, 5}},
		{skip: 2018, want: []byte{5, 9, 4, 1, 4, 2, 9, 8, 8, 2}},
	}
	for _, test := range tests {
		keep := len(test.want)
		got := evolve(test.skip, keep)
		if !cmp.Equal(got, test.want) {
			t.Errorf("evolve(%d, %d) = %v, want %v", test.skip, keep, got, test.want)
		}
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		scores []byte
		want   int
	}{
		{scores: []byte{0, 1, 2, 4, 5}, want: 5},
		{scores: []byte{5, 1, 5, 8, 9}, want: 9},
		{scores: []byte{9, 2, 5, 1, 0}, want: 18},
		{scores: []byte{5, 9, 4, 1, 4}, want: 2018},
	}
	for _, test := range tests {
		got := find(test.scores)
		if got != test.want {
			t.Errorf("find(%v) = %d, want %d", test.scores, got, test.want)
		}
	}
}
