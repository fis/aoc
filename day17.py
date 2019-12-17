#! /usr/bin/python3

import intcode
import sys

prog = intcode.load('day17-input.txt')

# part 1

def out2path(out):
    rows = ''.join(chr(c) for c in out).split('\n')
    while rows[-1] == '':
        rows.pop()
    path = set()
    for y, row in enumerate(rows):
        for x, c in enumerate(row):
            if c != '.': path.add((x,y))
    return path

out = []
intcode.run(prog, stdin=[], stdout=out)
print(''.join(chr(c) for c in out))

path = out2path(out)

align = 0
for x, y in path:
    if (x-1,y) in path and (x+1,y) in path and (x,y-1) in path and (x,y+1) in path:
        align += x * y

print(align)

# part 2

# "natural" path:
# R,8,L,4,R,4,R,10,R,8,R,8,L,4,R,4,R,10,R,8,L,12,L,12,R,8,R,8,R,10,R,4,R,4,L,12,L,12,R,8,R,8,R,10,R,4,R,4,L,12,L,12,R,8,R,8,R,10,R,4,R,4,R,10,R,4,R,4,R,8,L,4,R,4,R,10,R,8

# breakdown:
# --------------------
# R,8,L,4,R,4,R,10,R,8,  A
# R,8,L,4,R,4,R,10,R,8,  A
# L,12,L,12,R,8,R,8,     B
# R,10,R,4,R,4,          C
# L,12,L,12,R,8,R,8,     B
# R,10,R,4,R,4,          C
# L,12,L,12,R,8,R,8,     B
# R,10,R,4,R,4,          C
# R,10,R,4,R,4,          C
# R,8,L,4,R,4,R,10,R,8   A
# --------------------
# A,A,B,C,B,C,B,C,C,A

main = 'A,A,B,C,B,C,B,C,C,A\n'
sub_a = 'R,8,L,4,R,4,R,10,R,8\n'
sub_b = 'L,12,L,12,R,8,R,8\n'
sub_c = 'R,10,R,4,R,4\n'

robo_input = main + sub_a + sub_b + sub_c + 'n\n'

prog[0] = 2
out = []
intcode.run(prog, stdin=[ord(c) for c in robo_input], stdout=out)
print(out[-1])
