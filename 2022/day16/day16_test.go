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

package day16

import (
	"testing"

	"github.com/fis/aoc/util/fn"
)

func TestReleasePressure(t *testing.T) {
	scan, err := fn.MapE(ex, parseValveScan)
	if err != nil {
		t.Fatal(err)
	}
	sum := preprocess(scan)
	want := 1651
	if got := releasePressure(sum, 30); got != want {
		t.Errorf("releasePressure(ex, 30) = %d, want %d", got, want)
	}
}

func TestReleasePressure2(t *testing.T) {
	scan, err := fn.MapE(ex, parseValveScan)
	if err != nil {
		t.Fatal(err)
	}
	sum := preprocess(scan)
	want := 1707
	if got := releasePressure2(sum, 26); got != want {
		t.Errorf("releasePressure2(ex, 26) = %d, want %d", got, want)
	}
}
