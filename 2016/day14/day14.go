// Copyright 2023 Google LLC
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

// Package day14 solves AoC 2016 day 14.
package day14

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util"
)

func init() {
	glue.RegisterSolver(2016, 14, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line, got %d", len(lines))
	}
	salt := lines[0]
	p1 := findKey(newSimpleStreamer(salt), 64)
	p2 := findKey(newStretchedStreamer(salt, 2016), 64)
	return glue.Ints(p1, p2), nil
}

func findKey(ks keystreamer, keyIndex int) int {
	for keyIndex > 0 {
		candidate := ks.peek(0)
		if triple := candidate.findTriple(); triple != noTriple {
			for offset := 1; offset <= 1000; offset++ {
				if ks.peek(offset).matchPentad(triple) {
					keyIndex--
					break
				}
			}
		}
		ks.pop()
	}
	return ks.count() - 1
}

type hash [md5.Size]byte

const noTriple byte = 0x10

func (h hash) findTriple() byte {
	for i := 0; i < md5.Size-1; i++ {
		half := h[i]<<4 | h[i+1]>>4
		if h[i] == half || half == h[i+1] {
			return half & 0xf
		}
	}
	return noTriple
}

func (h hash) matchPentad(key byte) bool {
	kk := key<<4 | key
	for i := 0; i < md5.Size-2; i++ {
		if h[i] == kk && h[i+1] == kk && h[i+2]>>4 == key {
			return true
		}
		if h[i]&0x0f == key && h[i+1] == kk && h[i+2] == kk {
			return true
		}
	}
	return false
}

type keystreamer interface {
	count() int
	peek(offset int) hash
	pop()
}

type simpleStreamer struct {
	q        util.Queue[hash]
	cnt      int // index of q[0]
	saltSize int
	buf      [32]byte
}

func newSimpleStreamer(salt string) *simpleStreamer {
	ks := &simpleStreamer{saltSize: len(salt), q: util.MakeQueue[hash](1024)}
	for i, b := range []byte(salt) {
		ks.buf[i] = b
	}
	return ks
}

func (ks *simpleStreamer) count() int {
	return ks.cnt
}

func (ks *simpleStreamer) peek(offset int) hash {
	for offset >= ks.q.Len() {
		idx := int64(ks.cnt + offset)
		idxSize := len(strconv.AppendInt(ks.buf[ks.saltSize:ks.saltSize], idx, 10))
		ks.q.Push(md5.Sum(ks.buf[:ks.saltSize+idxSize]))
	}
	return ks.q.Index(offset)
}

func (ks *simpleStreamer) pop() {
	if !ks.q.Empty() {
		ks.q.Pop()
	}
	ks.cnt++
}

type stretchedStreamer struct {
	q        util.Queue[hash]
	cnt      int // index of q[0]
	saltSize int
	buf      [32]byte
	stretch  int
	hashBuf  [2 * md5.Size]byte
}

func newStretchedStreamer(salt string, stretch int) *stretchedStreamer {
	ks := &stretchedStreamer{saltSize: len(salt), stretch: stretch, q: util.MakeQueue[hash](1024)}
	for i, b := range []byte(salt) {
		ks.buf[i] = b
	}
	return ks
}

func (ks *stretchedStreamer) count() int {
	return ks.cnt
}

func (ks *stretchedStreamer) peek(offset int) hash {
	for offset >= ks.q.Len() {
		idx := int64(ks.cnt + offset)
		idxSize := len(strconv.AppendInt(ks.buf[ks.saltSize:ks.saltSize], idx, 10))
		h := md5.Sum(ks.buf[:ks.saltSize+idxSize])
		for i := 0; i < ks.stretch; i++ {
			hex.Encode(ks.hashBuf[:], h[:])
			h = md5.Sum(ks.hashBuf[:])
		}
		ks.q.Push(h)
	}
	return ks.q.Index(offset)
}

func (ks *stretchedStreamer) pop() {
	if !ks.q.Empty() {
		ks.q.Pop()
	}
	ks.cnt++
}
