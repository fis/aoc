#! /usr/bin/python3

import math
import sys

def visible(rocks, pos):
    visible = []
    for candidate in rocks:
        dx, dy = candidate[0] - pos[0], candidate[1] - pos[1]
        if dx == 0 and dy == 0:
            continue
        d = math.gcd(dx, dy)
        dx, dy = dx/d, dy/d
        test = (pos[0] + dx, pos[1] + dy)
        while test != candidate and test not in rocks:
            test = (test[0] + dx, test[1] + dy)
        if test == candidate:
            visible.append(candidate)
    return visible

def angle(start, end):
    dx, dy = end[0] - start[0], end[1] - start[1]
    return math.fmod(math.atan2(dx, -dy) + 2*math.pi, 2*math.pi)

# part 1

asteroids = set()
with open('day10-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for y, row in enumerate(f.readlines()):
        for x, c in enumerate(row.rstrip()):
            if c == '#':
                asteroids.add((x, y))

max_see = 0
obs_pos = (0, 0)

for pos in asteroids:
    see = len(visible(asteroids, pos))
    if see > max_see:
        max_see = see
        obs_pos = pos

print(max_see)

# part 2

asteroids.remove(obs_pos)

shot = []
while asteroids:
    see = visible(asteroids, obs_pos)
    see.sort(key=lambda p: angle(obs_pos, p))
    shot.extend(see)
    asteroids -= set(see)

print('shot[199] = {}'.format(shot[199]))
