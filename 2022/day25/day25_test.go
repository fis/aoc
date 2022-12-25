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

package day25

import (
	"testing"
)

var snafus = []struct {
	s string
	i int
}{
	{"1", 1},
	{"2", 2},
	{"1=", 3},
	{"1-", 4},
	{"10", 5},
	{"11", 6},
	{"12", 7},
	{"2=", 8},
	{"2-", 9},
	{"20", 10},
	{"1=0", 15},
	{"1-0", 20},
	{"1=11-2", 2022},
	{"1-0---0", 12345},
	{"1121-1110-1=0", 314159265},
	{"1=-0-2", 1747},
	{"12111", 906},
	{"2=0=", 198},
	{"21", 11},
	{"2=01", 201},
	{"111", 31},
	{"20012", 1257},
	{"112", 32},
	{"1=-1=", 353},
	{"1-12", 107},
	{"12", 7},
	{"1=", 3},
	{"122", 37},
}

func TestParseSNAFU(t *testing.T) {
	for _, test := range snafus {
		if got, err := parseSNAFU(test.s); err != nil {
			t.Errorf("parseSNAFU(%s): %v", test.s, err)
		} else if got != test.i {
			t.Errorf("parseSNAFU(%s) = %d, want %d", test.s, got, test.i)
		}
	}
}

func TestFormatSNAFU(t *testing.T) {
	for _, test := range snafus {
		if got := formatSNAFU(test.i); got != test.s {
			t.Errorf("formatSNAFU(%d) = %s, want %s", test.i, got, test.s)
		}
	}
}
