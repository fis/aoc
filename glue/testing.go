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
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"slices"
	"strconv"
	"testing"

	"github.com/fis/aoc/util"
	"github.com/google/go-cmp/cmp"
)

func RunTests(t *testing.T, testRoot string, year int) {
	tests, err := FindTests(testRoot, year)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.Day), func(t *testing.T) {
			if got, err := SolveFile(year, test.Day, test.InputFile); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.Want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func RunBenchmarks(b *testing.B, testRoot string, year int) {
	tests, err := FindTests(testRoot, year)
	if err != nil {
		b.Fatal(err)
	}
	for _, test := range tests {
		b.Run(fmt.Sprintf("day=%02d", test.Day), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if got, err := SolveFile(year, test.Day, test.InputFile); err != nil {
					b.Errorf("Solve: %v", err)
				} else if diff := cmp.Diff(test.Want, got); diff != "" {
					b.Errorf("Solve mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

type TestCase struct {
	Year      int
	Day       int
	InputFile string
	Want      []string
}

var reYearDir = regexp.MustCompile(`^\d{4}$`)

func FindAllTests(testRoot string) (tests []TestCase, err error) {
	var years []int
	items, err := os.ReadDir(testRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to read testdata directory: %w", err)
	}
	for _, item := range items {
		if item.IsDir() && reYearDir.MatchString(item.Name()) {
			year, _ := strconv.Atoi(item.Name())
			years = append(years, year)
		}
	}
	slices.Sort(years)
	for _, year := range years {
		subTests, err := FindTests(testRoot, year)
		if err != nil {
			return nil, err
		}
		tests = append(tests, subTests...)
	}
	return tests, nil
}

func FindTests(testRoot string, year int) (tests []TestCase, err error) {
	for day := 1; day <= 25; day++ {
		basePath := fmt.Sprintf("%s/%04d/day%02d", testRoot, year, day)
		want, err := util.ReadLines(basePath + ".out")
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return nil, fmt.Errorf("failed to read test output: %w", err)
		}
		tests = append(tests, TestCase{
			Year:      year,
			Day:       day,
			InputFile: basePath + ".txt",
			Want:      want,
		})
	}
	return tests, nil
}
