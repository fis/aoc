#! /usr/bin/python3
# Copyright 2019 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


import heapq
import math
import sys

# solver

def solve(graph):
    best = {}
    frontier = [(0, tuple(v['num'] for v in graph['starts']), frozenset())]
    while frontier:
        fd, fvis, fkeys = heapq.heappop(frontier)
        if fkeys == graph['keys']:
            return fd
        for idx, fvi in enumerate(fvis):
            fv = graph['verts'][fvi]
            for ev, ed, ekeys, edoors in fv['edges']:
                if not fkeys >= edoors:
                    continue
                nd = fd + ed
                nvis = tuple(ev['num'] if i == idx else vi for i, vi in enumerate(fvis))
                nkeys = fkeys | ekeys
                if nd >= best.get((nvis, nkeys), math.inf):
                    continue
                best[(nvis, nkeys)] = nd
                heapq.heappush(frontier, (nd, nvis, nkeys))
        pass
    raise RuntimeError('mission: impossible')

# loader

def read_file(path):
    tiles = {}
    with open(path) as f:
        for y, line in enumerate(f.readlines()):
            for x, c in enumerate(line):
                if c >= 'a' and c <= 'z':
                    tiles[(x,y)] = ('key', c)
                elif c >= 'A' and c <= 'Z':
                    tiles[(x,y)] = ('door', c.lower())
                elif c == '@':
                    tiles[(x,y)] = ('entry', None)
                elif c == '.':
                    tiles[(x,y)] = ('path', None)
    return tiles

def build_graph(tiles):
    verts = {}
    for p, t in tiles.items():
        kind, arg = t
        if kind == 'key':
            verts[p] = dict(num=len(verts), key=arg, edges=[])
        elif kind == 'entry':
            verts[p] = dict(num=len(verts), edges=[])
    for start, startv in verts.items():
        fringe = [dict(at=start, keys=frozenset(), doors=frozenset())]
        seen = set()
        d = 0
        while fringe:
            d += 1
            new_fringe = []
            for p in fringe:
                seen.add(p['at'])
                px, py = p['at']
                for step in [(px-1,py), (px+1,py), (px,py-1), (px,py+1)]:
                    if step in seen or step not in tiles:
                        continue
                    stepk, stepv = tiles[step]
                    keys, doors = p['keys'], p['doors']
                    if stepk == 'key':
                        keys = keys | frozenset((stepv,))
                        startv['edges'].append((verts[step], d, keys, doors))
                    elif stepk == 'door':
                        doors = doors | frozenset((stepv,))
                    new_fringe.append(dict(at=step, keys=keys, doors=doors))
            fringe = new_fringe
    return dict(
        verts=sorted(verts.values(), key=lambda v: v['num']),
        starts=[v for v in verts.values() if 'key' not in v],
        keys=frozenset(v['key'] for v in verts.values() if 'key' in v),
    )

# program

if len(sys.argv) == 2:
    tiles = read_file(sys.argv[1])
    graph = build_graph(tiles)
    print(solve(graph))
    sys.exit(0)

tiles = read_file('day18-input.txt')

graph = build_graph(tiles)
print(solve(graph))

s = [p for p, t in tiles.items() if t[0] == 'entry']
sx, sy = s[0]
for dx, dy in [(-1,0), (0,-1), (0,0), (0,1), (1,0)]:
    del tiles[(sx+dx,sy+dy)]
for dx, dy in [(-1,-1), (-1,1), (1,-1), (1,1)]:
    tiles[(sx+dx,sy+dy)] = ('entry', None)

graph = build_graph(tiles)
print(solve(graph))
