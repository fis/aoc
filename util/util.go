// Copyright 2019 Google LLC
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

// Package util contains shared functions for several AoC days.
package util

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fis/aoc/util/ix"
)

// Words returns the list of all nonempty contiguous sequences of non-whitespace characters
// in the input string. In other words, this is the list of tokens defined by the standard
// bufio.ScanWords function.
func Words(s string) (words []string) {
	words, _ = ScanAll(strings.NewReader(s), bufio.ScanWords)
	return words
}

// Lines splits the input string by newlines as if by bufio.ScanLines. In other words, it
// will not return an empty string even if the input has a trailing newline.
func Lines(s string) (lines []string) {
	lines, _ = ScanAll(strings.NewReader(s), bufio.ScanLines)
	return lines
}

// Chunks splits the input string as if by the ScanChunks function.
func Chunks(s string) (chunks []string) {
	chunks, _ = ScanAll(strings.NewReader(s), ScanChunks)
	return chunks
}

// Ints returns the list of all contiguous sequences of decimal digits parsed as integers,
// as defined by the ScanInts split function.
func Ints(s string) (ints []int) {
	cur, in, neg := 0, false, false
	for i := 0; i < len(s); i++ {
		b := s[i]
		if in && (b < '0' || b > '9') {
			if neg {
				cur = -cur
			}
			ints = append(ints, cur)
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
		ints = append(ints, cur)
	}
	return ints
}

// ReadLines returns the contents of a text file as a slice of strings representing the lines. The
// newline separators are not kept. The last line need not have a newline character at the end.
func ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAll(f, bufio.ScanLines)
}

// ReadChunks returns the contents of a text file as a slice of strings representing all paragraphs,
// as defined by text separated by a blank line (two consecutive newlines).
func ReadChunks(path string) (chunks []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAll(f, ScanChunks)
}

// ReadInts parses a text file containing integers separated by any non-digits (see ScanInts).
func ReadInts(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAllInts(f)
}

// ReadRegexp parses a text file using a regular expression; see ScanAllRegexp for details.
func ReadRegexp(path, pattern string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ScanAllRegexp(f, pattern)
}

// ScanAll runs a scanner on a reader with the given split function, and returns all tokens.
func ScanAll(r io.Reader, split bufio.SplitFunc) (tokens []string, err error) {
	s := bufio.NewScanner(r)
	s.Split(split)
	for s.Scan() {
		tokens = append(tokens, s.Text())
	}
	return tokens, s.Err()
}

// ScanAllInts extracts all decimal integers from the reader.
func ScanAllInts(r io.Reader) (ints []int, err error) {
	s := bufio.NewScanner(r)
	s.Split(ScanInts)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, s.Err()
}

