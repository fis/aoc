// Package fi contains the user-defined iterator counterparts of the fn package.
package fi

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// S returns the values of a slice as a sequence, like the standard slices.Values function.
func S[S ~[]E, E any](s S) iter.Seq[E] {
	return slices.Values(s)
}

// CS collects the values of a sequence as a slice, like the standard slices.Collect function.
func CS[E any](s iter.Seq[E]) []E {
	return slices.Collect(s)
}

// K returns the keys of a map as a sequence, like the standard maps.Keys function.
func K[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return maps.Keys(m)
}

// V returns the values of a map as a sequence, like the standard maps.Values function.
func V[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return maps.Values(m)
}

// Sum returns the sum of a sequence of integers. The sum of an empty sequence is 0.
func Sum[E integer](s iter.Seq[E]) (result E) {
	for e := range s {
		result += e
	}
	return result
}

// SumF returns the sum of the results of applying a function to a sequence. The sum of an empty sequence is 0.
func SumF[F ~func(I) O, I any, O integer](s iter.Seq[I], f F) (result O) {
	for e := range s {
		result += f(e)
	}
	return result
}

// Prod returns the product of a sequence of integers. The product of an empty slice is 1.
func Prod[E integer](s iter.Seq[E]) (result E) {
	result = 1
	for e := range s {
		result *= e
	}
	return result
}

// ProdF returns the product of the results of applying a function to a sequence. The product of an empty sequence is 1.
func ProdF[F ~func(I) O, I any, O integer](s iter.Seq[I], f F) (result O) {
	result = 1
	for e := range s {
		result *= f(e)
	}
	return result
}

// CountIf returns the number of items in a sequence that satisfy a predicate.
func CountIf[F ~func(E) bool, E any](s iter.Seq[E], f F) (count int) {
	for e := range s {
		if f(e) {
			count++
		}
	}
	return count
}

// Min returns the smallest value of a sequence of some ordered type.
func Min[E cmp.Ordered](s iter.Seq[E]) (result E) {
	next, stop := iter.Pull(s)
	defer stop()
	result, ok := next()
	if !ok {
		panic("Min of an empty sequence")
	}
	for {
		e, ok := next()
		if !ok {
			break
		}
		if e < result {
			result = e
		}
	}
	return result
}

// Max returns the largest value of a sequence of some ordered type.
func Max[E cmp.Ordered](s iter.Seq[E]) (result E) {
	next, stop := iter.Pull(s)
	defer stop()
	result, ok := next()
	if !ok {
		panic("Min of an empty sequence")
	}
	for {
		e, ok := next()
		if !ok {
			break
		}
		if e > result {
			result = e
		}
	}
	return result
}

// Head returns the first element of a sequence, or a default value for an empty sequence.
func Head[E any](s iter.Seq[E], def E) E {
	next, stop := iter.Pull(s)
	defer stop()
	result, ok := next()
	if ok {
		return result
	}
	return def
}

// Map returns a new sequence that contains the results of applying the given function to each element of the input sequence.
func Map[F ~func(I) O, I, O any](in iter.Seq[I], f F) iter.Seq[O] {
	return func(yield func(O) bool) {
		for x := range in {
			if !yield(f(x)) {
				return
			}
		}
	}
}

// MapRange returns a sequence with the results of calling function on a range of integers.
// The range is given in the usual half-open [start, end) style.
func MapRange[F ~func(I) E, I integer, E any](start, end I, f F) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i := start; i < end; i++ {
			if !yield(f(i)) {
				return
			}
		}
	}
}

// Filter returns a new sequence that contains just the elements matching a predicate.
func Filter[P ~func(E) bool, E any](s iter.Seq[E], p P) iter.Seq[E] {
	return func(yield func(E) bool) {
		for e := range s {
			if p(e) {
				if !yield(e) {
					return
				}
			}
		}
	}
}

// All returns true if the function f returns true for all elements of sequence s. Returns true for an empty sequence.
func All[F ~func(E) bool, E any](s iter.Seq[E], f F) bool {
	for e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

// Any returns true if the function f returns true for at least one element of sequence s. Returns false for an empty slice.
func Any[F ~func(E) bool, E any](s iter.Seq[E], f F) bool {
	for e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// Single returns an iterator that returns a single value when iterated over it.
func Single[T any](v T) iter.Seq[T] {
	return func(yield func(T) bool) {
		yield(v)
	}
}

// Seq returns an iterator that returns the provided argument values from the call, in order.
func Seq[T any](vs ...T) iter.Seq[T] {
	return S(vs)
}

// Chain returns an iterator that combines any number of iterators together.
func Chain[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Pairwise returns an iterator that iterates over all pairs of elements from a slice.
// The pairs are returned in order (0,1), (0,2), ..., (0,N-1), (1,2), ..., (N-2,N-1).
func Pairwise[E any, S ~[]E](s S) iter.Seq2[E, E] {
	return func(yield func(E, E) bool) {
		N := len(s)
		for i := 0; i < N-1; i++ {
			for j := i + 1; j < N; j++ {
				if !yield(s[i], s[j]) {
					return
				}
			}
		}
	}
}
