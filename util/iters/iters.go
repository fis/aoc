// Package iters contains custom iterators for convenience.
//
// Most of these iterators are variants of the list-returning functions of the util package.
package iters

import (
	"bufio"
	"io"
	"iter"
	"strings"

	"github.com/fis/aoc/util"
)

// Split returns an iterator over the string as split by the given delimiter.
func Split(s, delim string) iter.Seq[string] {
	return func(yield func(string) bool) {
		s := s
		for len(s) > 0 {
			next, tail, _ := strings.Cut(s, delim)
			s = tail
			if !yield(next) {
				return
			}
		}
	}
}

// Scanned returns an iterator over the tokens extracted by a bufio.Scanner with the given split function.
func Scanned(r io.Reader, split bufio.SplitFunc) iter.Seq[string] {
	return func(yield func(string) bool) {
		s := bufio.NewScanner(r)
		s.Split(split)
		for s.Scan() {
			if !yield(s.Text()) {
				break
			}
		}
	}
}

// Words iterates over all nonempty contiguous sequences of non-whitespace characters
// in the input string. In other words, this is the list of tokens defined by the standard
// bufio.ScanWords function.
func Words(s string) iter.Seq[string] {
	return Scanned(strings.NewReader(s), bufio.ScanWords)
}

// Lines iterates over the input string split by newlines, as if by bufio.ScanLines.
// In other words, it will not return an empty string even if the input has a trailing newline.
func Lines(s string) iter.Seq[string] {
	return Scanned(strings.NewReader(s), bufio.ScanLines)
}

// Chunks iterates over the input string as if by the ScanChunks function.
func Chunks(s string) iter.Seq[string] {
	return Scanned(strings.NewReader(s), util.ScanChunks)
}

// Ints iterates over all contiguous sequences of decimal digits parsed as integers,
// as defined by the ScanInts split function. Any non-integer data is ignored.
func Ints(s string) iter.Seq[int] {
	// This implementation is copied from util.Ints.
	return func(yield func(int) bool) {
		cur, in, neg := 0, false, false
		for i := range len(s) {
			b := s[i]
			if in && (b < '0' || b > '9') {
				if neg {
					cur = -cur
				}
				if !yield(cur) {
					return
				}
				cur, in, neg = 0, false, false
			}
			if b == '-' {
				neg = i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9'
			} else if b >= '0' && b <= '9' {
				in = true
				cur = cur*10 + int(b-'0')
			}
		}
		if in {
			if neg {
				cur = -cur
			}
			yield(cur)
		}
	}
}
