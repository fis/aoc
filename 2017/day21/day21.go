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

// Package day21 solves AoC 2017 day 21.
package day21

import (
	"fmt"
	"math/bits"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2017, 21, glue.RegexpSolver{
		Solver: solve,
		Regexp: `^(\S+) => (\S+)$`,
	})
}

func solve(lines [][]string) ([]string, error) {
	book, err := parseBook(lines)
	if err != nil {
		return nil, err
	}

	bmp := newBitmap(3)
	bmp.put3(0, 0, rootTile)
	bmp = iterate(bmp, book, 5)
	p1 := bmp.popCount()
	bmp = iterate(bmp, book, 18-5)
	p2 := bmp.popCount()

	return glue.Ints(p1, p2), nil
}

const rootTile tile3 = 0b010_001_111

func iterate(src *bitmap, book *ruleBook, n int) *bitmap {
	diag := util.IsDiag()

	for i := 0; i < n; i++ {
		if diag {
			util.Diagf("i = %d:\n", i)
			src.print()
			util.Diagln()
		}
		if (src.size & 1) == 0 {
			src = book.expand2(src)
		} else {
			src = book.expand3(src)
		}
	}
	if diag {
		util.Diagf("i = %d:\n", n)
		src.print()
		util.Diagln()
	}

	return src
}

type ruleBook struct {
	rules2 [16]tile3
	rules3 [512]tile4
}

func parseBook(lines [][]string) (*ruleBook, error) {
	book := &ruleBook{}
	for _, rule := range lines {
		if len(rule[0]) == 5 && len(rule[1]) == 11 {
			base := parseTile2(rule[0])
			to := parseTile3(rule[1])
			for _, from := range base.flips() {
				book.rules2[from] = to
			}
		} else if len(rule[0]) == 11 && len(rule[1]) == 19 {
			base := parseTile3(rule[0])
			to := parseTile4(rule[1])
			for _, from := range base.flips() {
				book.rules3[from] = to
			}
		} else {
			return nil, fmt.Errorf("invalid rule: %v", rule)
		}
	}
	return book, nil
}

func (book *ruleBook) expand2(src *bitmap) (dst *bitmap) {
	size := src.size / 2
	dst = newBitmap(size * 3)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dst.put3(3*x, 3*y, book.rules2[src.get2(2*x, 2*y)])
		}
	}
	return dst
}

func (book *ruleBook) expand3(src *bitmap) (dst *bitmap) {
	size := src.size / 3
	dst = newBitmap(size * 4)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			dst.put4(4*x, 4*y, book.rules3[src.get3(3*x, 3*y)])
		}
	}
	return dst
}

type tile2 uint32

func parseTile2(text string) (tile tile2) {
	tile |= tile2(text[0]&1) << 3
	tile |= tile2(text[1]&1) << 2
	tile |= tile2(text[3]&1) << 1
	tile |= tile2(text[4]&1) << 0
	return tile
}

func (src tile2) flips() (dst [8]tile2) {
	order := [8][4]int{
		{0, 1, 2, 3}, {0, 2, 1, 3}, {1, 0, 3, 2}, {2, 0, 3, 1},
		{3, 2, 1, 0}, {3, 1, 2, 0}, {2, 3, 0, 1}, {1, 3, 0, 2},
	}
	for b := 0; b < 4; b++ {
		if (src & 1) == 1 {
			for i := 0; i < 8; i++ {
				dst[i] |= 1 << order[i][b]
			}
		}
		src >>= 1
	}
	return dst
}

type tile3 uint32

func parseTile3(text string) (tile tile3) {
	tile |= tile3(text[0]&1) << 8
	tile |= tile3(text[1]&1) << 7
	tile |= tile3(text[2]&1) << 6
	tile |= tile3(text[4]&1) << 5
	tile |= tile3(text[5]&1) << 4
	tile |= tile3(text[6]&1) << 3
	tile |= tile3(text[8]&1) << 2
	tile |= tile3(text[9]&1) << 1
	tile |= tile3(text[10]&1) << 0
	return tile
}

