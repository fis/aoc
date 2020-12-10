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

package util

import (
	"fmt"
	"strconv"
)

type Solver interface {
	Solve(path string) ([]string, error)
}

var solvers map[int]Solver

func RegisterSolver(day int, s Solver) {
	if solvers == nil {
		solvers = make(map[int]Solver)
	}
	if _, ok := solvers[day]; ok {
		panic(fmt.Sprintf("duplicate solvers: %d", day))
	}
	solvers[day] = s
}

func CallSolver(day int, path string) ([]string, error) {
	s, ok := solvers[day]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d", day)
	}
	return s.Solve(path)
}

// GenericSolver wraps a solution function that does all the work itself.
type GenericSolver func(string) ([]string, error)

// LineSolver wraps a solution that wants the lines of the input as strings.
type LineSolver func([]string) ([]int, error)

// ChunkSolver wraps a solution that wants the blank-line-separated paragraphs of the input as strings.
type ChunkSolver func([]string) ([]int, error)

// IntSolver wraps a solution that wants the lines of the input converted to integers.
type IntSolver func([]int) ([]int, error)

// LevelSolver wraps a solution that wants the lines of the input converted to a 2D level structure.
type LevelSolver struct {
	Solver func(*Level) ([]int, error)
	Empty  byte
}

func (s GenericSolver) Solve(path string) ([]string, error) {
	return s(path)
}

func (s LineSolver) Solve(path string) ([]string, error) {
	data, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

func (s ChunkSolver) Solve(path string) ([]string, error) {
	data, err := ReadChunks(path)
	if err != nil {
		return nil, err
	}
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

func (s IntSolver) Solve(path string) ([]string, error) {
	data, err := ReadIntRows(path)
	if err != nil {
		return nil, err
	}
	ints, err := s(data)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

func (s LevelSolver) Solve(path string) ([]string, error) {
	level, err := ReadLevel(path, s.Empty)
	if err != nil {
		return nil, err
	}
	ints, err := s.Solver(level)
	if err != nil {
		return nil, err
	}
	return itoas(ints), nil
}

func itoas(in []int) (out []string) {
	for _, i := range in {
		out = append(out, strconv.Itoa(i))
	}
	return out
}
