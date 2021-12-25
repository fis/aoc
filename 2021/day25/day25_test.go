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

package day25

import (
	"testing"

	"github.com/fis/aoc/util"
)

var ex = []string{
	`v...>>.vv>`,
	`.vv>>.vv..`,
	`>>.>v>...v`,
	`>>v>>.>.v.`,
	`v>v.vv.v..`,
	`>.>>..v...`,
	`.vv..>.>v.`,
	`v.v..>>v.v`,
	`....v..v.>`,
}

var algos = []struct {
	name string
	init func(lines []string) interface{}
	copy func(data interface{}) interface{}
	run  func(data interface{}) int
}{
	{
		name: "byteCopying",
		init: func(lines []string) interface{} { return parseInput(lines) },
		copy: copySeaFloor,
		run:  func(data interface{}) int { return simulateCopying(data.(seaFloor)) },
	},
	{
		name: "byteInplace",
		init: func(lines []string) interface{} { return parseInput(lines) },
		copy: copySeaFloor,
		run:  func(data interface{}) int { return simulateInplace(data.(seaFloor)) },
	},
	{
		name: "bits",
		init: func(lines []string) interface{} { return parseInputBits(lines) },
		copy: copyBitFloor,
		run:  func(data interface{}) int { return simulateBits(data.(bitFloor)) },
	},
}

func copySeaFloor(data interface{}) interface{} {
	sf := data.(seaFloor)
	return seaFloor{w: sf.w, h: sf.h, cumbers: append([]seaCumber(nil), sf.cumbers...)}
}

func copyBitFloor(data interface{}) interface{} {
	bf := data.(bitFloor)
	return bitFloor{w: bf.w, wu: bf.wu, h: bf.h, bits: append([]uint32(nil), bf.bits...)}
}

func TestSimulate(t *testing.T) {
	want := 58
	for _, alg := range algos {
		data := alg.init(ex)
		if got := alg.run(data); got != want {
			t.Errorf("%s(ex) = %d, want %d", alg.name, got, want)
		}
	}
}

func BenchmarkSimulate(b *testing.B) {
	lines, err := util.ReadLines("../days/testdata/day25.txt")
	if err != nil {
		b.Fatal(err)
	}
	datas := make([]interface{}, len(algos))
	for i, alg := range algos {
		datas[i] = alg.init(lines)
	}
	want := 458
	for i, alg := range algos {
		b.Run(alg.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				data := alg.copy(datas[i])
				if got := alg.run(data); got != want {
					b.Errorf("%s(day25) = %d, want %d", alg.name, got, want)
				}
			}
		})
	}
}
