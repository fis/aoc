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
		Want: []string{"239", "141"},
	},
	{
		Day:  2,
		Want: []string{"97289", "9A7DC"},
	},
	{
		Day:  3,
		Want: []string{"869", "1544"},
	},
	{
		Day:  4,
		Want: []string{"158835", "993"},
	},
	{
		Day:  5,
		Want: []string{"f77a0e6e", "999828ec"},
	},
	{
		Day:  6,
		Want: []string{"umcvzsmw", "rwqoacfz"},
	},
	{
		Day:  7,
		Want: []string{"115", "231"},
	},
	{
		Day: 8,
		Want: []string{
			"119",
			"#### #### #  # ####  ### ####  ##   ##  ###   ##  ",
			"   # #    #  # #    #    #    #  # #  # #  # #  # ",
			"  #  ###  #### ###  #    ###  #  # #    #  # #  # ",
			" #   #    #  # #     ##  #    #  # # ## ###  #  # ",
			"#    #    #  # #       # #    #  # #  # #    #  # ",
			"#### #    #  # #    ###  #     ##   ### #     ##  ",
		},
	},
	{
		Day:  9,
		Want: []string{"150914", "11052855125"},
	},
	{
		Day:  10,
		Want: []string{"56", "7847"},
	},
}

func TestAllDays(t *testing.T) {
	glue.RunTests(t, tests, 2016)
}

func BenchmarkAllDays(b *testing.B) {
	glue.RunBenchmarks(b, tests, 2016)
}
