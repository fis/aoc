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

package day01

import (
	"testing"
)

func TestCalibrationValue(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"1abc2", 12},
		{"pqr3stu8vwx", 38},
		{"a1b2c3d4e5f", 15},
		{"treb7uchet", 77},
	}
	for _, test := range tests {
		if got := calibrationValue(test.input); got != test.want {
			t.Errorf("calibrationValue(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}

func TestCalibrationValueEx(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, test := range tests {
		if got := calibrationValueEx(test.input); got != test.want {
			t.Errorf("calibrationValueEx(%s) = %d, want %d", test.input, got, test.want)
		}
	}
}
