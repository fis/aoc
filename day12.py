#! /usr/bin/python3

import math
import re
import sys

def parse(line):
    m = re.match(r'<x=(-?\d+), y=(-?\d+), z=(-?\d+)>', line)
    if not m: raise RuntimeError('bad input: ' + line)
    return [int(m.group(1)), int(m.group(2)), int(m.group(3))]

with open('day12-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    moons_init = [dict(pos=parse(line.strip()), vel=[0,0,0]) for line in f.readlines()]

# part 1

moons = moons_init
steps = 1000 if len(sys.argv) < 3 else int(sys.argv[2])

for _ in range(steps):
    for i, a in enumerate(moons[:-1]):
        for b in moons[i+1:]:
            pa, pb = a['pos'], b['pos']
            va, vb = a['vel'], b['vel']
            for d in range(3):
                if pa[d] < pb[d]:
                    va[d] += 1
                    vb[d] -= 1
                elif pa[d] > pb[d]:
                    va[d] -= 1
                    vb[d] += 1
    for m in moons:
        p, v = m['pos'], m['vel']
        for d in range(3):
            p[d] += v[d]

print(sum(sum(abs(p) for p in m['pos']) * sum(abs(v) for v in m['vel']) for m in moons))

# part 2

n = len(moons)
cycles = []

for d in range(3):
    init_pos = [m['pos'][d] for m in moons]
    init_vel = [m['vel'][d] for m in moons]

    pos, vel = list(init_pos), list(init_vel)
    steps = 0
    while True:
        steps += 1
        for i in range(n-1):
            for j in range(i+1, n):
                if pos[i] < pos[j]:
                    vel[i] += 1
                    vel[j] -= 1
                elif pos[i] > pos[j]:
                    vel[i] -= 1
                    vel[j] += 1
        for i in range(n):
            pos[i] += vel[i]
        if pos == init_pos and vel == init_vel:
            break

    cycles.append(steps)

print(cycles)

while len(cycles) > 1:
    d = math.gcd(cycles[0], cycles[1])
    cycles = [cycles[0] * cycles[1] // d] + cycles[2:]

print(cycles[0])
