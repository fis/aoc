// Package fn contains the sort of non-Go-like, occasionally higher-order, utility functions you might find in a functional language.
package fn

import "golang.org/x/exp/constraints"

// Sum returns the sum of a slice of integers. The sum of an empty slice is 0.
func Sum[S ~[]E, E constraints.Integer](s S) (result E) {
	for _, e := range s {
		result += e
	}
	return result
}

// SumF returns the sum of the results of applying a function to a slice. The sum of an empty slice is 0.
func SumF[S ~[]I, F ~func(I) O, I any, O constraints.Integer](s S, f F) (result O) {
	for _, e := range s {
		result += f(e)
	}
	return result
}

// Prod returns the product of a slice of integers. The product of an empty slice is 1.
func Prod[S ~[]E, E constraints.Integer](s S) (result E) {
	result = 1
	for _, e := range s {
		result *= e
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

// If is an expression conditional, similar to the C `?:` operator (except both branches are evaluated).
func If[T any](c bool, t, f T) T {
	if c {
		return t
	}
	return f
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

// MaxF returns the largest result of applying a function to a slice.
func MaxF[S ~[]I, F ~func(I) O, I any, O constraints.Ordered](s S, f F) (result O) {
	result = f(s[0])
	for _, e := range s[1:] {
		if o := f(e); o > result {
			result = o
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

// MapRange returns a slice with the results of calling function on a range of integers.
func MapRange[F ~func(I) E, I constraints.Integer, E any](start, end I, f F) (out []E) {
	n := end - start
	out = make([]E, n)
	for i := start; i < end; i++ {
		out[i-start] = f(i)
	}
	return out
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

// All returns true if the function f returns true for all elements of slice s. Returns true for an empty slice.
func All[S ~[]E, F ~func(E) bool, E any](s S, f F) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

// Any returns true if the function f returns true for at least one element of slice s. Returns false for an empty slice.
func Any[S ~[]E, F ~func(E) bool, E any](s S, f F) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}
