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
	"cmp"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/google/subcommands"
)

// Main implements the main function for the combined AoC binary.
func Main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&solveCmd{}, "")
	subcommands.Register(&plotCmd{}, "")

	flag.Parse()
	os.Exit(int(subcommands.Execute(context.Background())))
}

// "aoc solve"

type solveCmd struct{}

func (*solveCmd) Name() string {
	return "solve"
}

func (*solveCmd) Synopsis() string {
	return "Solve an AoC puzzle."
}

func (*solveCmd) Usage() string {
	out := strings.Builder{}
	out.WriteString(`solve <year> <day> [input]:

  Solve one of the AoC puzzles.

  If an input file is not provided, but the file "testdata/YYYY/dayDD.txt"
  exists under the current directory, that file is used instead; this
  facilitates calling the command at the root of the repository. Otherwise,
  input is read from standard input, which you can explicitly request by
  passing "-" as the input file.

  Available days:
`)
	years, days := []int(nil), make(map[int][]int)
	for yd := range solvers {
		if _, ok := days[yd.Year]; !ok {
			years = append(years, yd.Year)
		}
		days[yd.Year] = append(days[yd.Year], yd.Day)
	}
	slices.Sort(years)
	for _, y := range years {
		fmt.Fprintf(&out, "    %d:", y)
		slices.Sort(days[y])
		for _, d := range days[y] {
			fmt.Fprintf(&out, " %d", d)
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (*solveCmd) SetFlags(_ *flag.FlagSet) {}

func (*solveCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() < 2 || f.NArg() > 3 {
		fmt.Fprintf(os.Stderr, "usage: solve <year> <day> [input]\n")
		return subcommands.ExitFailure
	}
	year, day, err := parseDay(f.Arg(0), f.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	in, inName, close, err := parseInput(f.Arg(2), year, day)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}
	defer close()

	out, err := Solve(year, day, in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", inName, err)
		return subcommands.ExitFailure
	}
	for _, line := range out {
		fmt.Println(line)
	}

	return subcommands.ExitSuccess
}

// "aoc plot"

type plotCmd struct{}

func (*plotCmd) Name() string {
	return "plot"
}

func (*plotCmd) Synopsis() string {
	return "Plot a graph related to an AoC puzzle."
}

func (*plotCmd) Usage() string {
	out := strings.Builder{}
	out.WriteString(`plot <year> <day> [input]:

  Plot a graph related to one of the AoC puzzles.

	Input handling is similar to the 'solve' command (see 'aoc help solve'), but
	the name can additionally be one of the built-in example inputs, typically
	"ex" or "exN". If you need to plot a file with the same name, use a path such
	as "./ex" instead.

	The output is by default as .dot source in standard output, but other options
  are available. TODO: implement other options

  Available days and their example inputs:
`)
	days, examples := []YearDay(nil), make(map[YearDay][]string)
	for yd, pr := range plotters {
		days = append(days, yd)
		for ex := range pr.examples {
			examples[yd] = append(examples[yd], ex)
		}
	}
	slices.SortFunc(days, func(a, b YearDay) int {
		if a.Year != b.Year {
			return cmp.Compare(a.Year, b.Year)
		}
		return cmp.Compare(a.Day, b.Day)
	})
	for _, yd := range days {
		fmt.Fprintf(&out, "    %d %d:", yd.Year, yd.Day)
		slices.Sort(examples[yd])
		for _, ex := range examples[yd] {
			fmt.Fprintf(&out, " %s", ex)
		}
		out.WriteRune('\n')
	}
	return out.String()
}

func (*plotCmd) SetFlags(_ *flag.FlagSet) {}

func (*plotCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if f.NArg() < 2 || f.NArg() > 3 {
		fmt.Fprintf(os.Stderr, "usage: plot <year> <day> [input]\n")
		return subcommands.ExitFailure
	}
	year, day, err := parseDay(f.Arg(0), f.Arg(1))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return subcommands.ExitFailure
	}

	pr, ok := plotters[YearDay{year, day}]
	if !ok {
		fmt.Fprintf(os.Stderr, "no plotter for: %d %d\n", year, day)
		return subcommands.ExitFailure
	}

	var (
		in     io.Reader
		inName string
	)
	if ex, ok := pr.examples[f.Arg(2)]; ok {
		in, inName = strings.NewReader(ex), fmt.Sprintf("<%s>", f.Arg(2))
	} else {
		var close func()
		in, inName, close, err = parseInput(f.Arg(2), year, day)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", inName, err)
			return subcommands.ExitFailure
		}
		defer close()
	}

	if err := pr.plotter.Plot(in, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", inName, err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

// common utilities

func parseDay(ya, da string) (y, d int, err error) {
	if y, err = strconv.Atoi(ya); err != nil {
		return 0, 0, fmt.Errorf("not a year: %s", ya)
	}
	if d, err = strconv.Atoi(da); err != nil {
		return 0, 0, fmt.Errorf("not a day: %s", da)
	}
	return y, d, nil
}

func parseInput(arg string, year, day int) (file *os.File, name string, close func(), err error) {
	// file, name, closable = os.Stdin, "<stdin>", false
	if arg != "" && arg != "-" {
		if file, err = os.Open(arg); err != nil {
			return nil, "", nil, err
		}
		return file, arg, func() { file.Close() }, nil
	} else if arg == "" {
		name = fmt.Sprintf("testdata/%04d/day%02d.txt", year, day)
		if file, err = os.Open(name); err == nil {
			return file, name, func() { file.Close() }, nil
		}
	}
	return os.Stdin, "<stdin>", func() {}, nil
}
