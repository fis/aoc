#! /usr/bin/python3

import heapq
import sys

# part 1

tiles = {}
with open('day20-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for y, line in enumerate(f.readlines()):
        for x, c in enumerate(line.rstrip()):
            tiles[(x,y)] = c

labelpairs = set()
labels = {}
paths = set()

min_x = min(x for x, _ in tiles.keys())
max_x = max(x for x, _ in tiles.keys())
min_y = min(y for _, y in tiles.keys())
max_y = max(y for _, y in tiles.keys())

for pos, c in tiles.items():
    if c == '.':
        paths.add(pos)
        continue
    if not c.isupper():
        continue
    x, y = pos
    if tiles.get((x,y+1), '').isupper() and tiles.get((x,y+2), '') == '.':
        l = c + tiles[(x,y+1)]
        if y == min_y:
            labelpairs.add(l)
            l += 'o'
        else:
            l += 'i'
        labels[l] = (x,y+2)
    elif tiles.get((x,y-1), '').isupper() and tiles.get((x,y-2), '') == '.':
        l = tiles[(x,y-1)] + c
        if y == max_y:
            labelpairs.add(l)
            l += 'o'
        else:
            l += 'i'
        labels[l] = (x,y-2)
    elif tiles.get((x+1,y), '').isupper() and tiles.get((x+2,y), '') == '.':
        l = c + tiles[(x+1,y)]
        if x == min_x:
            labelpairs.add(l)
            l += 'o'
        else:
            l += 'i'
        labels[l] = (x+2,y)
    elif tiles.get((x-1,y), '').isupper() and tiles.get((x-2,y), '') == '.':
        l = tiles[(x-1,y)] + c
        if x == max_x:
            labelpairs.add(l)
            l += 'o'
        else:
            l += 'i'
        labels[l] = (x-2,y)

rlabels = {}

for l, p in labels.items():
    rlabels[p] = l

nodes = dict((l, {}) for l in labels.keys())

for label, start in labels.items():
    seen = set([start])
    fringe = set([start])
    d = 0
    while fringe:
        new_fringe = set()
        for x, y in fringe:
            for step in ((x,y-1), (x,y+1), (x-1,y), (x+1,y)):
                if step in seen:
                    continue
                seen.add(step)
                if step in rlabels:
                    t = rlabels[step]
                    od = nodes[label].get(t,(None,))[0]
                    if od is None or d+1 < od:
                        nodes[label][t] = (d+1, 0)
                elif step in paths:
                    new_fringe.add(step)
        fringe = new_fringe
        d += 1

for l in labelpairs:
    lo, li = l+'o', l+'i'
    if lo in nodes and li in nodes:
        nodes[lo][li] = (1, -1)
        nodes[li][lo] = (1, 1)

def shortest_path(start, end):
    dist = {}
    seen = set()
    fringe = [(0, start)]
    while fringe:
        d, n = heapq.heappop(fringe)
        if n == end: return d
        if n in seen: continue
        seen.add(n)
        for m, e in nodes[n].items():
            w = e[0]
            if m in seen: continue
            md = d + w
            if m in dist and dist[m] <= md: continue  # have an okay path already
            dist[m] = md
            heapq.heappush(fringe, (md, m))

print(shortest_path('AAo', 'ZZo'))

# part 2

def shortest_recursive_path(start, end, max_dist):
    dist = {}
    seen = set()
    fringe = [(0, start, 0)]
    while fringe:
        d, n, nl = heapq.heappop(fringe)
        if d > max_dist: return None
        if n == end and nl == 0: return d
        if n in seen: continue
        seen.add((n, nl))
        for m, e in nodes[n].items():
            w, ml = e[0], nl + e[1]
            if ml < 0 or (m, ml) in seen: continue
            md = d + w
            if (m, ml) in dist and dist[(m, ml)] <= md: continue  # have an okay path already
            dist[(m, ml)] = md
            heapq.heappush(fringe, (md, m, ml))

print(shortest_recursive_path('AAo', 'ZZo', 100000))
