#! /usr/bin/python3

# this script renders the day14 input as a dependency graph in .dot format
# run: ./day14-dot.py input.txt | dot -Tpng -oinput.png

import sys

def pspec(spec):
    q, chem = spec.split(' ')
    return chem, int(q)

deps, quanta = {}, {}
with open('day14-input.txt' if len(sys.argv) < 2 else sys.argv[1]) as f:
    for line in f.readlines():
        inspecs, outspec = line.rstrip().split(' => ')
        outchem, outq = pspec(outspec)
        deps[outchem] = dict(pspec(inspec) for inspec in inspecs.split(', '))
        quanta[outchem] = outq

print('digraph reactions {')

for chem, q in quanta.items():
    print(f'  {chem} [label="{chem}\\n{q}"];')
for chem, chemdeps in deps.items():
    for dchem, dq in chemdeps.items():
        print(f'  {chem} -> {dchem} [label="{dq}"];')

print('}')
