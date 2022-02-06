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

package day05

import (
	"testing"
)

func TestSearch(t *testing.T) {
	want1, want2 := "18f47a30", "05ace8e3"
	got1, got2 := search("abc", 8)
	if got1 != want1 || got2 != want2 {
		t.Errorf("search(abc, 8) = (%s, %s), want (%s, %s)", got1, got2, want1, want2)
	}
}
