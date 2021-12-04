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

package glue

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type TestCase struct {
	Day  int
	Want []string
}

func RunTests(t *testing.T, tests []TestCase, year int) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.Day), func(t *testing.T) {
			if got, err := SolveFile(year, test.Day, fmt.Sprintf("testdata/day%02d.txt", test.Day)); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.Want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func RunBenchmarks(b *testing.B, tests []TestCase, year int) {
	for _, test := range tests {
		b.Run(fmt.Sprintf("day=%02d", test.Day), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if got, err := SolveFile(year, test.Day, fmt.Sprintf("testdata/day%02d.txt", test.Day)); err != nil {
					b.Errorf("Solve: %v", err)
				} else if diff := cmp.Diff(test.Want, got); diff != "" {
					b.Errorf("Solve mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}
