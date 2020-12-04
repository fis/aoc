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

package day09

import (
	"testing"

	"github.com/fis/aoc-go/intcode"
	"github.com/google/go-cmp/cmp"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		prog []int64
		want []int64
	}{
		{
			prog: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			want: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			prog: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			want: []int64{1219070632396864},
		},
		{
			prog: []int64{104, 1125899906842624, 99},
			want: []int64{1125899906842624},
		},
	}
	for _, test := range tests {
		got, _ := intcode.Run(test.prog, nil)
		if !cmp.Equal(got, test.want) {
			t.Errorf("%v -> %v, want %v", test.prog, got, test.want)
		}
	}
}
