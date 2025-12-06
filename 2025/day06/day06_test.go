// Copyright 2025 Google LLC
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

package day06

import "testing"

var ex = []string{
	"123 328  51 64 ",
	" 45 64  387 23 ",
	"  6 98  215 314",
	"*   +   *   +  ",
}

func TestWrong(t *testing.T) {
	want := 4277556
	if got := parseWorksheetWrong(ex).grandTotal(); got != want {
		t.Errorf("wrong(ex).grandtotal() = %d, want %d", got, want)
	}
}

func TestRight(t *testing.T) {
	want := 3263827
	if got := parseWorksheetRight(ex).grandTotal(); got != want {
		t.Errorf("right(ex).grandtotal() = %d, want %d", got, want)
	}
}
