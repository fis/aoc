#! /usr/bin/python3

import enum
import intcode
import sys

class Tile(enum.Enum):
    WALL = 0
    FLOOR = 1
    TARGET = 2

class FakeDroid:
    def __init__(self):
        self.pos = (0, 0)
        self.target = (-1, 1)
        self.walls = set([
            (-3,-1), (-3,0), (-3,1),
            (-2,-2), (-2,2),
            (-1,-2), (-1,0), (-1,2),
            (0,-1), (0,2),
            (1,-1), (1,1),
            (2,0)
        ])

    def try_move(self, coord):
        if abs(coord[0]-self.pos[0]) + abs(coord[1]-self.pos[1]) != 1:
            raise RuntimeError('bad try_move: {} to {}'.format(self.pos, coord))
        if coord in self.walls:
            return Tile.WALL
        else:
            self.pos = coord
            return Tile.TARGET if self.pos == self.target else Tile.FLOOR

    def must_move(self, coords):
        for coord in coords:
            if self.try_move(coord) == Tile.WALL:
                raise RuntimeError('bad must_move: {} to {}: hit wall'.format(self.pos, coord))

class RealDroid:
    def __init__(self):
        self.pos = (0, 0)
        self.vm = intcode.VM(intcode.load('day15-input.txt'))

    def try_move(self, coord):
        if   coord[0] == self.pos[0]   and coord[1] == self.pos[1]-1: cmd = 1
        elif coord[0] == self.pos[0]   and coord[1] == self.pos[1]+1: cmd = 2
        elif coord[0] == self.pos[0]-1 and coord[1] == self.pos[1]:   cmd = 3
        elif coord[0] == self.pos[0]+1 and coord[1] == self.pos[1]:   cmd = 4
        else: raise RuntimeError('bad try_move: {} to {}'.format(self.pos, coord))
        res = self.vm.step_out(stdin=[cmd])
        if type(res) is not int or res < 0 or res > 2:
            raise RuntimeError('bad try_move: {} to {}: out {}'.format(self.pos, coord, res))
        res = Tile(res)
        if res != Tile.WALL: self.pos = coord
        return res

    def must_move(self, coords):
        for coord in coords:
            if self.try_move(coord) == Tile.WALL:
                raise RuntimeError('bad must_move: {} to {}: hit wall'.format(self.pos, coord))

Droid = RealDroid if len(sys.argv) < 2 or sys.argv[1] != 'test' else FakeDroid

# part 1

def build_map(droid):
    space = {(0,0): Tile.FLOOR}
    target = [None]

    def dfs(pos):
        for step in ((pos[0],pos[1]-1), (pos[0],pos[1]+1), (pos[0]-1,pos[1]), (pos[0]+1,pos[1])):
            if step in space: continue  # already visited
            res = droid.try_move(step)
            space[step] = res
            if res == Tile.WALL: continue  # way blocked
            if res == Tile.TARGET: target[0] = step
            dfs(step)
            droid.must_move([pos])

    dfs((0,0))
    return space, target[0]

space, target = build_map(Droid())

tilech = { Tile.WALL: '#', Tile.FLOOR: '.', Tile.TARGET: 'O' }
min_x = min(c[0] for c in space.keys())
max_x = max(c[1] for c in space.keys())
min_y = min(c[1] for c in space.keys())
max_y = max(c[1] for c in space.keys())
for y in range(min_y, max_y+1):
    for x in range(min_x, max_x+1):
        ch = ' '
        if (x,y) in space: ch = tilech[space[(x,y)]]
        if x == 0 and y == 0: ch = '@'
        print(ch, end='')
    print('')

def distances(space, start):
    out = {}
    fringe = set([(start, 0)])
    while fringe:
        new_fringe = set()
        for pos, dist in fringe:
            out[pos] = dist
            for step in ((pos[0],pos[1]-1), (pos[0],pos[1]+1), (pos[0]-1,pos[1]), (pos[0]+1,pos[1])):
                if step not in space or space[step] == Tile.WALL: continue  # way blocked
                if step in out: continue  # already visited
                new_fringe.add((step, dist+1))
        fringe = new_fringe
    return out

dists = distances(space, (0,0))
print(dists[target])

# part 2

dists = distances(space, target)
print(max(dists.values()))
