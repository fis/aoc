// Copyright 2023 Google LLC
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

package day20

import "testing"

var (
	ex1 = []string{
		"broadcaster -> a, b, c",
		"%a -> b",
		"%b -> c",
		"%c -> inv",
		"&inv -> a",
	}
	ex2 = []string{
		"broadcaster -> a",
		"%a -> inv, con",
		"&inv -> b",
		"%b -> con",
		"&con -> output",
	}
)

func TestStepN(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"ex1", ex1, 32000000},
		{"ex2", ex2, 11687500},
	}
	for _, test := range tests {
		g, _, err := parseGraph(test.input)
		if err != nil {
			t.Errorf("parseGraph(%s): %v", test.name, err)
		} else if got := stepN(g, 1000); got != test.want {
			t.Errorf("stepN(%s, 1000) = %d, want %d", test.name, got, test.want)
		}
	}
}
