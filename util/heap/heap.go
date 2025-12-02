// Copyright 2024 Google LLC
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

// Package heap implements a basic binary min-heap on top of a slice.
package heap

type Heap[T any] []Item[T]

type Item[T any] struct {
	P int
	Value T
}

func (h *Heap[T]) Push(prio int, value T) {
	data := append(*h, Item[T]{})
	pos := len(data) - 1
	for pos > 0 {
		p := (pos-1) >> 1
		if data[p].P <= prio {
			break
		}
		pos, data[pos] = p, data[p]
	}
	data[pos] = Item[T]{prio, value}
	*h = data
}

func (h *Heap[T]) Pop() (prio int, value T) {
	data := *h
	prio, value = data[0].P, data[0].Value

	if len(data) == 1 {
		*h = data[:0]
		return prio, value
	}

	np, pos, data := data[len(data)-1].P, 0, data[:len(data)-1]
	for pos<<1|1 < len(data) {
		c := pos<<1|1
		if np <= data[c].P && (c+1 >= len(data) || np <= data[c+1].P) {
			break
		}
		if c+1 >= len(data) || data[c].P <= data[c+1].P {
			pos, data[pos] = c, data[c]
		} else {
			pos, data[pos] = c+1, data[c+1]
		}
	}
	data[pos] = Item[T]{np, (*h)[len(data)].Value}
	*h = data

	return prio, value
}
