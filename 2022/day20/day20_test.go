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

package day20

import (
	"fmt"
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []int{1, 2, -3, 3, -2, 0, 4}

func TestDecrypt(t *testing.T) {
	algos := []struct {
		name string
		f    func(file []int, key, rounds int) int
	}{
		{"plain", decryptPlain},
		{"skip 50", func(file []int, key, rounds int) int { return decrypt(file, key, rounds, 50) }},
	}
	file, err := util.ReadInts("../../testdata/2022/day20.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 8119137886612
	for _, algo := range algos {
		if got := algo.f(file, 811589153, 10); got != want {
			t.Errorf("decrypt[%s](day20, 811589153, 10) = %d, want %d", algo.name, got, want)
		}
	}
}

func TestDecryptPlain(t *testing.T) {
	tests := []struct {
		key    int
		rounds int
		want   int
	}{
		{1, 1, 3},
		{811589153, 10, 1623178306},
	}
	for _, test := range tests {
		if got := decryptPlain(ex, test.key, test.rounds); got != test.want {
			t.Errorf("decryptPlain(%v, %d, %d) = %d, want %d", ex, test.key, test.rounds, got, test.want)
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	file, err := util.ReadInts("../../testdata/2022/day20.txt")
	if err != nil {
		b.Fatal(err)
	}
	skipSizes := []int{5, 10, 20, 25, 40, 50, 100, 125, 200, 250, 500, 625, 1000, 1250, 2500}
	want := 8119137886612
	for _, skipSize := range skipSizes {
		b.Run(fmt.Sprintf("size=%d", skipSize), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				if got := decrypt(file, 811589153, 10, skipSize); got != want {
					b.Errorf("decrypt2(day20, 811589153, 10, %d) = %d, want %d", skipSize, got, want)
				}
			}
		})
	}
}
