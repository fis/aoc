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

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

func RunTests(t *testing.T, testRoot string, year int) {
	tests := findTests(testRoot, year)
	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.day), func(t *testing.T) {
			if got, err := SolveFile(year, test.day, test.inputFile); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func RunBenchmarks(b *testing.B, testRoot string, year int) {
	tests := findTests(testRoot, year)
	for _, test := range tests {
		b.Run(fmt.Sprintf("day=%02d", test.day), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if got, err := SolveFile(year, test.day, test.inputFile); err != nil {
					b.Errorf("Solve: %v", err)
				} else if diff := cmp.Diff(test.want, got); diff != "" {
					b.Errorf("Solve mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

type testCase struct {
	day       int
	inputFile string
	want      []string
}

func findTests(testRoot string, year int) []testCase {
	var tests []testCase
	for day := 1; day <= 25; day++ {
		basePath := fmt.Sprintf("%s/%04d/day%02d", testRoot, year, day)
		want, err := util.ReadLines(basePath + ".out")
		if err != nil {
			continue
		}
		tests = append(tests, testCase{
			day:       day,
			inputFile: basePath + ".txt",
			want:      want,
		})
	}
	return tests
}
