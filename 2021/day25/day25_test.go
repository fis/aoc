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
	lines, err := util.ReadLines("../../testdata/2021/day25.txt")
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
			for range b.N {
				data := alg.copy(datas[i])
				if got := alg.run(data); got != want {
					b.Errorf("%s(day25) = %d, want %d", alg.name, got, want)
				}
			}
		})
	}
}
