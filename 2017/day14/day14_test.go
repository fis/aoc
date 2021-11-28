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

package day14

import (
	"testing"
)

func TestHash(t *testing.T) {
	key := "flqrgnkx"
	wantUsed, wantRegions := 8108, 1242

	hashes, gotUsed := hash(key)
	lvl := buildLevel(hashes)
	gotRegions := countRegions(lvl)

	if gotUsed != wantUsed {
		t.Errorf("part1(%q) = %d, want %d", key, gotUsed, wantUsed)
	}
	if gotRegions != wantRegions {
		t.Errorf("part2(%q) = %d, want %d", key, gotRegions, wantRegions)
	}
}
