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

// Package day16 solves AoC 2021 day 16.
package day16

import (
	"encoding/hex"
	"fmt"

	"github.com/fis/aoc/glue"
)

func init() {
	glue.RegisterSolver(2021, 16, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line, got %d", len(lines))
	}
	p, _ := parsePacket(bitReaderHex(lines[0]))
	p1 := p.versionSum()
	p2 := p.eval()
	return glue.Ints(p1, p2), nil
}

type opCode int

const (
	opSum  opCode = 0
	opProd opCode = 1
	opMin  opCode = 2
	opMax  opCode = 3
	opLit  opCode = 4
	opGt   opCode = 5
	opLt   opCode = 6
	opEq   opCode = 7
)

type packet struct {
	ver int
	op  opCode
	lit int
	sub []packet
}

func (p packet) versionSum() (vs int) {
	vs = p.ver
	for _, sub := range p.sub {
		vs += sub.versionSum()
	}
	return vs
}

func (p packet) eval() (v int) {
	switch p.op {
	case opSum:
		for _, sub := range p.sub {
			v += sub.eval()
		}
		return v
	case opProd:
		v = 1
		for _, sub := range p.sub {
			v *= sub.eval()
		}
		return v
	case opMin:
		v = p.sub[0].eval()
		for _, sub := range p.sub[1:] {
			if sv := sub.eval(); sv < v {
				v = sv
			}
		}
		return v
	case opMax:
		v = p.sub[0].eval()
		for _, sub := range p.sub[1:] {
			if sv := sub.eval(); sv > v {
				v = sv
			}
		}
		return v
	case opLit:
		return p.lit
	case opGt:
		if p.sub[0].eval() > p.sub[1].eval() {
			return 1
		} else {
			return 0
		}
	case opLt:
		if p.sub[0].eval() < p.sub[1].eval() {
			return 1
		} else {
			return 0
		}
	case opEq:
		if p.sub[0].eval() == p.sub[1].eval() {
			return 1
		} else {
			return 0
		}
	}
	return -1
}

func parsePacket(br *bitReader) (p packet, totalLength int) {
	p.ver = br.take(3)
	p.op = opCode(br.take(3))
	totalLength += 6

	if p.op == opLit {
		for {
			chunk := br.take(5)
			totalLength += 5
			p.lit <<= 4
			p.lit |= chunk & 0b1111
			if chunk&0b10000 == 0 {
				break
			}
		}
		return p, totalLength
	}

	var bodyLength int
	inPackets := br.take(1) == 1
	totalLength++
	if inPackets {
		bodyLength = br.take(11)
		totalLength += 11
	} else {
		bodyLength = br.take(15)
		totalLength += 15
	}
	for bodyLength > 0 {
		sub, subLength := parsePacket(br)
		totalLength += subLength
		p.sub = append(p.sub, sub)
		if inPackets {
			bodyLength--
		} else {
			bodyLength -= subLength
		}
	}

	return p, totalLength
}

type bitReader struct {
	data []byte
	at   int
}

func bitReaderHex(str string) (br *bitReader) {
	src := []byte(str)
	if len(src)&1 == 1 {
		src = append(src, '0')
	}
	dst := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(dst, src)
	return &bitReader{data: dst, at: 0}
}

func (br *bitReader) take(bits int) (val int) {
	for br.at+bits >= 8 {
		n := 8 - br.at
		val <<= n
		val |= int(br.data[0]) & ((1 << n) - 1)
		bits -= n
		br.data, br.at = br.data[1:], 0
	}
	if bits == 0 {
		return val
	}
	val <<= bits
	val |= (int(br.data[0]) >> (8 - br.at - bits)) & ((1 << bits) - 1)
	br.at += bits
	return val
}
