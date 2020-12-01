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
