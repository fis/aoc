#! /usr/bin/python3

import intcode
import sys

prog = intcode.load('day05-input.txt' if len(sys.argv) < 2 else sys.argv[1])

# part 1

intcode.run(prog, stdin=[1])

# part 2

intcode.run(prog, stdin=[5])
