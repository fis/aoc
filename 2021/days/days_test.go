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

package days

import (
	"testing"

	"github.com/fis/aoc/glue"
)

var tests = []glue.TestCase{
	{
		Day:  1,
		Want: []string{"1529", "1567"},
	},
	{
		Day:  2,
		Want: []string{"1561344", "1848454425"},
	},
	{
		Day:  3,
		Want: []string{"2724524", "2775870"},
	},
	{
		Day:  4,
		Want: []string{"16716", "4880"},
	},
	{
		Day:  5,
		Want: []string{"5576", "18144"},
	},
	{
		Day:  6,
		Want: []string{"391671", "1754000560399"},
	},
	{
		Day:  7,
		Want: []string{"352997", "101571302"},
	},
	{
		Day:  8,
		Want: []string{"512", "1091165"},
	},
	{
		Day:  9,
		Want: []string{"478", "1327014"},
	},
	{
		Day:  10,
		Want: []string{"413733", "3354640192"},
	},
	{
		Day:  11,
		Want: []string{"1588", "517"},
	},
	{
		Day:  12,
		Want: []string{"5228", "131228"},
	},
	{
		Day: 13,
		Want: []string{
			"638",
			" ##    ##  ##  #  # ###   ##  ###  ### ",
			"#  #    # #  # # #  #  # #  # #  # #  #",
			"#       # #    ##   ###  #  # #  # ### ",
			"#       # #    # #  #  # #### ###  #  #",
			"#  # #  # #  # # #  #  # #  # #    #  #",
			" ##   ##   ##  #  # ###  #  # #    ### ",
		},
	},
	{
		Day:  14,
		Want: []string{"2509", "2827627697643"},
	},
	{
		Day:  15,
		Want: []string{"398", "2817"},
	},
	{
		Day:  16,
		Want: []string{"965", "116672213160"},
	},
	{
		Day:  17,
		Want: []string{"15931", "2555"},
	},
	{
		Day:  18,
		Want: []string{"4289", "4807"},
	},
	{
		Day:  19,
		Want: []string{"512", "16802"},
	},
	{
		Day:  20,
		Want: []string{"5359", "12333"},
	},
	{
		Day:  21,
		Want: []string{"908091", "190897246590017"},
	},
	{
		Day:  22,
		Want: []string{"546724", "1346544039176841"},
	},
	{
		Day:  23,
		Want: []string{"14350", "49742"},
	},
	{
		Day:  24,
		Want: []string{"49917929934999", "11911316711816"},
	},
}

func TestAllDays(t *testing.T) {
	glue.RunTests(t, tests, 2021)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, tests, 2021)
}
