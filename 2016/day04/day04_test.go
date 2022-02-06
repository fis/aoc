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

package day04

import (
	"testing"
)

func TestRoomValidate(t *testing.T) {
	tests := []struct {
		line       string
		wantSector int
		wantValid  bool
	}{
		{line: "aaaaa-bbb-z-y-x-123[abxyz]", wantSector: 123, wantValid: true},
		{line: "a-b-c-d-e-f-g-h-987[abcde]", wantSector: 987, wantValid: true},
		{line: "not-a-real-room-404[oarel]", wantSector: 404, wantValid: true},
		{line: "totally-real-room-200[decoy]", wantSector: 200, wantValid: false},
	}
	for _, test := range tests {
		if r, err := parseRoom(test.line); err != nil {
			t.Errorf("parseRoom(%s): %v", test.line, err)
		} else if r.sector != test.wantSector {
			t.Errorf("parseRoom(%s).sector = %d, want %d", test.line, r.sector, test.wantSector)
		} else if valid := r.validate(); valid != test.wantValid {
			t.Errorf("parseRoom(%s).validate() = %t, want %t", test.line, valid, test.wantValid)
		}
	}
}

func TestRoomDecode(t *testing.T) {
	line := "qzmt-zixmtkozy-ivhz-343[xxxxx]"
	want := "very encrypted name"
	if r, err := parseRoom(line); err != nil {
		t.Errorf("parseRoom(%s): %v", line, err)
	} else if got := r.decode(); got != want {
		t.Errorf("parseRoom(%s).decode() = %q, want %q", line, got, want)
	}
}
