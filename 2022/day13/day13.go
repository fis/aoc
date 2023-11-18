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

// Package day13 solves AoC 2022 day 13.
package day13

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/glue"
	"github.com/fis/aoc/util/fn"
)

func init() {
	glue.RegisterSolver(2022, 13, glue.LineSolver(solve))
}

func solve(lines []string) ([]string, error) {
	packets, err := parsePackets(lines)
	if err != nil {
		return nil, err
	}
	p1 := part1(packets)
	p2 := part2(packets)
	return glue.Ints(p1, p2), nil
}

func part1(packets []list) (sum int) {
	for i := 0; i*2+1 < len(packets); i++ {
		if cmp(packets[i*2], packets[i*2+1]) < 0 {
			sum += i + 1
		}
	}
	return sum
}

func part2(packets []list) int {
	d1, d2 := singletonList(2), singletonList(6)
	packets = append(packets, d1, d2)
	slices.SortFunc(packets, func(a, b list) int { return cmp(a, b) })
	i1 := slices.IndexFunc(packets, func(e list) bool { return cmp(e, d1) == 0 }) + 1
	i2 := slices.IndexFunc(packets, func(e list) bool { return cmp(e, d2) == 0 }) + 1
	return i1 * i2
}

func cmp(left, right list) int {
	llen, rlen := left.len(), right.len()
	n := min(llen, rlen)
	for i := 0; i < n; i++ {
		lt, rt := left.itemType(i), right.itemType(i)
		if lt == numItem && rt == numItem {
			switch ln, rn := left.num(i), right.num(i); {
			case ln < rn:
				return -1
			case ln > rn:
				return 1
			}
		} else {
			c := cmp(left.list(i), right.list(i))
			if c != 0 {
				return c
			}
		}
	}
	switch {
	case llen < rlen:
		return -1
	case llen > rlen:
		return 1
	}
	return 0
}

func parsePackets(lines []string) (packets []list, err error) {
	for _, line := range lines {
		if line == "" {
			continue
		}
		l, tail, err := parseList(line)
		if err != nil {
			return nil, err
		} else if tail != "" {
			return nil, fmt.Errorf("unexpected cruft at end of list: %q", tail)
		}
		packets = append(packets, l)
	}
	return packets, nil
}

// list represents a heterogenous list of either numbers or other similar lists.
type list interface {
	// len returns the length of the list.
	len() int
	// itemType returns the type of the item at index i: either `numItem` or `listItem`.
	itemType(i int) itemType
	// num returns the number at index i. It must only be called if there is a number at that index.
	num(i int) int
	// list returns the sublist at index i.
	// If there is actually a number at that index, it returns that number wrapped in a single-element list.
	list(i int) list

	fmt.Stringer
}

type itemType int

const (
	numItem itemType = iota
	listItem
)

func parseList(data string) (l list, tail string, err error) {
	if len(data) == 0 || data[0] != '[' {
		return nil, "", fmt.Errorf("expecting [")
	}
	data = data[1:]

	if len(data) > 0 && data[0] == ']' {
		return emptyList{}, data[1:], nil
	}

	first := -1
	if len(data) > 0 && data[0] >= '0' && data[0] <= '9' {
		first, data = parseNum(data)
		switch {
		case len(data) == 0:
			return nil, "", fmt.Errorf("unexpected EOF inside list")
		case data[0] == ']':
			return singletonList(first), data[1:], nil
		case data[0] != ',':
			return nil, "", fmt.Errorf("expecting , or ] but got %c", data[0])
		}
		data = data[1:]
	}

	var dl dataList
	if first >= 0 {
		dl.order = append(dl.order, 0)
		dl.nums = append(dl.nums, first)
	}
	for len(data) > 0 {
		switch {
		case data[0] >= '0' && data[0] <= '9':
			var n int
			n, data = parseNum(data)
			dl.order = append(dl.order, len(dl.nums))
			dl.nums = append(dl.nums, n)
		case data[0] == '[':
			var l list
			l, data, err = parseList(data)
			if err != nil {
				return nil, "", err
			}
			dl.order = append(dl.order, -len(dl.lists)-1)
			dl.lists = append(dl.lists, l)
		}
		if len(data) == 0 {
			break
		}
		switch {
		case data[0] == ']':
			return dl, data[1:], nil
		case data[0] != ',':
			return nil, "", fmt.Errorf("expecting , or ] but got %c", data[0])
		}
		data = data[1:]
	}
	return nil, "", fmt.Errorf("unexpected EOF inside list")
}

func parseNum(data string) (n int, tail string) {
	for len(data) > 0 && data[0] >= '0' && data[0] <= '9' {
		n = 10*n + int(data[0]-'0')
		data = data[1:]
	}
	return n, data
}

type dataList struct {
	order []int
	nums  []int
	lists []list
}

func (dl dataList) len() int                { return len(dl.order) }
func (dl dataList) itemType(i int) itemType { return fn.If(dl.order[i] >= 0, numItem, listItem) }
func (dl dataList) num(i int) int           { return dl.nums[dl.order[i]] }

func (dl dataList) list(i int) list {
	ord := dl.order[i]
	if ord >= 0 {
		return singletonList(dl.nums[ord])
	}
	return dl.lists[-ord-1]
}

func (dl dataList) String() string {
	var s strings.Builder
	s.WriteByte('[')
	for i := 0; i < dl.len(); i++ {
		if i > 0 {
			s.WriteByte(',')
		}
		if dl.itemType(i) == numItem {
			s.WriteString(strconv.Itoa(dl.num(i)))
		} else {
			s.WriteString(dl.list(i).String())
		}
	}
	s.WriteByte(']')
	return s.String()
}

type singletonList int

func (singletonList) len() int              { return 1 }
func (singletonList) itemType(int) itemType { return numItem }
func (sl singletonList) num(int) int        { return int(sl) }
func (sl singletonList) list(int) list      { return sl }
func (sl singletonList) String() string     { return fmt.Sprintf("(%d)", sl) }

type emptyList struct{}

func (emptyList) len() int              { return 0 }
func (emptyList) itemType(int) itemType { panic("not supported") }
func (emptyList) num(int) int           { panic("not supported") }
func (emptyList) list(int) list         { panic("not supported") }
func (emptyList) String() string        { return "{}" }
