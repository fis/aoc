#! /usr/bin/python3

import re
import sys

# part 1

N = 10007

prog = []
with open('day22-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for line in f.readlines():
        line = line.rstrip()
        if line == 'deal into new stack':
            prog.append(('deal', 0))
            continue
        m = re.match(r'deal with increment (\d+)', line)
        if m:
            prog.append(('interleave', int(m.group(1))))
            continue
        m = re.match(r'cut (-?\d+)', line)
        if m:
            prog.append(('cut', int(m.group(1))))
            continue
        raise RuntimeError(line)

deck = list(range(N))

for op, arg in prog:
    if op == 'deal':
        deck = deck[::-1]
        continue
    if op == 'cut':
        deck = deck[arg:] + deck[:arg]
        continue
    if op == 'interleave':
        odeck = deck.copy()
        for i, c in enumerate(odeck):
            deck[arg*i % N] = c
        continue
    raise RuntimeError('{} {}'.format(op, arg))

print(deck.index(2019))

# part 2

def egcd(a, b):
    if a == 0:
        return (b, 0, 1)
    else:
        g, y, x = egcd(b % a, a)
        return (g, x - (b // a) * y, y)

def modinv(a, m):
    g, x, y = egcd(a, m)
    if g != 1:
        raise RuntimeError('modular inverse does not exist')
    else:
        return x % m

N = 119315717514047
reps = 101741582076661
pos = 2020

# sanity checks
#N = 10007
# 2019 -> 5169 -> 7674 -> 3233 -> 9756 -> 7176 -> 9413 -> 3234
#reps, pos = 1, 5169
#reps, pos = 7, 3234

mul, off = 1, 0
for op, arg in prog[::-1]:
    if op == 'deal':
        mul, off = (-mul) % N, (N-1) - off
        continue
    if op == 'cut':
        off = (off + arg) % N
        continue
    if op == 'interleave':
        inv = modinv(arg, N)
        mul, off = (mul * inv) % N, (off * inv) % N
        continue
    raise RuntimeError('{} {}'.format(op, arg))

pmul, poff = mul, off
mul, off, bit = 1, 0, 1
while bit <= reps:
    if reps & bit != 0:
        mul, off = (mul * pmul) % N, (off * pmul + poff) % N
    pmul, poff = (pmul * pmul) % N, (poff * (pmul + 1)) % N
    bit <<= 1

print((pos*mul + off) % N)

# The straight-forward inverse below; infeasible for solving the puzzle, due to the number of repetitions.
sys.exit(0)

for _ in range(reps):
    for op, arg in prog[::-1]:
        if op == 'deal':
            pos = (N-1) - pos
            continue
        if op == 'cut':
            pos = (pos + arg) % N
            continue
        if op == 'interleave':
            inv = modinv(arg, N)
            pos = (pos * inv) % N
            continue
        raise RuntimeError('{} {}'.format(op, arg))

print(pos)
