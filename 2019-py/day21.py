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

prog = intcode.load('day21-input.txt' if len(sys.argv) < 2 else sys.argv[1])

def run(sprog):
    out = []
    intcode.run(prog, stdin=[ord(c) for c in '\n'.join(sprog) + '\n'], stdout=out)
    print(''.join(chr(c) for c in out if c < 256))
    return out[-1]

# part 1

sprog = [
    # jump if hole in any next three cells
    'NOT A T',
    'NOT T T',
    'AND B T',
    'AND C T',
    'NOT T J',
    # but only if ground available four cells in
    'AND D J',
    'WALK',
]
print(run(sprog))

# part 2

sprog = [
    # jump if hole in any next three cells & if ground available
    'NOT A T',
    'NOT T T',
    'AND B T',
    'AND C T',
    'NOT T J',
    'AND D J',
    # don't jump if it's a trap
    'NOT E T',
    'NOT T T',
    'OR H T',
    'AND T J',
    'RUN',
]
print(run(sprog))
