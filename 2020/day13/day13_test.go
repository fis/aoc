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

package day13

import (
	"testing"
)

func TestNextBus(t *testing.T) {
	wantID, wantWait := 59, 5
	gotID, gotWait := nextBus(939, []int{7, 13, -1, -1, 59, -1, 31, 19})
	if gotID != wantID || gotWait != wantWait {
		t.Errorf("nextBus = (%d, %d), want (%d, %d)", gotID, gotWait, wantID, wantWait)
	}
}

func TestBestTime(t *testing.T) {
	tests := []struct {
		buses []int
		want  int
	}{
		{buses: []int{7, 13, -1, -1, 59, -1, 31, 19}, want: 1068781},
		{buses: []int{17, -1, 13, 19}, want: 3417},
		{buses: []int{67, 7, 59, 61}, want: 754018},
		{buses: []int{67, -1, 7, 59, 61}, want: 779210},
		{buses: []int{67, 7, -1, 59, 61}, want: 1261476},
		{buses: []int{1789, 37, 47, 1889}, want: 1202161486},
	}
	for _, test := range tests {
		got := bestTime(test.buses)
		if got != test.want {
			t.Errorf("bestTime(%v) = %d, want %d", test.buses, got, test.want)
		}
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		ca, cb constraint
		want   constraint
	}{
		{ca: constraint{off: 12, mod: 13}, cb: constraint{off: 25, mod: 31}, want: constraint{off: 25, mod: 403}},
		{ca: constraint{off: 25, mod: 403}, cb: constraint{off: 12, mod: 19}, want: constraint{off: 4458, mod: 7657}},
	}
	for _, test := range tests {
		got := merge(test.ca, test.cb)
		if got.off != test.want.off || got.mod != test.want.mod {
			t.Errorf("merge(%v, %v) = %v, want %v", test.ca, test.cb, got, test.want)
		}
	}
}
