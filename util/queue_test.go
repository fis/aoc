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

package util

import "testing"

func TestQueue(t *testing.T) {
	q := MakeQueue[int](8)
	// [>__  __  __  __  __  __  __  __]
	for i := 0; i <= 6; i++ {
		q.Push(i)
	}
	// [>00  01  02  03  04  05  06  __]
	for i := 0; i <= 4; i++ {
		if j := q.Pop(); j != i {
			t.Fatal(i, j)
		}
	}
	// [ __  __  __  __  __ >05  06  __]
	for i := 7; i <= 12; i++ {
		q.Push(i)
	}
	// [ 08  09  10  11  12 >05  06  07]
	for i := 5; i <= 9; i++ {
		if j := q.Pop(); j != i {
			t.Fatal(i, j)
		}
	}
	// [ __  __ >10  11  12  __  __  __]
	for i := 13; i <= 17; i++ {
		q.Push(i)
	}
	// [ 16  17 >10  11  12  13  14  15]
	q.Push(18)
	// [>10  11  12  13  14  15  16  17  18  __  __  __  __  __  __  __]
	for i := 10; i <= 17; i++ {
		if j := q.Pop(); j != i {
			t.Fatal(i, j)
		}
	}
	// [ __  __  __  __  __  __  __  __ >18  __  __  __  __  __  __  __]
	for i := 19; i <= 27; i++ {
		q.Push(i)
	}
	// [ 26  27  __  __  __  __  __  __ >18  19  20  21  22  23  24  25]
	for i := 18; i <= 27; i++ {
		if j := q.Pop(); j != i {
			t.Fatal(i, j)
		}
	}
	// [ __  __ >__  __  __  __  __  __  __  __  __  __  __  __  __  __]
}
