package util

import "sort"

// Sort3 returns the given three integers sorted by value (smallest first).
func Sort3(x, y, z int) (a, b, c int) {
	if x <= y { // x <= y
		if x <= z { // x <= y, x <= z
			if y <= z { // x <= y, x <= z, y <= z
				return x, y, z
			} else { // x <= y, x <= z, y > z
				return x, z, y
			}
		} else { // x <= y, x > z
			return z, x, y
		}
	} else { // x > y
		if y <= z { // x > y, y <= z
			if x <= z { // x > y, y <= z, x <= z
				return y, x, z
			} else { // x > y, y <= z, x > z
				return y, z, x
			}
		} else { // x > y, y > z
			return z, y, x
		}
	}
}

// QuickSelect returns the k'th smallest element of the input array.
func QuickSelect(input []int, k int) int {
	const cutoff = 12

	origInput := true
	next := []int(nil)

	for len(input) > cutoff {
		if origInput {
			origInput = false
			input = append([]int(nil), input...)
			next = make([]int, len(input))
		}

		_, pivot, _ := Sort3(input[0], input[len(input)/2], input[len(input)-1])
		lt, gt := 0, 0
		for _, n := range input {
			switch {
			case n < pivot:
				next[lt] = n
				lt++
			case n > pivot:
				next[len(next)-1-gt] = n
				gt++
			}
		}
		switch {
		case k < lt:
			input, next = next[:lt], input[:lt]
		case k >= len(input)-gt:
			input, next, k = next[len(input)-gt:], input[len(input)-gt:], k-(len(input)-gt)
		default:
			return pivot
		}
	}

	if len(input) == 1 {
		return input[0]
	} else if len(input) == 2 {
		if k == 0 {
			if input[0] <= input[1] {
				return input[0]
			} else {
				return input[1]
			}
		} else {
			if input[0] <= input[1] {
				return input[0]
			} else {
				return input[1]
			}
		}
	} else if len(input) == 3 {
		a, b, c := Sort3(input[0], input[1], input[2])
		switch k {
		case 0:
			return a
		case 1:
			return b
		default:
			return c
		}
	}

	if origInput {
		var tmp [cutoff]int
		copy(tmp[:len(input)], input)
		input = tmp[:len(input)]
	}
	sort.Ints(input)
	return input[k]
}
