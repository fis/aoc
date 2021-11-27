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

package day04

import (
	"testing"

	"github.com/fis/aoc/util"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		passphrase string
		want       bool
	}{
		{passphrase: "aa bb cc dd ee", want: true},
		{passphrase: "aa bb cc dd aa", want: false},
		{passphrase: "aa bb cc dd aaa", want: true},
	}
	for _, test := range tests {
		if got := part1(util.Words(test.passphrase)); got != test.want {
			t.Errorf("part1(%s) = %t, want %t", test.passphrase, got, test.want)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		passphrase string
		want       bool
	}{
		{passphrase: "abcde fhgij", want: true},
		{passphrase: "abcde xyz ecdab", want: false},
		{passphrase: "a ab abc abd abf abj", want: true},
		{passphrase: "iiii oiii ooii oooi oooo", want: true},
		{passphrase: "oiii ioii iioi iiio", want: false},
	}
	for _, test := range tests {
		if got := part2(util.Words(test.passphrase)); got != test.want {
			t.Errorf("part2(%s) = %t, want %t", test.passphrase, got, test.want)
		}
	}
}
