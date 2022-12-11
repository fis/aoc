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

// Package day05 solves AoC 2016 day 5.
package day05

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/fis/aoc/glue"
	"golang.org/x/exp/slices"
)

func init() {
	glue.RegisterSolver(2016, 5, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line of input, got %d", len(lines))
	}
	p1, p2 := search(lines[0], 8)
	return []string{p1, p2}, nil
}

var hexDigits = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

const (
	parallelism = 8
	chunkBits   = 16
	chunkSize   = 1 << chunkBits
)

func search(prefix string, count int) (string, string) {
	sortedCh, doneCh := make(chan [][2]byte), make(chan struct{})
	go controller(prefix, sortedCh, doneCh)

	p1, p2 := make([]byte, 0, count), make([]byte, count)
	found := 0
outer:
	for results := range sortedCh {
		for _, pair := range results {
			if len(p1) < count {
				p1 = append(p1, hexDigits[pair[0]])
			}
			if pos := int(pair[0]); pos < count && p2[pos] == 0 {
				p2[pos] = hexDigits[pair[1]]
				found++
				if found >= count {
					close(doneCh)
					break outer
				}
			}
		}
	}

	return string(p1), string(p2)
}

type result struct {
	start int64
	found [][2]byte
}

func controller(prefix string, sortedCh chan<- [][2]byte, doneCh <-chan struct{}) {
	var (
		next    int64 = 0
		expect  int64 = 0
		pending []result
		chunkCh = make(chan int64)
		foundCh = make(chan result)
	)
	for i := 0; i < parallelism; i++ {
		go searchChunk(prefix, chunkCh, foundCh)
	}
loop:
	for {
		select {
		case <-doneCh:
			break loop
		case chunkCh <- next:
			next += chunkSize
		case got := <-foundCh:
			if got.start > expect {
				pending = append(pending, got)
				continue
			}
			sortedCh <- got.found
			expect += chunkSize
			slices.SortFunc(pending, func(a, b result) bool { return a.start < b.start })
			for len(pending) > 0 && pending[0].start == expect {
				sortedCh <- pending[0].found
				expect += chunkSize
				pending = pending[1:]
			}
		}
	}
	running := parallelism
	for running > 0 {
		select {
		case chunkCh <- -1:
			running--
		case <-foundCh:
			// no longer needed
		}
	}
}

func searchChunk(prefix string, chunkCh <-chan int64, foundCh chan<- result) {
	buf := append([]byte(prefix), make([]byte, 16)...)
	for start := range chunkCh {
		if start < 0 {
			return
		}
		var found [][2]byte
		for i := int64(0); i < chunkSize; i++ {
			in := strconv.AppendInt(buf[:len(prefix)], start+i, 10)
			out := md5.Sum(in)
			if out[0] == 0 && out[1] == 0 && (out[2]&0xf0) == 0 {
				found = append(found, [2]byte{out[2] & 0x0f, out[3] >> 4})
			}
		}
		foundCh <- result{start: start, found: found}
	}
}
