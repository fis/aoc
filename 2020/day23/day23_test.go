// Copyright 2020 Google LLC
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

package day23

import (
	"testing"
)

func TestRingMove(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"389125467", "289154673"},
		{"289154673", "546789132"},
		{"546789132", "891346725"},
		{"891346725", "467913258"},
		{"467913258", "136792584"},
		{"136792584", "936725841"},
		{"936725841", "258367419"},
		{"258367419", "674158392"},
		{"674158392", "574183926"},
		{"574183926", "837419265"},
	}
	for _, test := range tests {
		if r1, err := newRing(test.in); err != nil {
			t.Errorf("newRing(%s): %v", test.in, err)
		} else if got := r1.move().String(); got != test.want {
			t.Errorf("%s.move() = %s, want %s", test.in, got, test.want)
		}
	}
}

func TestRingKey(t *testing.T) {
	r, _ := newRing("837419265")
	want := 92658374
	got := r.key()
	if got != want {
		t.Errorf("(837419265).key() = %d, want %d", got, want)
	}
}

func TestBigRing(t *testing.T) {
	r, _ := newBigRing("389125467")
	for i := 0; i < 10000000; i++ {
		r.move()
	}
	want := 149245887792
	got := r.key()
	if got != want {
		t.Errorf("big ring x10M key = %d, want %d", got, want)
	}
}
