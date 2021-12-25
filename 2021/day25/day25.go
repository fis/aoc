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

// Package day25 solves AoC 2021 day 25.
package day25

import (
	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 25, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	sf := parseInputBits(lines)
	p1 := simulateBits(sf)
	return glue.Ints(p1), nil
}

type seaCumber byte

const (
	noCumber    = 0
	eastCumber  = 1
	southCumber = 2
)

type seaFloor struct {
	w, h    uint32
	cumbers []seaCumber
}

func simulateCopying(sf seaFloor) (steps int) {
	halfStep := make([]seaCumber, sf.w*sf.h)
	for {
		steps++
		moves := false
		for y := uint32(0); y < sf.h; y++ {
			for x := uint32(0); x < sf.w; x++ {
				if c := sf.cumbers[y*sf.w+x]; c == eastCumber && sf.cumbers[y*sf.w+(x+1)%sf.w] == noCumber {
					moves = true
					halfStep[y*sf.w+x] = noCumber
					halfStep[y*sf.w+(x+1)%sf.w] = eastCumber
					x++
				} else {
					halfStep[y*sf.w+x] = c
				}
			}
		}
		for x := uint32(0); x < sf.w; x++ {
			for y := uint32(0); y < sf.h; y++ {
				if c := halfStep[y*sf.w+x]; c == southCumber && halfStep[((y+1)%sf.h)*sf.w+x] == noCumber {
					moves = true
					sf.cumbers[y*sf.w+x] = noCumber
					sf.cumbers[((y+1)%sf.h)*sf.w+x] = southCumber
					y++
				} else {
					sf.cumbers[y*sf.w+x] = c
				}
			}
		}
		if !moves {
			break
		}
	}
	return steps
}

func simulateInplace(sf seaFloor) (steps int) {
	for {
		steps++
		moves := false
		for y := uint32(0); y < sf.h; y++ {
			old1 := sf.cumbers[y*sf.w]
			for x := uint32(0); x < sf.w; x++ {
				if c := sf.cumbers[y*sf.w+x]; c == eastCumber {
					var nc seaCumber
					if x == sf.w-1 {
						nc = old1
					} else {
						nc = sf.cumbers[y*sf.w+(x+1)]
					}
					if nc == noCumber {
						moves = true
						sf.cumbers[y*sf.w+x] = noCumber
						sf.cumbers[y*sf.w+(x+1)%sf.w] = eastCumber
						x++
					}
				}
			}
		}
		for x := uint32(0); x < sf.w; x++ {
			old1 := sf.cumbers[x]
			for y := uint32(0); y < sf.h; y++ {
				if c := sf.cumbers[y*sf.w+x]; c == southCumber {
					var nc seaCumber
					if y == sf.h-1 {
						nc = old1
					} else {
						nc = sf.cumbers[(y+1)*sf.w+x]
					}
					if nc == noCumber {
						moves = true
						sf.cumbers[y*sf.w+x] = noCumber
						sf.cumbers[((y+1)%sf.h)*sf.w+x] = southCumber
						y++
					}
				}
			}
		}
		if !moves {
			break
		}
	}
	return steps
}

func parseInput(lines []string) (sf seaFloor) {
	w, h := len(lines[0]), len(lines)
	cumbers := make([]seaCumber, w*h)
	for y, line := range lines {
		for x := 0; x < w; x++ {
			switch line[x] {
			case '>':
				cumbers[y*w+x] = eastCumber
			case 'v':
				cumbers[y*w+x] = southCumber
			}
		}
	}
	return seaFloor{w: uint32(w), h: uint32(h), cumbers: cumbers}
}

type bitFloor struct {
	w, wu, h uint32
	bits     []uint32
}

func simulateBits(bf bitFloor) (steps int) {
	for {
		steps++
		moves := uint32(0)
		for y := uint32(0); y < bf.h; y++ {
			old1 := bf.bits[y*bf.wu] >> 30
			carry := uint32(0)
			for x := uint32(0); x < bf.wu; x++ {
				var b, nb uint32
				b = bf.bits[y*bf.wu+x]
				if x < bf.wu-1 {
					nb = bf.bits[y*bf.wu+x+1] >> 30
				} else if bf.w&0b1111 != 0 {
					b |= old1 << (2 * (15 - (bf.w & 0b1111)))
					nb = 0
				} else {
					nb = old1
				}
				easts := b & 0b01010101010101010101010101010101
				easts |= easts << 1
				empty := (b << 2) | nb
				empty = ^(empty ^ (empty >> 1))
				empty &= 0b01010101010101010101010101010101
				empty |= empty << 1
				moving := easts & empty
				bf.bits[y*bf.wu+x] = (b & ^moving) | ((b & moving) >> 2) | carry
				carry = (0b01 & moving) << 30
				if x == bf.wu-1 && bf.w&0b1111 != 0 {
					slopBits := 2 * (16 - (bf.w & 0b1111))
					slop := bf.bits[y*bf.wu+x] >> (slopBits - 2)
					slopMask := ^((uint32(1) << slopBits) - 1)
					bf.bits[y*bf.wu] |= slop << 30
					bf.bits[y*bf.wu+x] &= slopMask
					b &= slopMask
				}
				moves |= (b & moving)
			}
			if bf.w&0b1111 == 0 {
				bf.bits[y*bf.wu] |= carry
			}
		}
		for x := uint32(0); x < bf.wu; x++ {
			old1 := bf.bits[x]
			carry := uint32(0)
			for y := uint32(0); y < bf.h; y++ {
				var b, nb uint32
				b = bf.bits[y*bf.wu+x]
				if y < bf.h-1 {
					nb = bf.bits[(y+1)*bf.wu+x]
				} else {
					nb = old1
				}
				souths := b & 0b10101010101010101010101010101010
				souths |= souths >> 1
				empty := ^(nb ^ (nb >> 1))
				empty &= 0b01010101010101010101010101010101
				empty |= empty << 1
				moving := souths & empty
				bf.bits[y*bf.wu+x] = (b & ^moving) | carry
				carry = b & moving
				moves |= (b & moving)
			}
			bf.bits[x] |= carry
		}
		if moves == 0 {
			break
		}
	}
	return steps
}

func parseInputBits(lines []string) (bf bitFloor) {
	w, h := len(lines[0]), len(lines)
	wu := (w + 15) >> 4
	cumbers := make([]uint32, wu*h)
	for y, line := range lines {
		for x := 0; x < w; x++ {
			switch line[x] {
			case '>':
				cumbers[y*wu+(x>>4)] |= 0b01 << (30 - ((x & 0b1111) << 1))
			case 'v':
				cumbers[y*wu+(x>>4)] |= 0b10 << (30 - ((x & 0b1111) << 1))
			}
		}
	}
	return bitFloor{w: uint32(w), wu: uint32(wu), h: uint32(h), bits: cumbers}
}
