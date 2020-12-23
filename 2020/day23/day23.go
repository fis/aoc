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

// Package day23 solves AoC 2020 day 23.
package day23

import (
	"fmt"

	"github.com/fis/aoc-go/glue"
)

func init() {
	glue.RegisterSolver(2020, 23, glue.LineSolver(solve))
}

func solve(lines []string) ([]int, error) {
	if len(lines) != 1 {
		return nil, fmt.Errorf("expected one line, got %d", len(lines))
	}

	r1, err := newRing(lines[0])
	if err != nil {
		return nil, err
	}
	for i := 0; i < 100; i++ {
		r1 = r1.move()
	}
	p1 := r1.key()

	r2, err := newBigRing(lines[0])
	if err != nil {
		return nil, err
	}
	for i := 0; i < 10000000; i++ {
		r2.move()
	}
	p2 := r2.key()

	return []int{p1, p2}, nil
}

type ring uint64

func newRing(s string) (r ring, err error) {
	if len(s) != 9 {
		return 0, fmt.Errorf("invalid labels (expected 9): %q", s)
	}
	r = 0
	for i := 0; i < 9; i++ {
		b := s[i]
		if b < '1' || b > '9' {
			return 0, fmt.Errorf("invalid label: %v", b)
		}
		r |= ring(b-'1') << (4 * i)
	}
	return r, nil
}

func (r ring) String() string {
	labels := make([]byte, 9)
	for i := 0; i < 9; i++ {
		labels[i] = '1' + byte((r>>(4*i))&0xf)
	}
	return string(labels)
}

func (r ring) move() ring {
	cl := r & 0xf
	pick := (r >> 4) & 0xfff
	dl := (cl - 1 + 9) % 9
	for dl == pick&0xf || dl == (pick>>4)&0xf || dl == (pick>>8)&0xf {
		dl = (dl - 1 + 9) % 9
	}
	dp := 16
	for (r>>dp)&0xf != dl {
		dp += 4
	}
	b1p, b1w := 16, dp-16+4
	b2p, b2w := dp+4, 32-dp
	b1m, b2m := (ring(1)<<b1w)-1, (ring(1)<<b2w)-1
	return ((r >> b1p) & b1m) | (pick << b1w) | (((r >> b2p) & b2m) << (12 + b1w)) | (cl << 32)
}

func (r ring) key() (n int) {
	one := 0
	for (r>>one)&0xf != 0 {
		one += 4
	}
	for at := (one + 4) % 36; at != one; at = (at + 4) % 36 {
		n = 10*n + int((r>>at)&0xf) + 1
	}
	return n
}

type bigRing struct {
	cur  *bigRingCup
	cups []*bigRingCup
}

type bigRingCup struct {
	label int
	next  *bigRingCup
}

func newBigRing(s string) (r *bigRing, err error) {
	r = &bigRing{cups: make([]*bigRingCup, 1000000)}
	prev, max := &r.cur, 0
	for _, b := range s {
		if b < '1' || b > '9' {
			return nil, fmt.Errorf("invalid label: %v", b)
		}
		i := int(b) - '1'
		cup := &bigRingCup{label: i}
		r.cups[i] = cup
		*prev, prev = cup, &cup.next
		if i > max {
			max = i
		}
	}
	for i := max + 1; i <= 999999; i++ {
		cup := &bigRingCup{label: i}
		r.cups[i] = cup
		*prev, prev = cup, &cup.next
	}
	*prev = r.cur
	return r, nil
}

func (r *bigRing) move() {
	p1, p2, p3 := r.cur.next, r.cur.next.next, r.cur.next.next.next
	r.cur.next = p3.next
	dl := (r.cur.label - 1 + 1000000) % 1000000
	for dl == p1.label || dl == p2.label || dl == p3.label {
		dl = (dl - 1 + 1000000) % 1000000
	}
	d := r.cups[dl]
	d.next, p3.next = p1, d.next
	r.cur = r.cur.next
}

func (r bigRing) key() int {
	k0 := r.cups[0]
	k1, k2 := k0.next, k0.next.next
	return (k1.label + 1) * (k2.label + 1)
}

/*
Big ring "symbolic" simulation, looking for shortcuts.

0:
(3) 8 9 1 2 5 4 6 7 10..1000000
    ~~~~~ *
1:
(2) 8 9 1 5 4 6 7 10..1000000 3
    ~~~~~             *
2:
(5) 4 6 7 10..1000000 8 9 1 3 2
    ~~~~~                   *
3:
(10) 11..13 14..1000000 8 9 1 3 4 6 7 2 5
     ~~~~~~               *
4:
(14) 15..17 18..1000000 8 9 11..13 1 3 4 6 7 2 5 10
     ~~~~~~                     *
5:
(18) 19..21 22..1000000 8 9 11..13 15..17 1 3 4 6 7 2 5 10 14
		 ~~~~~~                            *
...
250000:
(999998) 999999 1000000 8 9 11..13/15..17//999995..999997 1 3 4 6 7 2 5 10/14//999994
         ~~~~~~~~~~~~~~~~                          *
250001:
(9) 11 12 13 15..17/19..21//999995..999997 999999 1000000 8 1 3 4 6 7 2 5 10/14//999998
    ~~~~~~~~                                              *
250002:
(15) 16 17 19 20 21 23..25/27..29//999995..999997 999999 1000000 8 11..13 1 3 4 6 7 2 5 10/14//999998 9
...
Seems to get pretty shuffled up.
*/
