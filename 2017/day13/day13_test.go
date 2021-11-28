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

package day13

import (
	"testing"
)

func TestSeverity(t *testing.T) {
	lines := []string{"0: 3", "1: 2", "4: 4", "6: 4"}
	tests := []struct {
		delay        int
		wantCaught   bool
		wantSeverity int
	}{
		{delay: 0, wantCaught: true, wantSeverity: 24},
		{delay: 10, wantCaught: false, wantSeverity: 0},
	}
	layers, err := parseInput(lines)
	if err != nil {
		t.Fatalf("parseInput: %v", err)
	}
	for _, test := range tests {
		gotCaught, gotSeverity := severity(layers, test.delay)
		if gotCaught != test.wantCaught || gotSeverity != test.wantSeverity {
			t.Errorf("severity(..., %d) = (%t, %d), want (%t, %d)", test.delay, gotCaught, gotSeverity, test.wantCaught, test.wantSeverity)
		}
	}
}

func TestCrack(t *testing.T) {
	lines := []string{"0: 3", "1: 2", "4: 4", "6: 4"}
	want := 10
	if layers, err := parseInput(lines); err != nil {
		t.Errorf("parseInput: %v", err)
	} else if got := crack(layers); got != want {
		t.Errorf("crack = %d, want %d", got, want)
	}
}
