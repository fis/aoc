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
// limitations under the License.package day16

package util

// BucketQ is a priority queue using the "bucket queue" data structure.
type BucketQ[T any] struct {
	count   int
	at      int
	buckets []bucket[T] // [span]bucket[T]
}

type bucket[T any] struct {
	prio  int
	items []T
}

// NewBucketQ makes a new queue with the given maximum span, which must be a power of 2.
// The distance between the lowest and highest priority item in the queue must not exceed the span.
func NewBucketQ[T any](span int) *BucketQ[T] {
	return &BucketQ[T]{
		buckets: make([]bucket[T], span),
	}
}

// Len returns the number of elements currently in the queue.
func (q *BucketQ[T]) Len() int {
	return q.count
}

// Pop removes the lowest priority item currently in the queue.
func (q *BucketQ[T]) Pop() (prio int, item T) {
	q.count--
	for len(q.buckets[q.at].items) == 0 {
		q.at = (q.at + 1) & (len(q.buckets) - 1)
	}
	items := q.buckets[q.at].items
	item = items[len(items)-1]
	q.buckets[q.at].items = items[:len(items)-1]
	return q.buckets[q.at].prio, item
}

// Push adds an item to the queue with the given priority.
func (q *BucketQ[T]) Push(prio int, item T) {
	q.count++
	i := prio & (len(q.buckets) - 1)
	q.buckets[i].prio = prio
	q.buckets[i].items = append(q.buckets[i].items, item)
}