func (src tile3) flips() (dst [8]tile3) {
	order := [8][9]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8}, {0, 3, 6, 1, 4, 7, 2, 5, 8},
		{2, 1, 0, 5, 4, 3, 8, 7, 6}, {6, 3, 0, 7, 4, 1, 8, 5, 2},
		{8, 7, 6, 5, 4, 3, 2, 1, 0}, {8, 5, 2, 7, 4, 1, 6, 3, 0},
		{6, 7, 8, 3, 4, 5, 0, 1, 2}, {2, 5, 8, 1, 4, 7, 0, 3, 6},
	}
	for b := 0; b < 9; b++ {
		if (src & 1) == 1 {
			for i := 0; i < 8; i++ {
				dst[i] |= 1 << order[i][b]
			}
		}
		src >>= 1
	}
	return dst
}

type tile4 uint32

func parseTile4(text string) (tile tile4) {
	tile |= tile4(text[0]&1) << 15
	tile |= tile4(text[1]&1) << 14
	tile |= tile4(text[2]&1) << 13
	tile |= tile4(text[3]&1) << 12
	tile |= tile4(text[5]&1) << 11
	tile |= tile4(text[6]&1) << 10
	tile |= tile4(text[7]&1) << 9
	tile |= tile4(text[8]&1) << 8
	tile |= tile4(text[10]&1) << 7
	tile |= tile4(text[11]&1) << 6
	tile |= tile4(text[12]&1) << 5
	tile |= tile4(text[13]&1) << 4
	tile |= tile4(text[15]&1) << 3
	tile |= tile4(text[16]&1) << 2
	tile |= tile4(text[17]&1) << 1
	tile |= tile4(text[18]&1) << 0
	return tile
}

type bitmap struct {
	size    int
	rowSize int
	data    []uint32
}

func newBitmap(size int) *bitmap {
	rowSize := (size + 31) / 32
	return &bitmap{
		size:    size,
		rowSize: rowSize,
		data:    make([]uint32, size*rowSize),
	}
}

func (b *bitmap) get2(x, y int) (v tile2) {
	xu, xb := x>>5, x&0b11111
	for row := 0; row < 2; row++ {
		off := (y + row) * b.rowSize
		v <<= 2
		v |= tile2(b.data[off+xu]>>(30-xb)) & 0b11
	}
	return v
}

func (b *bitmap) get3(x, y int) (v tile3) {
	xu, xb := x>>5, x&0b11111
	if xb <= 29 { // fast path: contiguous
		for row := 0; row < 3; row++ {
			off := (y + row) * b.rowSize
			v <<= 3
			v |= tile3(b.data[off+xu]>>(29-xb)) & 0b111
		}
	} else { // slow path: straddling the unit boundary
		split := 32 - xb // bits in xu (2 or 1)
		for row := 0; row < 3; row++ {
			off := (y + row) * b.rowSize
			v <<= split
			v |= tile3(b.data[off+xu] & uint32(split|1))
			v <<= 3 - split
			v |= tile3(b.data[off+xu+1] >> (29 + split))
		}
	}
	return v
}

func (b *bitmap) put3(x, y int, v tile3) {
	xu, xb := x>>5, x&0b11111
	if xb <= 29 { // fast path: contiguous
		for row := 0; row < 3; row++ {
			off := (y + row) * b.rowSize
			rv := (uint32(v) >> 6) & 0b111
			v <<= 3
			b.data[off+xu] |= rv << (29 - xb)
		}
	} else { // slow path: straddling the unit boundary
		split := 32 - xb
		for row := 0; row < 3; row++ {
			off := (y + row) * b.rowSize
			rv := (uint32(v) >> 6) & 0b111
			v <<= 3
			b.data[off+xu] |= rv >> (3 - split)
			b.data[off+xu+1] |= (rv & uint32((3-split)|1)) << (29 + split)
		}
	}
}

func (b *bitmap) put4(x, y int, v tile4) {
	xu, xb := x>>5, x&0b11111
	for row := 0; row < 4; row++ {
		off := (y + row) * b.rowSize
		rv := (uint32(v) >> 12) & 0b1111
		v <<= 4
		b.data[off+xu] |= rv << (28 - xb)
	}
}

func (b *bitmap) popCount() (count int) {
	for _, v := range b.data {
		count += bits.OnesCount32(v)
	}
	return count
}

func (b *bitmap) print() {
	for y := 0; y < b.size; y++ {
		for xo, v := range b.data[y*b.rowSize : (y+1)*b.rowSize] {
			for x := 0; x < 32 && (xo<<5)+x < b.size; x++ {
				util.Diagf("%c", [2]rune{'.', '#'}[v>>31])
				v <<= 1
			}
		}
		util.Diagln()
	}
}
