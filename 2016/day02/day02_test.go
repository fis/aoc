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

package day02

import (
	"testing"
)

func TestDecode(t *testing.T) {
	sheet := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
	tests := []struct {
		name    string
		decoder func([]string) string
		want    string
	}{
		{name: "decode", decoder: decode, want: "1985"},
		{name: "decodeCross", decoder: decodeCross, want: "5DB3"},
	}
	for _, test := range tests {
		got := test.decoder(sheet)
		if got != test.want {
			t.Errorf("%s(%v) = %s, want %s", test.name, sheet, got, test.want)
		}
	}
}
