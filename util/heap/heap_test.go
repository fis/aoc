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

package heap

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	data := []Item[string]{
		{5, "frump"},
		{2, "baz"},
		{1, "bar"},
		{3, "quux"},
		{0, "foo"},
		{4, "zuul"},
	}
	h := Heap[string](nil)
	for _, item := range data {
		h.Push(item.P, item.Value)
	}
	got := []string(nil)
	for len(h) > 0 {
		_, value := h.Pop()
		got = append(got, value)
	}
	want := "[foo bar baz quux zuul frump]"
	if fmt.Sprint(got) != want {
		t.Errorf("got %v want %s", got, want)
	}
}

func TestHeapify(t *testing.T) {
	h := Heap[string]{
		{5, "frump"},
		{2, "baz"},
		{1, "bar"},
		{3, "quux"},
		{0, "foo"},
		{4, "zuul"},
	}
	h.Heapify()
	got := []string(nil)
	for len(h) > 0 {
		_, value := h.Pop()
		got = append(got, value)
	}
	want := "[foo bar baz quux zuul frump]"
	if fmt.Sprint(got) != want {
		t.Errorf("got %v want %s", got, want)
	}
}
