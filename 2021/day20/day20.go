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

// Package day20 solves AoC 2021 day 20.
package day20

import (
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2021, 20, glue.ChunkSolver(solve))
}

func solve(chunks []string) ([]string, error) {
	p1, p2 := enhanceBits(chunks[0], util.Lines(chunks[1]))
	return glue.Ints(p1, p2), nil
}

func enhanceBytes(algoLine string, imgLines []string) (lit2, lit50 int) {
	const steps = 50

	var algo [512]byte
	for i := 0; i < 512; i++ {
		if algoLine[i] == '#' {
			algo[i] = 1
		}
	}

	x0, x1 := steps+1, steps+len(imgLines[0])
	y0, y1 := steps+1, steps+len(imgLines)
	w, h := x1+steps+2, y1+steps+2
	img, next := make([]byte, w*h), make([]byte, w*h)
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			if imgLines[y-(steps+1)][x-(steps+1)] == '#' {
				img[y*w+x] = 1
			}
		}
	}

	invert1, invert9 := [2]byte{0, 0}, [2]uint32{0, 0}
	if algo[0] == 1 {
		if algo[511] != 0 {
			panic("an infinitude of lit bits will result from this")
		}
		invert1[0], invert9[1] = 1, 0b111_111_111
	}

	for step := 0; step < steps; step++ {
		x0, x1, y0, y1 = x0-1, x1+1, y0-1, y1+1
		for y := y0; y <= y1; y++ {
			for x := x0; x <= x1; x++ {
				neigh := uint32(img[(y-1)*w+(x-1)])<<8 | uint32(img[(y-1)*w+x])<<7 | uint32(img[(y-1)*w+(x+1)])<<6
				neigh |= uint32(img[(y+0)*w+(x-1)])<<5 | uint32(img[(y+0)*w+x])<<4 | uint32(img[(y+0)*w+(x+1)])<<3
				neigh |= uint32(img[(y+1)*w+(x-1)])<<2 | uint32(img[(y+1)*w+x])<<1 | uint32(img[(y+1)*w+(x+1)])<<0
				neigh ^= invert9[step&1]
				next[y*w+x] = algo[neigh] ^ invert1[step&1]
			}
		}
		img, next = next, img
		if step == 1 || step == steps-1 {
			lit := 0
			for y := y0; y <= y1; y++ {
				for x := x0; x <= x1; x++ {
					lit += int(img[y*w+x])
				}
			}
			if step == 1 {
				lit2 = lit
			} else {
				lit50 = lit
			}
		}
	}

	return lit2, lit50
}

func enhanceBits(algoLine string, imgLines []string) (lit2, lit50 int) {
	const (
		steps uint32 = 50
		bitw  uint32 = 32
		shift uint32 = 5
		mask  uint32 = 0b11111
	)

	var algo [512]byte
	for i := 0; i < 512; i++ {
		if algoLine[i] == '#' {
			algo[i] = 1
		}
	}

	x0, x1 := steps+1, steps+uint32(len(imgLines[0]))
	y0, y1 := steps+1, steps+uint32(len(imgLines))
	w, h := (x1+steps+2+bitw-1)&^mask, (y1+steps+2+bitw-1)&^mask
	wu := w >> shift
	img, next := make([]uint32, wu*h), make([]uint32, wu*h)
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			if imgLines[y-(steps+1)][x-(steps+1)] == '#' {
				iu, ib := y*wu+(x>>shift), (^x)&mask
				img[iu] |= 1 << ib
			}
		}
	}

	invert1, invert9 := [2]byte{0, 0}, [2]uint32{0, 0}
	if algo[0] == 1 {
		if algo[511] != 0 {
			panic("an infinitude of lit bits will result from this")
		}
		invert1[0], invert9[1] = 1, 0b111_111_111
	}

	for step := uint32(0); step < steps; step++ {
		x0, x1, y0, y1 = x0-1, x1+1, y0-1, y1+1
		for y := y0; y <= y1; y++ {
			for xu := x0 >> shift; xu < (x1+bitw)>>shift; xu++ {
				next[y*wu+xu] = 0
			}
			for x := x0; x <= x1; x++ {
				iu, ib := y*wu+(x>>shift), (^x)&mask
				var neigh uint32
				switch ib {
				case 0:
					neigh = (img[iu-wu]&0b11)<<7 | ((img[iu-wu+1] >> (bitw - 1 - 6)) & 0b1_000_000)
					neigh |= (img[iu]&0b11)<<4 | ((img[iu+1] >> (bitw - 1 - 3)) & 0b1_000)
					neigh |= (img[iu+wu]&0b11)<<1 | (img[iu+wu+1] >> (bitw - 1))
				case bitw - 1:
					neigh = (img[iu-wu-1]&1)<<8 | ((img[iu-wu] >> (bitw - 2 - 6)) & 0b11_000_000)
					neigh |= (img[iu-1]&1)<<5 | ((img[iu] >> (bitw - 2 - 3)) & 0b11_000)
					neigh |= (img[iu+wu-1]&1)<<2 | (img[iu+wu] >> (bitw - 2))
				default:
					neigh = ((img[iu-wu] >> (ib - 1)) & 0b111) << 6
					neigh |= ((img[iu] >> (ib - 1)) & 0b111) << 3
					neigh |= (img[iu+wu] >> (ib - 1)) & 0b111
				}
				neigh ^= invert9[step&1]
				next[iu] |= uint32((algo[neigh] ^ invert1[step&1])) << ib
			}
		}
		img, next = next, img
		if step == 1 || step == steps-1 {
			lit := 0
			for y := y0; y <= y1; y++ {
				for xu := x0 >> shift; xu < (x1+bitw)>>shift; xu++ {
					lit += bits.OnesCount32(img[y*wu+xu])
				}
			}
			if step == 1 {
				lit2 = lit
			} else {
				lit50 = lit
			}
		}
	}

	return lit2, lit50
}
