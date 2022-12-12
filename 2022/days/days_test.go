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

package days

import (
	"testing"

	"github.com/fis/aoc/glue"
)

var tests = []glue.TestCase{
	{
		Day:  1,
		Want: []string{"69836", "207968"},
	},
	{
		Day:  2,
		Want: []string{"12586", "13193"},
	},
	{
		Day:  3,
		Want: []string{"8185", "2817"},
	},
	{
		Day:  4,
		Want: []string{"550", "931"},
	},
	{
		Day:  5,
		Want: []string{"QGTHFZBHV", "MGDMPSZTM"},
	},
	{
		Day:  6,
		Want: []string{"1655", "2665"},
	},
	{
		Day:  7,
		Want: []string{"2061777", "4473403"},
	},
	{
		Day:  8,
		Want: []string{"1672", "327180"},
	},
	{
		Day:  9,
		Want: []string{"6391", "2593"},
	},
	{
		Day: 10,
		Want: []string{
			"11780",
			"###  #### #  # #    ###   ##  #  #  ##  ",
			"#  #    # #  # #    #  # #  # #  # #  # ",
			"#  #   #  #  # #    ###  #  # #  # #  # ",
			"###   #   #  # #    #  # #### #  # #### ",
			"#    #    #  # #    #  # #  # #  # #  # ",
			"#    ####  ##  #### ###  #  #  ##  #  # ",
		},
	},
	{
		Day:  11,
		Want: []string{"66124", "19309892877"},
	},
	{
		Day:  12,
		Want: []string{"437", "430"},
	},
}

func TestAllDays(t *testing.T) {
	glue.RunTests(t, tests, 2022)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, tests, 2022)
}
