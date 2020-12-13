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


import curses
import intcode
import sys
import time

prog = intcode.load('day13-input.txt' if len(sys.argv) < 2 else sys.argv[1])

tiles = {0: ' ', 1: '@', 2: '#', 3: '-', 4: 'o'}

# part 1

out = []
intcode.run(prog, stdin=[], stdout=out)
if len(out) % 3 != 0:
    raise RuntimeError('output not multiple of three')

screen = dict()
for i in range(0, len(out), 3):
    x, y, t = out[i:i+3]
    screen[(x,y)] = t

W = max(c[0] for c in screen) + 1
H = max(c[1] for c in screen) + 1

for y in range(H):
    for x in range(W):
        t = screen.get((x,y), 0)
        print(tiles[t], end='')
    print('')

print(sum(1 for t in screen.values() if t == 2))

# part 2, interactive mode

prog[0] = 2  # free play

def game(scr):
    curses.curs_set(0)
    out = ['x', 0, 0]

    def game_in():
        ch = scr.getch()
        if ch == curses.KEY_LEFT: return -1
        if ch == curses.KEY_RIGHT: return 1
        return 0

    def game_out(n):
        if out[0] == 'x':
            out[0] = 'y'
            out[1] = n
        elif out[0] == 'y':
            out[0] = 't'
            out[2] = n
        else:
            out[0] = 'x'
            x, y = out[1:3]
            if x == -1 and y == 0:
                scr.addstr(0, 0, str(n))
            else:
                scr.addch(y+1, x, tiles[n])
            scr.refresh()

    intcode.run(prog, stdin=game_in, stdout=game_out)

#curses.wrapper(game)

# part 2, demo mode

def demo():
    out = ['x', 0, 0]
    pad_x, ball_x = [0], [0]
    score = [0]

    def demo_in():
        if ball_x[0] < pad_x[0]: return -1
        if ball_x[0] > pad_x[0]: return 1
        return 0

    def demo_out(n):
        if out[0] == 'x':
            out[0] = 'y'
            out[1] = n
        elif out[0] == 'y':
            out[0] = 't'
            out[2] = n
        else:
            out[0] = 'x'
            x, y = out[1:3]
            if x == -1 and y == 0:
                score[0] = n
            else:
                if n == 3: pad_x[0] = x
                if n == 4: ball_x[0] = x

    intcode.run(prog, stdin=demo_in, stdout=demo_out)
    return score[0]

print(demo())
