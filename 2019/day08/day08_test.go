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

package day08

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRender(t *testing.T) {
	layers := decode([]byte("0222112222120000"), 2, 2)
	img := render(layers)
	want := []byte{0, 1, 1, 0}
	if !cmp.Equal(img, want) {
		t.Errorf("render(%v) = %v, want %v", layers, img, want)
	}
}
