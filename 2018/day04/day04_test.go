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

package day04

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var ex = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-05 00:55] wakes up",
}

func TestParseLog(t *testing.T) {
	want := map[int][]sleepMask{
		10: []sleepMask{
			mask(".....####################.....#########################....."),
			mask("........................#####..............................."),
		},
		99: []sleepMask{
			mask("........................................##########.........."),
			mask("....................................##########.............."),
			mask(".............................................##########....."),
		},
	}
	got := parseLog(ex)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parseLog mismatch (-want +got):\n%s", diff)
	}
}

func TestStrategies(t *testing.T) {
	tests := []struct {
		name         string
		strategy     func(map[int][]sleepMask) (int, int)
		wantG, wantM int
	}{
		{name: "strategy1", strategy: strategy1, wantG: 10, wantM: 24},
		{name: "strategy2", strategy: strategy2, wantG: 99, wantM: 45},
	}
	log := parseLog(ex)
	for _, test := range tests {
		gotG, gotM := test.strategy(log)
		if gotG != test.wantG || gotM != test.wantM {
			t.Errorf("%s = (%d, %d), want (%d, %d)", test.name, gotG, gotM, test.wantG, test.wantM)
		}
	}
}

func mask(s string) (m sleepMask) {
	for i, c := range s {
		m[i] = c == '#'
	}
	return m
}
