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
    space = {(0,0): dict(tile=Tile.FLOOR, path=[(0,0)])}
    fringe = set([(0, 0)])
    target = None

    while fringe:
        new_fringe = set()
        for pos in fringe:
            pos_path, moved = space[pos]['path'], False
            for step in ((pos[0],pos[1]-1), (pos[0],pos[1]+1), (pos[0]-1,pos[1]), (pos[0]+1,pos[1])):
                if step in space: continue  # already visited
                if not moved:
                    droid.must_move(pos_path[1:])
                    moved = True
                res = droid.try_move(step)
                if res == Tile.TARGET: target = step
                space[step] = dict(tile=res)
                if res != Tile.WALL:
                    space[step]['path'] = pos_path + [step]
                    new_fringe.add(step)
                    droid.must_move([pos])
            if moved:
                droid.must_move(pos_path[-2::-1])
        fringe = new_fringe

    return space, target

space, target = build_map(Droid())

tilech = { Tile.WALL: '#', Tile.FLOOR: '.', Tile.TARGET: 'O' }
min_x = min(c[0] for c in space.keys())
max_x = max(c[1] for c in space.keys())
min_y = min(c[1] for c in space.keys())
max_y = max(c[1] for c in space.keys())
for y in range(min_y, max_y+1):
    for x in range(min_x, max_x+1):
        ch = ' '
        if (x,y) in space: ch = tilech[space[(x,y)]['tile']]
        if x == 0 and y == 0: ch = '@'
        print(ch, end='')
    print('')

path = space[target]['path']
print(path)
print(len(path)-1)

# part 2

empty = set(p for p, c in space.items() if c['tile'] == Tile.FLOOR)

clock = 0
fringe = set([target])
while fringe:
    new_fringe = set()
    for pos in fringe:
        for step in ((pos[0],pos[1]-1), (pos[0],pos[1]+1), (pos[0]-1,pos[1]), (pos[0]+1,pos[1])):
            if step not in empty: continue
            empty.remove(step)
            new_fringe.add(step)
    fringe = new_fringe
    clock += 1

print(clock-1)
