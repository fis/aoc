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

package ix

import "testing"

func TestAbs(t *testing.T) {
	tests := [][2]int{{-100, 100}, {-20, 20}, {-1, 1}, {0, 0}, {1, 1}, {20, 20}, {100, 100}}
	for _, test := range tests {
		if got := Abs(test[0]); got != test[1] {
			t.Errorf("Abs(%d) = %d, want %d", test[0], got, test[1])
		}
	}
}

func TestMax(t *testing.T) {
	tests := [][3]int{
		{-100, 100, 100}, {-20, 20, 20}, {-1, 1, 1},
		{-5, -5, -5}, {0, 0, 0}, {5, 5, 5},
		{1, -1, 1}, {20, -20, 20}, {100, -100, 100},
	}
	for _, test := range tests {
		if got := Max(test[0], test[1]); got != test[2] {
			t.Errorf("Max(%d, %d) = %d, want %d", test[0], test[1], got, test[2])
		}
	}
}

func TestMin(t *testing.T) {
	tests := [][3]int{
		{-100, 100, -100}, {-20, 20, -20}, {-1, 1, -1},
		{-5, -5, -5}, {0, 0, 0}, {5, 5, 5},
		{1, -1, -1}, {20, -20, -20}, {100, -100, -100},
	}
	for _, test := range tests {
		if got := Min(test[0], test[1]); got != test[2] {
			t.Errorf("Min(%d, %d) = %d, want %d", test[0], test[1], got, test[2])
		}
	}
}

func TestSign(t *testing.T) {
	tests := [][2]int{{-100, -1}, {-20, -1}, {-1, -1}, {0, 0}, {1, 1}, {20, 1}, {100, 1}}
	for _, test := range tests {
		if got := Sign(test[0]); got != test[1] {
			t.Errorf("Sign(%d) = %d, want %d", test[0], got, test[1])
		}
	}
}

func TestGCD(t *testing.T) {
	tests := [][3]int{
		{10, 15, 5},
		{20, 30, 10},
		{15, 10, 5},
		{30, 20, 10},
		{17, 19, 1},
		{20, 10, 10},
		{10, 10, 10},
		{1234567890, 987654321, 9},
		{0, 0, 0},
	}
	for _, test := range tests {
		if got := GCD(test[0], test[1]); got != test[2] {
			t.Errorf("GCD(%d, %d) = %d, want %d", test[0], test[1], got, test[2])
		}
	}
}

func TestLCM(t *testing.T) {
	tests := [][3]int{
		{10, 15, 30},
		{20, 30, 60},
		{15, 10, 30},
		{30, 20, 60},
		{17, 19, 323},
		{20, 10, 20},
		{10, 10, 10},
		{1234567890, 987654321, 135480701236261410},
	}
	for _, test := range tests {
		if got := LCM(test[0], test[1]); got != test[2] {
			t.Errorf("LCM(%d, %d) = %d, want %d", test[0], test[1], got, test[2])
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := [][2]int{
		{0, 0},
		{1, 1}, {2, 1}, {3, 1},
		{4, 2}, {5, 2}, {8, 2},
		{9, 3}, {10, 3}, {15, 3},
		{16, 4}, {17, 4}, {24, 4},
		{25, 5},
		{9999, 99}, {10000, 100},
	}
	for _, test := range tests {
		if got := Sqrt(test[0]); got != test[1] {
			t.Errorf("Sqrt(%d) = %d, want %d", test[0], got, test[1])
		}
	}
}
