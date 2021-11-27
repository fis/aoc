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
	"testing"

	"github.com/fis/aoc/glue"
)

func TestAllDays(t *testing.T) {
	tests := []glue.TestCase{
		{
			Day:  1,
			Want: []string{"970816", "96047280"},
		},
		{
			Day:  2,
			Want: []string{"506", "443"},
		},
		{
			Day:  3,
			Want: []string{"151", "7540141059"},
		},
		{
			Day:  4,
			Want: []string{"170", "103"},
		},
		{
			Day:  5,
			Want: []string{"813", "612"},
		},
		{
			Day:  6,
			Want: []string{"6249", "3103"},
		},
		{
			Day:  7,
			Want: []string{"289", "30055"},
		},
		{
			Day:  8,
			Want: []string{"1749", "515"},
		},
		{
			Day:  9,
			Want: []string{"32321523", "4794981"},
		},
		{
			Day:  10,
			Want: []string{"2376", "129586085429248"},
		},
		{
			Day:  11,
			Want: []string{"2164", "1974"},
		},
		{
			Day:  12,
			Want: []string{"362", "29895"},
		},
		{
			Day:  13,
			Want: []string{"3882", "867295486378319"},
		},
		{
			Day:  14,
			Want: []string{"6631883285184", "3161838538691"},
		},
		{
			Day:  15,
			Want: []string{"1259", "689"},
		},
		{
			Day:  16,
			Want: []string{"27802", "279139880759"},
		},
		{
			Day:  17,
			Want: []string{"230", "1600"},
		},
		{
			Day:  18,
			Want: []string{"1890866893020", "34646237037193"},
		},
		{
			Day:  19,
			Want: []string{"205", "329"},
		},
		{
			Day:  20,
			Want: []string{"28057939502729", "2489"},
		},
		{
			Day:  21,
			Want: []string{"2517", "rhvbn,mmcpg,kjf,fvk,lbmt,jgtb,hcbdb,zrb"},
		},
		{
			Day:  22,
			Want: []string{"31809", "32835"},
		},
		{
			Day:  23,
			Want: []string{"24987653", "442938711161"},
		},
		{
			Day:  24,
			Want: []string{"320", "3777"},
		},
		{
			Day:  25,
			Want: []string{"1478097"},
		},
	}
	glue.RunTests(t, tests, 2020)
}
