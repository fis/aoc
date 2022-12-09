// Copyright 2020 Google LLC
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

package glue

import (
	"bufio"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/fis/aoc/util"
)

// GenericSolver wraps a solution function that does all the work itself.
type GenericSolver func(io.Reader) ([]string, error)

// LineSolver wraps a solution that wants the lines of the input as strings.
type LineSolver func([]string) ([]string, error)

// ParsableLineSolver extends the regular LineSolver to also parse each line to an object of the given type.
type ParsableLineSolver[T any] struct {
	Solver func([]T) ([]string, error)
	Parser func(string) (T, error)
}

// ChunkSolver wraps a solution that wants the blank-line-separated paragraphs of the input as strings.
type ChunkSolver func([]string) ([]string, error)

// IntSolver wraps a solution that wants the input read in as a list of decimal integers.
// The separators can be any non-digit characters.
type IntSolver func([]int) ([]string, error)

// RegexpSolver wraps a solution that wants to match a single regular expression to each of the input lines.
// The solver is called with a slice of slices of all submatches. Note that this excludes the full pattern
// match that is typically included in regexp match outputs.
type RegexpSolver struct {
	Solver func([][]string) ([]string, error)
	Regexp string
}

// ParsableRegexpSolver extends the regular RegexpSolver to also parse each line to an object of the given type.
type ParsableRegexpSolver[T any] struct {
	Solver func([]T) ([]string, error)
	Regexp string
	Parser func([]string) (T, error)
}

// LevelSolver wraps a solution that wants the lines of the input converted to a 2D level structure.
type LevelSolver struct {
	Solver func(*util.Level) ([]string, error)
	Empty  byte
}

// Solve implements the Solver interface.
func (s GenericSolver) Solve(input io.Reader) ([]string, error) {
	return s(input)
}

// Solve implements the Solver interface.
func (s LineSolver) Solve(input io.Reader) ([]string, error) {
	data, err := util.ScanAll(input, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	out, err := s(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s ParsableLineSolver[T]) Solve(input io.Reader) ([]string, error) {
	lines, err := util.ScanAll(input, bufio.ScanLines)
	if err != nil {
		return nil, err
	}
	parsed := make([]T, len(lines))
	for i, line := range lines {
		parsed[i], err = s.Parser(line)
		if err != nil {
			return nil, err
		}
	}
	out, err := s.Solver(parsed)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s ChunkSolver) Solve(input io.Reader) ([]string, error) {
	data, err := util.ScanAll(input, util.ScanChunks)
	if err != nil {
		return nil, err
	}
	out, err := s(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s IntSolver) Solve(input io.Reader) ([]string, error) {
	data, err := util.ScanAllInts(input)
	if err != nil {
		return nil, err
	}
	out, err := s(data)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s RegexpSolver) Solve(input io.Reader) ([]string, error) {
	parsed, err := util.ScanAllRegexp(input, s.Regexp)
	if err != nil {
		return nil, err
	}
	out, err := s.Solver(parsed)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s ParsableRegexpSolver[T]) Solve(input io.Reader) ([]string, error) {
	matches, err := util.ScanAllRegexp(input, s.Regexp)
	if err != nil {
		return nil, err
	}
	parsed := make([]T, len(matches))
	for i, match := range matches {
		parsed[i], err = s.Parser(match)
		if err != nil {
			return nil, err
		}
	}
	out, err := s.Solver(parsed)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Solve implements the Solver interface.
func (s LevelSolver) Solve(input io.Reader) ([]string, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	level := util.ParseLevel(data, s.Empty)
	out, err := s.Solver(level)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Ints converts a list of ints to a list of strings.
func Ints(in ...int) (out []string) {
	for _, i := range in {
		out = append(out, strconv.Itoa(i))
	}
	return out
}
