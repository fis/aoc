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

package day25

import "testing"

var algos = []struct {
	name string
	f    func(b, a, m int) int
}{
	{name: "trialMultiplication", f: trialMultiplication},
	{name: "babyStep", f: babyStep},
}

func TestFindKey(t *testing.T) {
	pub1, pub2 := 5764801, 17807724
	want := 14897079
	for _, algo := range algos {
		got := findKey(pub1, pub2, algo.f)
		if got != want {
			t.Errorf("findKey(%d, %d, %s) = %d, want %d", pub1, pub2, algo.name, got, want)
		}
	}
}

func BenchmarkAlgos(b *testing.B) {
	inB, inA, inM := 7, 9232416, 20201227
	for _, algo := range algos {
		b.Run(algo.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algo.f(inB, inA, inM)
			}
		})
	}
}
