// Package fn contains the sort of non-Go-like, occasionally higher-order, utility functions you might find in a functional language.
package fn

import "golang.org/x/exp/constraints"

// Sum returns the sum of a slice of integers.
func Sum[S ~[]E, E constraints.Integer](s S) (result E) {
	for _, e := range s {
		result += e
	}
	return result
}

// CountIf returns the number of items in a slice that satisfy a predicate.
func CountIf[S ~[]E, F ~func(E) bool, E any](s S, f F) (count int) {
	for _, e := range s {
		if f(e) {
			count++
		}
	}
	return count
}

// Min returns the smallest value of a slice of some ordered type.
func Min[S ~[]E, E constraints.Ordered](s S) (result E) {
	result = s[0]
	for _, e := range s[1:] {
		if e < result {
			result = e
		}
	}
	return result
}

// Max returns the largest value of a slice of some ordered type.
func Max[S ~[]E, E constraints.Ordered](s S) (result E) {
	result = s[0]
	for _, e := range s[1:] {
		if e > result {
			result = e
		}
	}
	return result
}

// Head returns the first line of a slice, or a default value for an empty slice.
func Head[S ~[]E, E any](s S, def E) E {
	if len(s) > 0 {
		return s[0]
	}
	return def
}

// Map returns a new slice that contains the results of applying the given function to each element of the input slice.
func Map[S ~[]I, F ~func(I) O, I, O any](in S, f F) (out []O) {
	out = make([]O, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

// MapE is a variant of Map that allows the function to fail.
func MapE[S ~[]I, F ~func(I) (O, error), I, O any](in S, f F) (out []O, err error) {
	out = make([]O, len(in))
	for i, x := range in {
		o, err := f(x)
		if err != nil {
			return nil, err
		}
		out[i] = o
	}
	return out, nil
}

// Filter returns a new slice that contains just the elements matching a predicate.
func Filter[S ~[]E, P ~func(E) bool, E any](s S, p P) (out []E) {
	for _, e := range s {
		if p(e) {
			out = append(out, e)
		}
	}
	return out
}

// ForEach invokes a function (which must not return anything) for each element of a slice in order.
func ForEach[S ~[]E, F ~func(E), E any](s S, f F) {
	for _, e := range s {
		f(e)
	}
}