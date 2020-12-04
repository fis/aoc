// Copyright 2019 Google LLC
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

package day22

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExamples(t *testing.T) {
	const deckSize int64 = 10
	tests := []struct {
		comment string
		shuffle []string
		want    []int64
	}{
		{
			comment: "deal",
			shuffle: []string{"deal into new stack"},
			want:    []int64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			comment: "cut (positive)",
			shuffle: []string{"cut 3"},
			want:    []int64{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			comment: "cut (negative)",
			shuffle: []string{"cut -4"},
			want:    []int64{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
		{
			comment: "interleave",
			shuffle: []string{"deal with increment 3"},
			want:    []int64{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		{
			comment: "example 1",
			shuffle: []string{
				"deal with increment 7",
				"deal into new stack",
				"deal into new stack",
			},
			want: []int64{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			comment: "example 2",
			shuffle: []string{
				"cut 6",
				"deal with increment 7",
				"deal into new stack",
			},
			want: []int64{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			comment: "example 3",
			shuffle: []string{
				"deal with increment 7",
				"deal with increment 9",
				"cut -2",
			},
			want: []int64{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			comment: "example 4",
			shuffle: []string{
				"deal into new stack",
				"cut -2",
				"deal with increment 7",
				"cut 8",
				"cut -4",
				"deal with increment 7",
				"cut 3",
				"deal with increment 9",
				"deal with increment 3",
				"cut -1",
			},
			want: []int64{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}
	for _, test := range tests {
		ops, err := parseShuffle(test.shuffle)
		if err != nil {
			t.Errorf("%s: failed parse: %v", test.comment, err)
			continue
		}
		deck := make([]int64, deckSize)
		for in := int64(0); in < deckSize; in++ {
			out := shuffleForward(in, deckSize, ops)
			deck[out] = in
		}
		if !cmp.Equal(deck, test.want) {
			t.Errorf("%s: got %v, want %v", test.comment, deck, test.want)
		}
	}
}
