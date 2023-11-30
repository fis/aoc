// Binary z80run executes an AoC Z80 program.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/fis/aoc/z80/z80ex"
)

const usage = `usage: z80run prog.bin
`

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Print(usage)
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
	steps := cpu.Run(bufIn, bufOut)
	bufOut.Flush()

	fmt.Printf("# T = %d\n", steps)
}
