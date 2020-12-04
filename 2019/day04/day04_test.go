// Copyright 2019 Google LLC
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
)

func TestValidate1(t *testing.T) {
	tests := []struct {
		digits string
		want   bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
	}
	for _, test := range tests {
		got := validate1([]byte(test.digits))
		if got != test.want {
			t.Errorf("validate1(%s) = %v, want %v", test.digits, got, test.want)
		}
	}
}

func TestValidate2(t *testing.T) {
	tests := []struct {
		digits string
		want   bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}
	for _, test := range tests {
		got := validate2([]byte(test.digits))
		if got != test.want {
			t.Errorf("validate2(%s) = %v, want %v", test.digits, got, test.want)
		}
	}
}
