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

package day24

import "testing"

var (
	sampleImmuneSystem = []string{
		`Immune System:`,
		`17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2`,
		`989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3`,
	}
	sampleInfection = []string{
		`Infection:`,
		`801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1`,
		`4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`,
	}
)

func TestCombat(t *testing.T) {
	tests := []struct {
		boost int
		want  int
	}{
		{boost: 0, want: 5216},
		{boost: 1570, want: 51},
	}

	var s systemState
	if err := s.initialize(sampleImmuneSystem, sampleInfection); err != nil {
		t.Fatalf("initialize: %v", err)
	}

	for _, test := range tests {
		ts := s.copy()
		ts.boost(test.boost)
		if got, _ := ts.combat(); got != test.want {
			t.Errorf("[boost %d] combat() = %d, want %d", test.boost, got, test.want)
		}
	}
}
