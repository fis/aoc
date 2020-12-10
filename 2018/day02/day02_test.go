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
)

func TestChecksum(t *testing.T) {
	ex := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}
	want := 12
	got := checksum(ex)
	if got != want {
		t.Errorf("checksum = %d, want %d", got, want)
	}
}

func TestDiff1(t *testing.T) {
	tests := []struct {
		left, right string
		wantOK      bool
		wantS       string
	}{
		{left: "abcde", right: "axcye", wantOK: false, wantS: ""},
		{left: "fghij", right: "fguij", wantOK: true, wantS: "fgij"},
	}
	for _, test := range tests {
		if ok, s := diff1(test.left, test.right); ok != test.wantOK || s != test.wantS {
			t.Errorf("diff1(%q, %q) = (%v, %q), want (%v, %q)", test.left, test.right, ok, s, test.wantOK, test.wantS)
		}
	}
}

func TestFindBox(t *testing.T) {
	ex := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}
	want := "fgij"
	got := findBox(ex)
	if got != want {
		t.Errorf("findBox = %q, want %q", got, want)
	}
}
