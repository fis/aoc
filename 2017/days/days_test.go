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

func TestAllDays(t *testing.T) {
	tests := []glue.TestCase{
		{
			Day:  1,
			Want: []string{"1044", "1054"},
		},
		{
			Day:  2,
			Want: []string{"44670", "285"},
		},
		{
			Day:  3,
			Want: []string{"430", "312453"},
		},
		{
			Day:  4,
			Want: []string{"455", "186"},
		},
	}
	glue.RunTests(t, tests, 2017)
}
