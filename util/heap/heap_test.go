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
