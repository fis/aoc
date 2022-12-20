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

package obs

const bucketSpan = 32

type bucketQ[T any] struct {
	count   int
	at      int
	buckets [bucketSpan]struct {
		prio  int
		items []T
	}
}

func (q *bucketQ[T]) len() int {
	return q.count
}

func (q *bucketQ[T]) pop() (prio int, item T) {
	q.count--
	for len(q.buckets[q.at].items) == 0 {
		q.at = (q.at + 1) & (bucketSpan - 1)
	}
	items := q.buckets[q.at].items
	item = items[len(items)-1]
	q.buckets[q.at].items = items[:len(items)-1]
	return q.buckets[q.at].prio, item
}

func (q *bucketQ[T]) push(prio int, item T) {
	q.count++
	i := prio & (bucketSpan - 1)
	q.buckets[i].prio = prio
	q.buckets[i].items = append(q.buckets[i].items, item)
}
