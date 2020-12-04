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
import itertools
import sys

prog = intcode.load('day25-input.txt')

# run interactively:
#intcode.run(prog, stdin='ascii', stdout='ascii')
#sys.exit(0)

# solution:
# - asterisk, astronaut ice cream, fixed point, ornament
# - 134227456

# solver below

class Computer:
    def __init__(self, prog):
        self._vm = intcode.VM(prog)
        self._vm._stdin = self._in
        self._vm._stdout = self._out
        self._input = []
        self._output = []

    def read(self):
        while not self._output or self._output[-1] != 10:
            if not self._vm.step():
                break
        if not self._output:
            return None
        line = ''.join(chr(c) for c in self._output[:-1])
        self._output.clear()
        return line

    def read_until(self, key):
        lines = []
        while True:
            line = self.read()
            if line is None or line == key:
                break
            lines.append(line)
        return lines

    def write(self, line):
        self._input.extend([ord(c) for c in line])
        self._input.append(10)
        self._need_more_input = False

    def _in(self):
        if not self._input:
            raise RuntimeError('wanted input but nothing provided')
        return self._input.pop(0)

    def _out(self, n):
        self._output.append(n)

pickups = [
    'east',
    'take antenna',
    'north',
    'north',
    'take asterisk',
    'south',
    'west',
    'south',
    'take hologram',
    'north',
    'west',
    'take astronaut ice cream',
    'east',
    'east',
    'south',
    'east',
    'take ornament',
    'north',
    'west',
    'take fixed point',
    'east',
    'south',
    'west',
    'west',
    'south',
    'south',
    'south',
    'take dark matter',
    'north',
    'west',
    'north',
    'take monolith',
    'north',
    'north',
]

comp = Computer(prog)

comp.read_until('Command?')
for cmd in pickups:
    comp.write(cmd)
    comp.read_until('Command?')

items = sorted([
    'antenna', 'asterisk', 'hologram', 'astronaut ice cream',
    'ornament', 'fixed point', 'dark matter', 'monolith',
])

have = set(items)

for size in range(1, len(items)+1):
    for attempt in itertools.combinations(items, size):
        for item in items:
            if item in have and item not in attempt:
                comp.write(f'drop {item}')
                comp.read_until('Command?')
                have.remove(item)
            elif item not in have and item in attempt:
                comp.write(f'take {item}')
                comp.read_until('Command?')
                have.add(item)
        comp.write('east')
        result = comp.read_until('Command?')
        if '== Security Checkpoint ==' not in result:
            print('\n'.join(result))
            print('items:\n' + ', '.join(attempt))
            sys.exit(0)
