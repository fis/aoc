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

// Package glue contains the glue code to link together the different days' solvers.
//
// The expected way to use this package is:
// - Have your daily solver packages not export anything, but just have init functions
//   that call this package's `RegisterFoo` functions.
// - For each year, have a `days` package that blank-imports each day to pull them in.
// - To build the standard omnibus binary, blank-import each year, and call `Main`.
package glue

import (
	"fmt"
	"io"
	"os"
)

// Solver represents a function capable of solving one of the AoC puzzles.
type Solver interface {
	// Solves the puzzle by reading the given input, and producing a set of output lines.
	Solve(r io.Reader) ([]string, error)
}

// Plotter represents a function capable of turning an AoC puzzle input to a GraphViz dot graph.
type Plotter interface {
	Plot(r io.Reader, w io.Writer) error
}

type plotterRecord struct {
	plotter  Plotter
	examples map[string]string
}

// YearDay names an AoC puzzle by the event year (2015+) and puzzle day (1-25).
type YearDay struct{ Year, Day int }

var (
	solvers  map[YearDay]Solver
	plotters map[YearDay]plotterRecord
)

func init() {
	solvers = make(map[YearDay]Solver)
	plotters = make(map[YearDay]plotterRecord)
}

// RegisterSolver makes a solver known to the glue code as the nominated solver of the given day.
// This function is expected to be called from an `init` func.
func RegisterSolver(year, day int, s Solver) {
	yd := YearDay{year, day}
	if _, ok := solvers[yd]; ok {
		panic(fmt.Sprintf("duplicate solvers: %d %d", year, day))
	}
	solvers[yd] = s
}

// RegisterPlotter makes a plotter known to the glue code as the nominated plotter of the given day.
func RegisterPlotter(year, day int, p Plotter, examples map[string]string) {
	yd := YearDay{year, day}
	if _, ok := plotters[yd]; ok {
		panic(fmt.Sprintf("duplicate plotters: %d %d", year, day))
	}
	plotters[yd] = plotterRecord{plotter: p, examples: examples}
}

// Solve solves the given AoC puzzle using the provided input.
func Solve(year, day int, input io.Reader) ([]string, error) {
	s, ok := solvers[YearDay{year, day}]
	if !ok {
		return nil, fmt.Errorf("unknown day: %d %d", year, day)
	}
	return s.Solve(input)
}

// SolveFile calls Solve on an input file.
func SolveFile(year, day int, path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Solve(year, day, f)
}
