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

package day08

import "testing"

var example = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}

func TestChecksum(t *testing.T) {
	n, _ := parseLicense(example)
	want := 138
	got := n.checksum()
	if got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}

func TestValue(t *testing.T) {
	n, _ := parseLicense(example)
	want := 66
	got := n.value()
	if got != want {
		t.Errorf("value = %d, want %d", got, want)
	}
}
