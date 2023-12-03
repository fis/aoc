// Copyright 2023 Google LLC
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

// Binary z80run executes an AoC Z80 program.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fis/aoc/z80/z80ex"
)

var (
	bound = flag.Int64("bound", -1, "only emulate up to (approximately) N t-states")
)

const usage = `usage: z80run [flags] prog.bin
Optional flags:
`

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprint(os.Stderr, usage)
		flag.PrintDefaults()
		os.Exit(1)
	}
	prog := flag.Arg(0)

	bin, err := os.ReadFile(prog)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if len(bin) > 65536 {
		fmt.Fprintf(os.Stderr, "%s: too large: %d bytes\n", prog, len(bin))
		os.Exit(1)
	}

	bufIn, bufOut := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	cpu := z80ex.NewCPU()
	cpu.WriteMem(bin, 0)

	var steps uint64
	if *bound <= 0 {
		steps = cpu.Run(bufIn, bufOut)
	} else {
		steps, _ = cpu.RunBounded(bufIn, bufOut, uint64(*bound))
	}

	bufOut.Flush()
	fmt.Printf("# T = %d\n", steps)
}
