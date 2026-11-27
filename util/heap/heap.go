// Package heap implements a basic binary min-heap on top of a slice.
package heap

type Heap[T any] []Item[T]

type Item[T any] struct {
	P     int
	Value T
}

func (h *Heap[T]) Push(prio int, value T) {
	data := append(*h, Item[T]{})
	pos := len(data) - 1
	for pos > 0 {
		p := (pos - 1) >> 1
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
		c := pos<<1 | 1
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

func (h *Heap[T]) Heapify() {
	data := *h

	if len(data) < 2 {
		return
	}

	for root := (len(data) - 2) / 2; root >= 0; root-- {
		pos := root
		it := data[pos]
		for pos<<1|1 < len(data) {
			c := pos<<1 | 1
			if it.P <= data[c].P && (c+1 >= len(data) || it.P <= data[c+1].P) {
				break
			}
			if c+1 >= len(data) || data[c].P <= data[c+1].P {
				pos, data[pos] = c, data[c]
			} else {
				pos, data[pos] = c+1, data[c+1]
			}
		}
		if pos != root {
			data[pos] = it
		}
	}
}
