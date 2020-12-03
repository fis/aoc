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

package day02

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var inputs = []struct {
	pol, pass string
}{
	{pol: "1-3 a", pass: "abcde"},
	{pol: "1-3 b", pass: "cdefg"},
	{pol: "2-9 c", pass: "ccccccccc"},
}

func TestParse(t *testing.T) {
	want := []policy{
		{Min: 1, Max: 3, C: 'a'},
		{Min: 1, Max: 3, C: 'b'},
		{Min: 2, Max: 9, C: 'c'},
	}
	for i, input := range inputs {
		if pol, err := parsePolicy(input.pol); err != nil {
			t.Errorf("invalid policy: %v", err)
		} else if !cmp.Equal(pol, want[i]) {
			t.Errorf("parsePolicy(%q) = %v, want %v", input.pol, pol, want[i])
		}
	}
}

func TestSled(t *testing.T) {
	want := map[string]bool{
		"abcde":     true,
		"cdefg":     false,
		"ccccccccc": true,
	}
	for _, input := range inputs {
		if pol, err := parsePolicy(input.pol); err != nil {
			t.Errorf("invalid policy: %v", err)
		} else if got := pol.validateSled(input.pass); got != want[input.pass] {
			t.Errorf("%v.validateSled(%q) = %v, want %v", pol, input.pass, got, want[input.pass])
		}
	}
}

func TestToboggan(t *testing.T) {
	want := map[string]bool{
		"abcde":     true,
		"cdefg":     false,
		"ccccccccc": false,
	}
	for _, input := range inputs {
		if pol, err := parsePolicy(input.pol); err != nil {
			t.Errorf("invalid policy: %v", err)
		} else if got := pol.validateToboggan(input.pass); got != want[input.pass] {
			t.Errorf("%v.validateToboggan(%q) = %v, want %v", pol, input.pass, got, want[input.pass])
		}
	}
}
