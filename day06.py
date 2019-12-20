#! /usr/bin/python3

import sys

# part 1

orbits = dict()

with open('day06-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for ln in f.readlines():
        parent, child = ln.rstrip().split(')')
        orbits.setdefault(parent, []).append(child)
    modules = [int(ln.rstrip()) for ln in f.readlines()]

def count_orbits(parent, depth=1):
    orbs = 0
    for child in orbits.get(parent, []):
        orbs += depth
        orbs += count_orbits(child, depth+1)
    return orbs

print(count_orbits('COM'))

# part 2

you, san = None, None

def find_endpoints(parent, path=[]):
    global you, san
    if parent == 'YOU': you = path
    if parent == 'SAN': san = path
    for child in orbits.get(parent, []):
        find_endpoints(child, path + [parent])

find_endpoints('COM')

#print(repr(you))
#print(repr(san))

if you is None or san is None:
    raise RuntimeException('not found')

while you and san and you[0] == san[0]:
    you = you[1:]
    san = san[1:]

#print(repr(you))
#print(repr(san))

print(len(you) + len(san))
