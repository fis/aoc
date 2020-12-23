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

package days

import (
	"fmt"
	"testing"

	"github.com/fis/aoc-go/glue"
	"github.com/google/go-cmp/cmp"
)

func TestAllDays(t *testing.T) {
	tests := []struct {
		day  int
		want []string
	}{
		{
			day:  1,
			want: []string{"970816", "96047280"},
		},
		{
			day:  2,
			want: []string{"506", "443"},
		},
		{
			day:  3,
			want: []string{"151", "7540141059"},
		},
		{
			day:  4,
			want: []string{"170", "103"},
		},
		{
			day:  5,
			want: []string{"813", "612"},
		},
		{
			day:  6,
			want: []string{"6249", "3103"},
		},
		{
			day:  7,
			want: []string{"289", "30055"},
		},
		{
			day:  8,
			want: []string{"1749", "515"},
		},
		{
			day:  9,
			want: []string{"32321523", "4794981"},
		},
		{
			day:  10,
			want: []string{"2376", "129586085429248"},
		},
		{
			day:  11,
			want: []string{"2164", "1974"},
		},
		{
			day:  12,
			want: []string{"362", "29895"},
		},
		{
			day:  13,
			want: []string{"3882", "867295486378319"},
		},
		{
			day:  14,
			want: []string{"6631883285184", "3161838538691"},
		},
		{
			day:  15,
			want: []string{"1259", "689"},
		},
		{
			day:  16,
			want: []string{"27802", "279139880759"},
		},
		{
			day:  17,
			want: []string{"230", "1600"},
		},
		{
			day:  18,
			want: []string{"1890866893020", "34646237037193"},
		},
		{
			day:  19,
			want: []string{"205", "329"},
		},
		{
			day:  20,
			want: []string{"28057939502729", "2489"},
		},
		{
			day:  21,
			want: []string{"2517", "rhvbn,mmcpg,kjf,fvk,lbmt,jgtb,hcbdb,zrb"},
		},
		{
			day:  22,
			want: []string{"31809", "32835"},
		},
		{
			day:  23,
			want: []string{"24987653", "442938711161"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("day=%02d", test.day), func(t *testing.T) {
			if got, err := glue.SolveFile(2020, test.day, fmt.Sprintf("testdata/day%02d.txt", test.day)); err != nil {
				t.Errorf("Solve: %v", err)
			} else if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("Solve mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
