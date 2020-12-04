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

package day03

import (
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		w1, w2        string
		closest, best int
	}{
		{
			w1:      "R8,U5,L5,D3",
			w2:      "U7,R6,D4,L4",
			closest: 6,
			best:    30,
		},
		{
			w1:      "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			w2:      "U62,R66,U55,R34,D71,R55,D58,R83",
			closest: 159,
			best:    610,
		},
		{
			w1:      "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			w2:      "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			closest: 135,
			best:    410,
		},
	}
	for _, test := range tests {
		c, b := compute(test.w1, test.w2)
		if c != test.closest {
			t.Errorf("%q, %q -> closest %d, want %d", test.w1, test.w2, c, test.closest)
		}
		if b != test.best {
			t.Errorf("%q, %q -> best %d, want %d", test.w1, test.w2, b, test.best)
		}
	}
}
