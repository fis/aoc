#! /usr/bin/python3

import intcode
import sys

# part 1

prog = intcode.load('day05-input.txt' if len(sys.argv) < 2 else sys.argv[1])
intcode.run(prog)
