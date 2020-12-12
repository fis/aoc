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

	"github.com/fis/aoc-go/util"
)

// GenericSolver wraps a solution function that does all the work itself.
type GenericSolver func(io.Reader) ([]string, error)

// LineSolver wraps a solution that wants the lines of the input as strings.
type LineSolver func([]string) ([]int, error)

// ChunkSolver wraps a solution that wants the blank-line-separated paragraphs of the input as strings.
type ChunkSolver func([]string) ([]int, error)

// IntSolver wraps a solution that wants the input read in as whitespace-separated decimal integers.
type IntSolver func([]int) ([]int, error)

// LevelSolver wraps a solution that wants the lines of the input converted to a 2D level structure.
type LevelSolver struct {
	Solver func(*util.Level) ([]int, error)
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
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

// Solve implements the Solver interface.
func (s ChunkSolver) Solve(input io.Reader) ([]string, error) {
	data, err := util.ScanAll(input, util.ScanChunks)
	if err != nil {
		return nil, err
	}
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

// Solve implements the Solver interface.
func (s IntSolver) Solve(input io.Reader) ([]string, error) {
	rawData, err := util.ScanAll(input, bufio.ScanWords)
	if err != nil {
		return nil, err
	}
	data, err := atois(rawData)
	if err != nil {
		return nil, err
	}
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

// Solve implements the Solver interface.
func (s LevelSolver) Solve(input io.Reader) ([]string, error) {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	level := util.ParseLevel(data, s.Empty)
	ints, err := s.Solver(level)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

func atois(in []string) (out []int, err error) {
	for _, s := range in {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

func itoas(in []int) (out []string) {
	for _, i := range in {
		out = append(out, strconv.Itoa(i))
	}
	return out
}
