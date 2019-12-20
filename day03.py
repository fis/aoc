#! /usr/bin/python3

import sys

# part 1

with open('day03-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    line_a = f.readline().rstrip().split(',')
    line_b = f.readline().rstrip().split(',')

def to_segs(line):
    segs = []
    at = (0, 0)
    for cmd in line:
        d, n = cmd[0], int(cmd[1:])
        if d == 'U':
            end = (at[0], at[1]+n)
            seg = ('V', at[0], at[1], end[1])
        elif d == 'D':
            end = (at[0], at[1]-n)
            seg = ('V', at[0], end[1], at[1])
        elif d == 'L':
            end = (at[0]-n, at[1])
            seg = ('H', at[1], end[0], at[0])
        elif d == 'R':
            end = (at[0]+n, at[1])
            seg = ('H', at[1], at[0], end[0])
        else: sys.exit('bah')
        segs.append(seg)
        at = end
    return segs

segs_a = to_segs(line_a)
segs_b = to_segs(line_b)

crosses = []
for seg_a in segs_a:
    for seg_b in segs_b:
        h, v = seg_a, seg_b
        if h[0] == v[0]: continue  # assume parallel lines never cross
        if h[0] == 'V': h, v = v, h
        h_y, h_x0, h_x1 = h[1:]
        v_x, v_y0, v_y1 = v[1:]
        if h_y >= v_y0 and h_y <= v_y1 and v_x >= h_x0 and v_x <= h_x1 and (v_x != 0 or h_y != 0):
            crosses.append((v_x, h_y))

print(min(abs(x) + abs(y) for x, y in crosses))

# part 2

grid = {}

for i, line in ((0, line_a), (1, line_b)):
    at, delay = (0, 0), 0
    for cmd in line:
        d, n = cmd[0], int(cmd[1:])
        if d == 'U': dx, dy = 0, 1
        elif d == 'D': dx, dy = 0, -1
        elif d == 'L': dx, dy = -1, 0
        elif d == 'R': dx, dy = 1, 0
        else: sys.exit('bah')
        for _ in range(n):
            at = (at[0]+dx, at[1]+dy)
            delay += 1
            cell = grid.setdefault(at, [-1, -1])
            cell[i] = delay

delays = [v for v in grid.values() if v[0] != -1 and v[1] != -1]

print(min(sum(d) for d in delays))
