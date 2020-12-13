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


import intcode
import sys

prog = intcode.load('day19-input.txt' if len(sys.argv) < 2 else sys.argv[1])

def test_beam(x, y):
    out = []
    intcode.run(prog, stdin=[x, y], stdout=out)
    return out[0]

# part 1

N = 50

beam = set()

for y in range(N):
    for x in range(N):
        b = test_beam(x, y)
        if b: beam.add((x,y))
        print('#' if b else ' ', end='')
    print()

print(len(beam))

# part 2

cur_y = N-1
l = min(x for x, y in beam if y == cur_y)
r = max(x for x, y in beam if y == cur_y)

N = 100
beam_l, beam_r = [l] * N, [r] * N

box_x, box_y = None, None

while True:
    l, r = beam_l[cur_y % N], beam_r[cur_y % N]
    cur_y += 1
    while not test_beam(l, cur_y):
        l += 1
    while test_beam(r+1, cur_y):
        r += 1
    beam_l[cur_y % N] = l
    beam_r[cur_y % N] = r
    pl, pr = beam_l[(cur_y - N + 1) % N], beam_r[(cur_y - N + 1) % N]
    if l >= pl and l+N-1 <= pr:
        box_x, box_y = l, cur_y - N + 1
        break

print(box_x, box_y, 10000*box_x + box_y)
