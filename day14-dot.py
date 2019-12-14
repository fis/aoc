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
        inspecs, outspec = line.strip().split(' => ')
        outchem, outq = pspec(outspec)
        deps[outchem] = dict(pspec(inspec) for inspec in inspecs.split(', '))
        quanta[outchem] = outq

print('digraph reactions {')

for chem, q in quanta.items():
    print('  {0} [label="{0}\\n{1}"];'.format(chem, q))
for chem, chemdeps in deps.items():
    for dchem, dq in chemdeps.items():
        print('  {} -> {} [label="{}"];'.format(chem, dchem, dq))

print('}')
