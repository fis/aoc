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

// Binary aoc2019 runs the AoC 2019 puzzle solutions.
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/fis/aoc-go/2019/days"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintln(os.Stderr, "Usage: aoc2019 [flags] N input.txt")
		os.Exit(1)
	}
	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Not a number: %q: %v\n", flag.Arg(0), err)
		os.Exit(1)
	}
	out, err := days.Solve(day, flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Solution failed: %v\n", err)
		os.Exit(1)
	}
	for _, s := range out {
		fmt.Println(s)
	}
}