// ScanAllRegexp parses a reader's contents using a regular expression. The return value is a list of lists,
// containing each line's submatches. Note that unlike the usual convention, the match of the entire
// regular expression is not included.
func ScanAllRegexp(r io.Reader, pattern string) ([][]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	lines, err := ScanAll(r, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	parsed := make([][]string, len(lines))
	for i, line := range lines {
		parts := re.FindStringSubmatch(line)
		if parts == nil {
			return nil, fmt.Errorf("line %q does not match pattern %s", line, pattern)
		}
		parsed[i] = parts[1:]
	}
	return parsed, nil
}

// ScanChunks implements a bufio.SplitFunc for scanning paragraphs delimited by a blank line
// (i.e., two consecutive '\n' bytes).
func ScanChunks(data []byte, atEOF bool) (advance int, token []byte, err error) {
	delim := []byte{'\n', '\n'}

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, delim); i >= 0 {
		return i + 2, data[0:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

// ScanInts implements a bufio.SplitFunc for scanning decimal integers separated by any non-digits.
// An optional - can also be included as the first character of the token. "123-456" will be split
// to the two tokens "123", "-456".
func ScanInts(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	start := -1
	for i, b := range data {
		if (b >= '0' && b <= '9') || b == '-' {
			start = i
			break
		}
	}
	if start < 0 {
		return len(data), nil, nil
	}
	data = data[start:]
	advance += start
	end := 1
	for end < len(data) && data[end] >= '0' && data[end] <= '9' {
		end++
	}
	if end == len(data) && !atEOF {
		return advance, nil, nil
	}
	if end == 1 && data[0] == '-' {
		return advance + 1, nil, nil
	}
	return advance + end, data[:end], nil
}

// CheckPrefix tests if a string contains a prefix, and if so, removes it.
// If the string does not contain the prefix, it's returned unmodified.
// The `ok` result is true if the prefix was found.
func CheckPrefix(s, prefix string) (tail string, ok bool) {
	if !strings.HasPrefix(s, prefix) {
		return s, false
	}
	return s[len(prefix):], true
}

// NextWord returns the prefix of s up to the first space character, if any.
// If there is no space, the entire string is returned. The second return value
// gives the remainder of the input string.
func NextWord(s string) (word, tail string) {
	if sep := strings.IndexByte(s, ' '); sep >= 0 {
		return s[:sep], s[sep:]
	}
	return s, ""
}

// NextInt parses the leading decimal digits of s as a (nonnegative) decimal number,
// and returns both the parsed number and the remainder of the input. If there are
// no decimal digits, `ok` will be false.
func NextInt(s string) (n int, ok bool, tail string) {
	for len(s) > 0 && s[0] >= '0' && s[0] <= '9' {
		n, ok, s = n*10+int(s[0]-'0'), true, s[1:]
	}
	return n, ok, s
}

// SortBy is like slices.SortFunc except applies an accessor function to the objects.
// Comparing is then done using the natural ordering of the results.
func SortBy[S ~[]I, F ~func(I) O, I any, O cmp.Ordered](x S, f F) {
	slices.SortFunc(x, func(a, b I) int {
		return cmp.Compare(f(a), f(b))
	})
}

// LabelMap is a mapping from string labels to integer indices, with automatic allocation.
type LabelMap map[string]int

// Retrieves the corresponding index for a label, or allocates the next free one if it is new.
func (m LabelMap) Get(label string) int {
	if id, ok := m[label]; ok {
		return id
	}
	id := len(m)
	m[label] = id
	return id
}

// Splitter is a convenience type for iterating over the results of splitting a string without allocating a slice.
type Splitter string

// Empty returns true if the current state of the splitter is the empty string (no components remain).
func (s Splitter) Empty() bool {
	return len(s) == 0
}

// Count returns how many parts there would be in the string if it were to be split with a delimiter.
func (s Splitter) Count(delim string) int {
	return strings.Count(string(s), delim) + 1
}

// Next returns the part of the string leading up to the delimiter (if found), and also updates the splitter to retain the trailing part.
// If there is no delimiter, the entire contents are returned and the splitter becomes empty.
func (s *Splitter) Next(delim string) string {
	next, tail, _ := strings.Cut(string(*s), delim)
	*s = Splitter(tail)
	return next
}

// P represents a two-dimensional integer-valued coordinate.
type P struct {
	X, Y int
}

var (
	// MinP is the P with the most negative coordinates possible.
	MinP = P{math.MinInt, math.MinInt}
	// MaxP is the P with the most positive coordinates possible.
	MaxP = P{math.MaxInt, math.MaxInt}
)

// String formats the point in the most common (X,Y) style.
func (p P) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// GoString formats the point in the style of a Go structure.
func (p P) GoString() string {
	return fmt.Sprintf("util.P{%d,%d}", p.X, p.Y)
}

// Add returns the point with coordinates corresponding to the sum of the receiver and the other point.
func (p P) Add(q P) P {
	return P{p.X + q.X, p.Y + q.Y}
}

// AddXY returns Add(P{x, y}).
func (p P) AddXY(x, y int) P {
	return P{p.X + x, p.Y + y}
}

// Scale returns the point multiplied by a scalar. Useful for points representing vectors.
func (p P) Scale(n int) P {
	return P{n * p.X, n * p.Y}
}

// Neigh returns a point's von Neumann neighbourhood (the 4 orthogonally adjacent elements).
//
// The directions will be returned in this order: north, south, west, east.
func (p P) Neigh() [4]P {
	return [4]P{{p.X, p.Y - 1}, {p.X, p.Y + 1}, {p.X - 1, p.Y}, {p.X + 1, p.Y}}
}

// Neigh8 returns a point's Moore neighbourhood (the 8 orthogonally or diagonally adjacent elements).
func (p P) Neigh8() [8]P {
	return [8]P{
		{p.X - 1, p.Y - 1}, {p.X, p.Y - 1}, {p.X + 1, p.Y - 1},
		{p.X - 1, p.Y}, {p.X + 1, p.Y},
		{p.X - 1, p.Y + 1}, {p.X, p.Y + 1}, {p.X + 1, p.Y + 1},
	}
}

// DistM returns the Manhattan (taxicab, 4-neighbor) distance between two points.
func DistM(a, b P) int {
	return ix.Abs(a.X-b.X) + ix.Abs(a.Y-b.Y)
}

// DistC returns the Chebyshev (chessboard, 8-neighbor) distance between two points.
func DistC(a, b P) int {
	return max(ix.Abs(a.X-b.X), ix.Abs(a.Y-b.Y))
}

// Bounds returns the bounding box of a list of points.
func Bounds(points []P) (min, max P) {
	min, max = points[0], points[0]
	for _, p := range points[1:] {
		if p.X < min.X {
			min.X = p.X
		}
		if p.Y < min.Y {
			min.Y = p.Y
		}
		if p.X > max.X {
			max.X = p.X
		}
		if p.Y > max.Y {
			max.Y = p.Y
		}
	}
	return min, max
}

// ParseP parses a string in the "X,Y" format as a P.
func ParseP(s string) (P, error) {
	comma := strings.IndexByte(s, ',')
	if comma < 0 {
		return P{}, fmt.Errorf("no , in point: %q", s)
	}
	x, err := strconv.Atoi(s[:comma])
	if err != nil {
		return P{}, fmt.Errorf("bad X coordinate: %q: %w", s[:comma], err)
	}
	y, err := strconv.Atoi(s[comma+1:])
	if err != nil {
		return P{}, fmt.Errorf("bad Y coordinate: %q: %w", s[comma+1:], err)
	}
	return P{x, y}, nil
}
