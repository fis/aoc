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

package util

import "fmt"

// Queue is a simple array-backed ring buffer implementation of an unbounded queue.
//
// Using it in place of a simple slice might help avoid some copying,
// at least in a scenario where the "active" size stays roughly the same
// but a lot of elements go through the queue.
type Queue[T any] struct {
	// The backing store.
	//
	// len(q) is the length of the queue, even when it doesn't start at the front. (This is silly.)
	// cap(q) is the size of the backing array, and always a power of 2.
	q []T
	// Current head of the queue as an index into q.
	head int
}

// MakeQueue makes a new, empty queue of the given initial size, which must be a power of 2.
func MakeQueue[T any](size int) Queue[T] {
	return Queue[T]{
		q:    make([]T, 0, size),
		head: 0,
	}
}

// Len returns the number of items currently in the queue.
func (q Queue[T]) Len() int {
	return len(q.q)
}

// Empty returns true if the queue has no items.
func (q Queue[T]) Empty() bool {
	return len(q.q) == 0
}

// Index returns the i'th item in the queue.
// The head of the queue (least recently pushed item) is at index 0.
func (q Queue[T]) Index(i int) T {
	i = (q.head + i) & (cap(q.q) - 1)
	return q.q[:i+1][i]
}

// Slice returns a segment of the contents of the queue as a slice.
// The semantics of the parameters are the same as the `s[low:high]` expression on a slice.
// If possible, the returned slice will share backing array with the queue.
// However, if the selected region wraps around, a copy will be made.
func (q *Queue[T]) Slice(low, high int) []T {
	if low < 0 || high < 0 || low > len(q.q) || high > len(q.q) || low > high {
		panic(fmt.Sprintf("runtime error: queue slice bounds violate 0 <= %d <= %d <= %d", low, high, len(q.q)))
	}
	n := high - low
	low = (q.head + low) & (cap(q.q) - 1)
	if low+n <= cap(q.q) {
		return q.q[low : low+n]
	}
	return append(q.q[low:cap(q.q)], q.q[:n-(cap(q.q)-low)]...) // guaranteed to make a copy
}

// Push adds an item at the end of the queue.
func (q *Queue[T]) Push(t T) {
	n := len(q.q)
	if n == cap(q.q) {
		newQ := make([]T, n, 2*n)
		if q.head == 0 {
			copy(newQ, q.q)
		} else {
			n1 := n - q.head
			copy(newQ[:n1], q.q[q.head:])
			copy(newQ[n1:], q.q[:q.head])
		}
		q.q, q.head = newQ, 0
	}
	i := (q.head + n) & (cap(q.q) - 1)
	q.q[:i+1][i] = t
	q.q = q.q[:n+1]
}

// Pop removes (and returns) the first (least recently added) item in the queue.
// The queue must not be empty.
func (q *Queue[T]) Pop() T {
	if len(q.q) == 0 {
		panic("runtime error: pop of an empty queue")
	}
	head := q.head
	q.head = (head + 1) & (cap(q.q) - 1)
	q.q = q.q[:len(q.q)-1]
	return q.q[:head+1][head]
}
